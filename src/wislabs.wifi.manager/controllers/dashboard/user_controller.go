package dashboard

import (
	"wislabs.wifi.manager/utils"
	"wislabs.wifi.manager/dao"
	log "github.com/Sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"database/sql"
	"wislabs.wifi.manager/commons"
)

func IsUserAuthenticated(user dao.DashboardUser) bool {
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()

	var hashedPassword sql.NullString
	smtOut, err := dbMap.Db.Prepare("SELECT password FROM users where username=? ANd tenantid=? and status='active'")
	defer smtOut.Close()
	err = smtOut.QueryRow(user.Username, user.TenantId).Scan(&hashedPassword)
	if err == nil && hashedPassword.Valid {
		if (len(hashedPassword.String) > 0) {
			err = bcrypt.CompareHashAndPassword([]byte(hashedPassword.String), []byte(user.Password))
			if err == nil {
				log.Debug("User authenticated successfully " + user.Username)
				return true
			}
		}
	}else {
		log.Debug("User authentication failed " + user.Username)
		return false
	}
	return false
}

func RegisterDashboardUser(user dao.DashboardUser) error {
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	stmtIns, err := dbMap.Db.Prepare(commons.CREATE_DASHBOARD_USER)
	defer stmtIns.Close()

	if err != nil {
		return err
	}
	result, err := stmtIns.Exec(user.TenantId, user.Username, string(hashedPassword), user.Email, user.Status)
	if err != nil {
		return err
	} else {
		userId, _ := result.LastInsertId()
		//AddDashboardUserPermissions(userId, user)
		AddDashboardUserApGroups(userId, user)
	}
	return err
}

func UpdateDashboardUser(user dao.DashboardUser) error {
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()

	stmtIns, err := dbMap.Db.Prepare(commons.UPDATE_DASHBOARD_USER)
	defer stmtIns.Close()

	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(user.Email, user.Status, user.Username, user.TenantId)
	if err != nil {
		return err
	}
	AddDashboardUserApGroups(GetUserId(user.TenantId, user.Username), user)
	return err
}

func UpdateDashboardUserPassword(tenantId int, username string, oldPassword string, newPassword string)error {
	var user dao.DashboardUser
	user.Username = username
	user.Password = oldPassword
	user.TenantId = tenantId
	var err error
	if (IsUserAuthenticated(user)) {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		dbMap := utils.GetDBConnection("dashboard");
		defer dbMap.Db.Close()

		stmtIns, err := dbMap.Db.Prepare(commons.UPDATE_DASHBOARD_USER_PASSWORD)
		defer stmtIns.Close()

		if err != nil {
			return err
		}
		_, err = stmtIns.Exec(string(hashedPassword), username, tenantId)
		if err != nil {
			return err
		}
	}
	return err
}

func DeleteDashboardUser(tenantId int, username string) {
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()

	stmtIns, err := dbMap.Db.Prepare(commons.DELETE_DASHBOARD_USER)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	_, err = stmtIns.Exec(tenantId, username)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtIns.Close()
}

func GetDashboardUser(tenantId int, username string) dao.DashboardUser {
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()

	var user dao.DashboardUser
	err := dbMap.SelectOne(&user, commons.GET_DASHBOARD_USER, username, tenantId)
	if err != nil {
		//panic(err.Error()) // proper error handling instead of panic in your app
		return user
	}
	user.TenantId = tenantId
	user.Permissions = GetDashboardUserPermissions(tenantId, username)
	user.ApGroups = GetDashboardUserApGroups(tenantId, username)
	return user
}

func GetAllDashboardUsers(tenantId int) []dao.DashboardUser {
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()
	var users []dao.DashboardUser
	_, err := dbMap.Select(&users, commons.GET_ALL_DASHBOARD_USERS, tenantId)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	return users
}

func GetDashboardUserRoles(tenantId int) []dao.Role {
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()
	var roles []dao.Role
	_, err := dbMap.Select(&roles, "SELECT name, tenantId FROM roles WHERE tenantid=?", tenantId)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	return roles
}

func GetDashboardUserPermissions(tenantId int, username string) []string {
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()
	var permissions []string
	_, err := dbMap.Select(&permissions, commons.GET_DASHBOARD_USER_PERMISSIONS, username, tenantId)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	return permissions
}

func AddDashboardUserPermissions(userId int64, user dao.DashboardUser) error {
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()
	stmtIns, err := dbMap.Db.Prepare(commons.ADD_DASHBOARD_USER_PERMISSIONS)
	defer stmtIns.Close()

	if err != nil {
		return err
	}
	for _, permission := range user.Permissions {
		_, err := stmtIns.Exec(GetPermissionId(user.TenantId, permission), userId)
		if err != nil {
			return err
		}
	}
	return err
}

func UpdateDashboardUserPermissions() {

}

func GetAllDashboardUserPermissions(tenantId int) []dao.Permission {
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()
	var permissions []dao.Permission
	_, err := dbMap.Select(&permissions, commons.GET_ALL_PERMISSIONS, tenantId)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	return permissions
}

func GetPermissionId(tenantId int, permission string) int64 {
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()

	permissionId, err := dbMap.SelectInt(commons.GET_PERMISSION_ID, permission, tenantId)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	return permissionId
}

func GetUserId(tenantId int, username string) int64 {
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()

	permissionId, err := dbMap.SelectInt(commons.GET_USER_ID, username, tenantId)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	return permissionId
}

func AddDashboardUserApGroups(userId int64, user dao.DashboardUser) error {
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()
	stmtIns, err := dbMap.Db.Prepare(commons.ADD_DASHBOARD_USER_AP_GROUP)
	defer stmtIns.Close()

	if err != nil {
		return err
	}
	for _, groupName := range user.ApGroups {
		_, err := stmtIns.Exec(GetApGroupId(user.TenantId, groupName), userId)
		if err != nil {
			return err
		}
	}
	return err
}

func GetApGroupId(tenantId int, groupName string) int64 {
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()

	permissionId, err := dbMap.SelectInt(commons.GET_AP_GROUP_ID, groupName, tenantId)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	return permissionId
}

func GetDashboardUserApGroups(tenantId int, username string) []string {
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()
	var groups []string
	_, err := dbMap.Select(&groups, commons.GET_DASHBOARD_USER_AP_GROUPS, username, tenantId)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	return groups
}

