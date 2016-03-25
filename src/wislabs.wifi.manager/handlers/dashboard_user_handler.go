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
	tenantid, err := strconv.Atoi(vars["tenantid"])
	if (err != nil) {
		log.Fatalln("Error while reading tenantid", err)
	}
	dashboard.DeleteDashboardUser(tenantid, username)
	w.WriteHeader(http.StatusOK)
}

func GetUserProfile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]
	tenantId := utils.GetTenantId(r)
	var user dao.DashboardUser

	user = dashboard.GetDashboardUser(tenantId, username)
	if (user.Username != "") {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(user); err != nil {
			panic(err)
		}
	}else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func GetDashboardUsersHandler(w http.ResponseWriter, r *http.Request) {
	if (!authenticator.IsAuthorized("dashboard_users", authenticator.ACTION_READ, r)) {
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

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(users); err != nil {
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

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(permissions); err != nil {
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
