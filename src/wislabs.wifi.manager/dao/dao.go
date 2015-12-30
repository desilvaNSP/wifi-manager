package dao
import (
	"wislabs.wifi.manager/utils"
)

type NameValue struct{
	Name string  `db:"name"json:"name"`
	Value int  `db:"value"json:"value"`
}

type DashboardUser struct {
	TenantId int 	`db:"tenantid"json:"tenantid"`
	Username string `db:"username"json:"username"`
	Password string `db:"password"json:"password"`
	Email string 	`db:"email"json:"email"`
	Status string	`db:"status"json:"status"`
	Roles []string  `json:"roles"`
}

type PortalUser struct{
	TenantId int 				 		 `db:"tenantid"json:"tenantid"`
	Username string                      `db:"username"json:"username"`
	Password string                      `json:"password"`
	Acctstarttime utils.NullString       `db:"acctstarttime"json:"acctstarttime"`
	Acctlastupdatedtime utils.NullString `db:"acctlastupdatedtime"json:"acctlastupdatedtime"`
	Acctactivationtime utils.NullString  `db:"acctactivationtime"json:"acctactivationtime"`
	Acctstoptime utils.NullString        `db:"acctstoptime"json:"acctstoptime"`
	GroupName utils.NullString           `db:"groupname"json:"groupname"`
	ACL utils.NullString        		 `db:"acl"json:"acl"`
	Visits int64                         `db:"visits"json:"visits"`
}

type Role struct{
   Name string `json:"name"`
   TenantId string `json:"tenantId"`
}

type AuthUser struct {
	Username string `json:"username"`
	Role Role   `json:"role"`
}

type Constrains struct {
	From string `json:"from"`
	To string   `json:"to"`
	LocationId string    `json:"locationid"`
}

type ApLocation struct {
	TenantId int 				  `db:"tenantid"json:"tenantid"`
	LocationId int64    		  `db:"locationid"json:"locationid"`
	SSID string 				  `db:"ssid"json:"ssid"`
	MAC string   				  `db:"mac"json:"mac"`
	Longitude utils.NullFloat64   `db:"longitude"json:"longitude"`
	Latitude utils.NullFloat64    `db:"latitude"json:"latitude"`
	GroupName string			  `db:"groupname"json:"groupname"`
}

type DashboardMetric struct {
	TenantId int 				  `db:"tenantid"json:"tenantid"`
	MetricId int 				  `db:"metricid"json:"metricid"`
	Name string			 		  `db:"name"json:"name"`
}

type DashboardAppInfo struct {
	AppId int 				 	  `db:"appid"json:"appid"`
	TenantId int 				  `db:"tenantid"json:"tenantid"`
	Aggregate string			  `db:"aggregate"json:"aggregate"`
	Name string			 		  `db:"name"json:"name"`
	Users []DashboardAppUser	  `db:"users"json:"users"`
	Groups []DashboardAppGroup	  `db:"groups"json:"groups"`
	Metrics []DashboardAppMetric  `db:"metrics"json:"metrics"`
}

type DashboardApp struct {
	AppId int 				 	  `db:"appid"json:"appid"`
	TenantId int 				  `db:"tenantid"json:"tenantid"`
	Name string			 		  `db:"name"json:"name"`
	Aggregate string			  `db:"aggregate"json:"aggregate"`
}

type DashboardAppUser struct {
	TenantId int 				  `db:"tenantid"json:"tenantid"`
	AppId int 				 	  `db:"appid"json:"appid"`
	UserName string			 	  `db:"username"json:"username"`
}

type DashboardAppMetric struct {
	AppId int 				 	  `db:"appid"json:"appid"`
	MetricId int			 	  `db:"metricid"json:"metricid"`
}

type DashboardAppGroup struct {
	AppId int 				 	  `db:"appid"json:"appid"`
	GroupName string			  `db:"groupname"json:"groupname"`
}

type Response struct {
	Status string `json:"status"`
	Message string `json:"message"`
}

type DBConfigs struct{
	Username string
	Password string
	DBName string
	Host string
	Port int
}

type ServerConfigs struct{
	Sample string
	HttpPort int
	HttpsPort int
	ReadTimeOut int
	WriteTimeOut int
}