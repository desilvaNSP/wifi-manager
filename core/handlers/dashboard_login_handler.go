package handlers

import(
	"wifi-manager/core/dao"
	"wifi-manager/core/utils"
	"wifi-manager/core/common"
	"net/http"
	"encoding/json"
	"strconv"
	"github.com/gorilla/mux"
)

func Login(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var u dao.User
	err := decoder.Decode(&u)
	if err != nil {
		panic("Error while decoding json")
	}

	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()
	var count int64
	count, err = dbMap.SelectInt("SELECT COUNT(*) FROM dashboarduser where username =  ? AND password = ? AND activated=1", u.Username, u.Password)
	checkErr(err, "Select failed")

	cookie := http.Cookie{}
	cookie.Name = "status"
	cookie.Value="success"
	response := dao.Response{}
	response.Status = "success"

	if(count!=1){
		cookie.Value = "fail"
		response.Status = "fail"
		response.Message = strconv.FormatInt(count,10) + u.Username
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
	var user dao.User
	err := decoder.Decode(&user)
	if err != nil {
		panic("Error while decoding json")
	}

	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()

	stmtIns, err := dbMap.Db.Prepare("INSERT INTO dashboarduser (username, password, email) VALUES( ?, ?, ? )")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	_, err = stmtIns.Exec(user.Username, user.Password, user.Email)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtIns.Close()

	response := dao.Response{}
	cookie := http.Cookie{}
	if err == nil {
		cookie.Name = "status"
		cookie.Value="success"
		response.Status = "success"
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

func GetUserInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()
	var user dao.User

	err := dbMap.SelectOne(&user, "SELECT username,password,email FROM dashboarduser where username = ?",username)
	checkErr(err, "Select failed")

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(user); err != nil {
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
