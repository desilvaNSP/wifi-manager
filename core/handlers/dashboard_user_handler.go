package handlers

import(
	"wifi-manager/core/dao"
	"wifi-manager/core/common"
	"wifi-manager/core/controllers/dashboard"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"strconv"
)

func AuthenticateUser(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)
	var user dao.DashboardUser
	err := decoder.Decode(&user)
	if err != nil {
		panic("Error while decoding json")
	}

	cookie := http.Cookie{}
	cookie.Name = "failed"
	cookie.Value="success"

	response := dao.Response{}
	response.Status = "failed"
	if(dashboard.IsUserAuthenticated(user)){
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
	err = dashboard.RegisterUser(user)

	response := dao.Response{}
	cookie := http.Cookie{}
	if err == nil {
		cookie.Name = "status"
		cookie.Value="success"
		response.Status = "pending"
	}else{
		cookie.Value = "fail"
		response.Status = "fail"
		response.Message = "Error while registering the user"
	}
	r.AddCookie(&cookie)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}

func DeleteDashboardUsersHandler(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	tenantId := vars["tenantid"]
	username := vars["username"]
	tenantIdInt,_ := strconv.Atoi(tenantId)
	dashboard.DeleteUser(tenantIdInt, username)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode("{}"); err != nil {
		panic(err)
	}
}

func GetUserProfile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tenantId := vars["tenantid"]
	username := vars["username"]
    var user dao.DashboardUser
	tenantIdInt, _ := strconv.Atoi(tenantId)

	user = dashboard.GetUser(tenantIdInt, username)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(user); err != nil {
		panic(err)
	}
}

func GetDashboardUsersHandler(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	tenantId := vars["tenantid"]
	tenantIdInt,_ := strconv.Atoi(tenantId)
	users := dashboard.GetAllUsers(tenantIdInt)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	b, err := json.Marshal(users)
	if err != nil {
		panic(err)
	}
	if err := json.NewEncoder(w).Encode(string(b)); err != nil {
		panic(err)
	}
}

func GetTenantRolesHandler(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	tenantId := vars["tenantid"]
	tenantIdInt,_ := strconv.Atoi(tenantId)
	roles := dashboard.GetRoles(tenantIdInt)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	b, err := json.Marshal(roles)
	if err != nil {
		panic(err)
	}
	if err := json.NewEncoder(w).Encode(string(b)); err != nil {
		panic(err)
	}
}

func GetLoginPage(w http.ResponseWriter, r *http.Request) {
	r.URL.Path = common.ServerHome + "/webapps/dashboard/login.html"
	http.ServeFile(w, r, r.URL.Path)
}

func GetRegistrationPage(w http.ResponseWriter, r *http.Request) {
	r.URL.Path = common.ServerHome + "/webapps/dashboard/register.html"
	http.ServeFile(w, r, r.URL.Path)
}
