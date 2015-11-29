package dashboard

import (
	"wifi-manager/core/utils"
	"wifi-manager/core/dao"
	log "github.com/Sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"database/sql"
)

func IsUserAuthenticated(user dao.DashboardUser) bool{
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()

	var hashedPassword sql.NullString
	smtOut, err := dbMap.Db.Prepare("SELECT password FROM users where username=? ANd tenantid=? and status='active'")
	defer  smtOut.Close()
	err = smtOut.QueryRow(user.Username, user.TenantId).Scan(&hashedPassword)
	if err == nil && hashedPassword.Valid {
		if(len(hashedPassword.String) > 0){
			err = bcrypt.CompareHashAndPassword([]byte(hashedPassword.String), []byte(user.Password))
			if err == nil{
				log.Debug("User authenticated successfully " + user.Username)
			    return true
			}
		}
	}else{
		log.Debug("User authentication failed " + user.Username)
		return false
	}
	return false
}

func RegisterUser(user dao.DashboardUser) error{
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	stmtIns, err := dbMap.Db.Prepare("INSERT INTO users (tenantid, username, password, email, status) VALUES( ?, ?, ?, ?, ?)")
	defer stmtIns.Close()

	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(user.TenantId, user.Username, string(hashedPassword), user.Email, user.Status)
	return err
}

func UpdateUser(user dao.DashboardUser) error{
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()

	stmtIns, err := dbMap.Db.Prepare("UPDATE users SET email=?, status=? WHERE username=? and tenantid=?")
	defer stmtIns.Close()

	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(user.Email, user.Status, user.Username, user.TenantId)
	return err
}


func DeleteUser(tenantId int, username string){
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()

	stmtIns, err := dbMap.Db.Prepare("DELETE FROM users WHERE tenantid=? AND username=?")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	_, err = stmtIns.Exec(tenantId, username)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtIns.Close()
}

func GetUser(tenantId int, username string) dao.DashboardUser{
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()

	var user dao.DashboardUser
	err := dbMap.SelectOne(&user, "SELECT username,password,email FROM users WHERE username=? AND tenantid=?", username, tenantId)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	return user
}

func GetAllUsers(tenantId int) []dao.DashboardUser{
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()
	var users []dao.DashboardUser
	_,err := dbMap.Select(&users, "SELECT username, email, status FROM users WHERE tenantid=?",tenantId)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	return users
}

func GetRoles(tenantId int) []dao.Role{
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()
	var roles []dao.Role
	_,err := dbMap.Select(&roles, "SELECT name, tenantId FROM roles WHERE tenantid=?",tenantId)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	return roles
}