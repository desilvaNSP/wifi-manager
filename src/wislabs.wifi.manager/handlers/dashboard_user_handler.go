package handlers

import (
	"wislabs.wifi.manager/dao"
	"wislabs.wifi.manager/commons"
	"wislabs.wifi.manager/controllers/dashboard"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"strconv"
	"log"
	"wislabs.wifi.manager/authenticator"
	"wislabs.wifi.manager/utils"
	"wislabs.wifi.manager/controllers/location"
	"strings"
)

func AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user dao.DashboardUser
	err := decoder.Decode(&user)
	if err != nil {
		panic("Error while decoding json")
	}

	cookie := http.Cookie{}
	cookie.Name = "failed"
	cookie.Value = "success"

	response := dao.Response{}
	response.Status = "failed"
	if (dashboard.IsUserAuthenticated(user)) {
		cookie.Value = "success"
		response.Status = "success"
		response.Message = user.Username
	}

	r.AddCookie(&cookie)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}

func UserExistInTenantHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]
	tenantId, err := strconv.Atoi(r.Header.Get("tenantid"))
	if (err != nil) {
		log.Fatalln("Error while reading tenantid", err)
	}
	checkuser := dashboard.IsUserExistInTenant(tenantId, username);
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(checkuser); err != nil {
		panic(err)
	}
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user dao.DashboardUser
	err := decoder.Decode(&user)
	if err != nil {
		panic("Error while decoding json")
	}
	err = dashboard.RegisterDashboardUser(user)
	if (err != nil) {
		//log.Error(err.Error())
	}

	w.WriteHeader(http.StatusOK)
}

func UpdateUserbyAdminHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user dao.DashboardUser
	err := decoder.Decode(&user)
	if err != nil {
		panic("Error while decoding json")
	}
	err = dashboard.UpdateDashboardUser(user)
	w.WriteHeader(http.StatusOK)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var updateuser dao.DashboardUserDetails
	err := decoder.Decode(&updateuser)
	if err != nil {
		panic("Error while decoding json")
	}
	err = dashboard.UpdateDashboardUserDetails(updateuser)
	w.WriteHeader(http.StatusOK)
}

func UpdateUserPasswordHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user dao.DashboardUserResetPassword
	err := decoder.Decode(&user)
	if err != nil {
		panic("Error while decoding json")
	}
	tenantId, err := strconv.Atoi(r.Header.Get("tenantid"))
	err = dashboard.UpdateDashboardUserPassword(tenantId, user.Username, user.OldPassword, user.NewPassword)
	w.WriteHeader(http.StatusOK)
}

func DeleteDashboardUsersHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]
	tenantId, err := strconv.Atoi(vars["tenantid"])
	if (err != nil) {
		log.Fatalln("Error while reading tenantid", err)
	}
	dashboard.DeleteDashboardUser(tenantId, username)
	w.WriteHeader(http.StatusOK)
}

func GetUserProfile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]
	tenantId := utils.GetTenantId(r)
	user := dashboard.GetDashboardUser(tenantId, username)
	// Admin user is allowed to all the groups
	if (authenticator.IsUserAuthorized(username, authenticator.ADMIN, authenticator.ACTION_EXECUTE, r)) {
		user.ApGroups = location.GetAllLocationGroups(tenantId)
	}
	if (user.Username != "") {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(user); err != nil {
			panic(err)
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func GetDashboardUsersHandler(w http.ResponseWriter, r *http.Request) {
	if (!authenticator.IsAuthorized(authenticator.DASHBOARD_USERS, authenticator.ACTION_READ, r)) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	vars := mux.Vars(r)
	tenantid, err := strconv.Atoi(vars["tenantid"])
	if (err != nil) {
		log.Fatalln("Error while reading tenantid", err)
	}

	users := dashboard.GetAllDashboardUsers(tenantid)
	userInfo := make([]dao.UserInfo, len(users))

	for index, user := range users {
		userInfo[index].TenantId = user.TenantId
		userInfo[index].Username = user.Username
		userInfo[index].Email = user.Email
		userInfo[index].Status = user.Status
		userInfo[index].Permissions = dashboard.GetDashboardUserPermissions(user.TenantId, user.Username)
		userInfo[index].ApGroups = dashboard.GetDashboardUserApGroups(user.TenantId, user.Username)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(userInfo); err != nil {
		panic(err)
	}
}

func GetTenantRolesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tenantId := vars["tenantid"]
	tenantIdInt, _ := strconv.Atoi(tenantId)
	roles := dashboard.GetDashboardUserRoles(tenantIdInt)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(roles); err != nil {
		panic(err)
	}
}

func GetAllUserPermissionsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tenantId, err := strconv.Atoi(vars["tenantid"])
	if (err != nil) {
		log.Fatalln("Error while reading tenantid", err)
	}
	permissions := dashboard.GetAllDashboardUserPermissions(tenantId)

	userscopes := make(map[string][]dao.Permission)

	for _, scope := range permissions {
		userscopes[scope.Name] = append(userscopes[scope.Name], scope)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(userscopes); err != nil {
		panic(err)
	}
}

func GetLoginPage(w http.ResponseWriter, r *http.Request) {
	r.URL.Path = commons.ServerHome + "/webapps/dashboard/login.html"
	http.ServeFile(w, r, r.URL.Path)
}

func GetRegistrationPage(w http.ResponseWriter, r *http.Request) {
	r.URL.Path = commons.ServerHome + "/webapps/dashboard/register.html"
	http.ServeFile(w, r, r.URL.Path)
}

func GetUsersOfSSIDs(w http.ResponseWriter, r *http.Request) {
	ssids := strings.Split(r.FormValue("ssids"), ",")
	usernames := dashboard.GetUsernamesOfSSIDS(ssids)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(usernames); err != nil {
		panic(err)
	}
}

func AddSSIDsOfUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]
	tenantId, err := strconv.Atoi(r.Header.Get("tenantid"))
	if (err != nil) {
		log.Fatalln("Error while reading tenantid", err)
	}
	userId := dashboard.GetUserId(tenantId, username)
	decoder := json.NewDecoder(r.Body)
	var ssids []string
	err = decoder.Decode(&ssids)
	if (err != nil) {
		log.Fatalln("Error while decoding ssids json")
	}
	dashboard.AddUserSSIDS(userId, ssids)
	w.WriteHeader(http.StatusOK)
}

func GetSSIDsOfUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]
	tenantId, err := strconv.Atoi(r.Header.Get("tenantid"))
	if (err != nil) {
		log.Fatalln("Error while reading tenantid", err)
	}
	userId := dashboard.GetUserId(tenantId, username)

	ssids := dashboard.GetUserSSIDS(userId)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(ssids); err != nil {
		panic(err)
	}
}

func UpdateSSIDsOfUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]
	tenantId, err := strconv.Atoi(r.Header.Get("tenantid"))
	if (err != nil) {
		log.Fatalln("Error while reading tenantid", err)
	}
	userId := dashboard.GetUserId(tenantId, username)
	decoder := json.NewDecoder(r.Body)
	var ssids []string
	err = decoder.Decode(&ssids)
	if (err != nil) {
		log.Fatalln("Error while decoding ssids json")
	}
	dashboard.UpdateUserSSIDS(userId, ssids)
	w.WriteHeader(http.StatusOK)
}

