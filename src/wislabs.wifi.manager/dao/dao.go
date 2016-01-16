package dao
import (
	"wislabs.wifi.manager/utils"
)

type NameValue struct {
	Name  string  `db:"name"json:"name"`
	Value float64  `db:"value"json:"value"`
}

type Tenant struct {
	TenantId  int       `db:"tenantid"json:"tenantid"`
	Domain    string    `db:"domain"json:"domain"`
	Status    string    `db:"status"json:"status"`
	CreatedOn string    `db:"createdon"json:"createdon"`
}

type DashboardUser struct {
	UserId      int64    `db:"userid"json:"userid"`
	TenantId    int    `db:"tenantid"json:"tenantid"`
	Username    string `db:"username"json:"username"`
	Password    string `db:"password"json:"password"`
	Email       string    `db:"email"json:"email"`
	Status      string    `db:"status"json:"status"`
	Roles       []string  `json:"roles"`
	Permissions []string `json:"permissions"`
	ApGroups    []string `json:"apgroups"`
}

type PortalUser struct {
	TenantId            int                         `db:"tenantid"json:"tenantid"`
	Username            string                      `db:"username"json:"username"`
	Password            string                      `json:"password"`
	Acctstarttime       utils.NullString       `db:"acctstarttime"json:"acctstarttime"`
	Acctlastupdatedtime utils.NullString `db:"acctlastupdatedtime"json:"acctlastupdatedtime"`
	Acctactivationtime  utils.NullString  `db:"acctactivationtime"json:"acctactivationtime"`
	Acctstoptime        utils.NullString        `db:"acctstoptime"json:"acctstoptime"`
	GroupName           utils.NullString           `db:"groupname"json:"groupname"`
	ACL                 utils.NullString                 `db:"acl"json:"acl"`
	Visits              int64                         `db:"visits"json:"visits"`
}

type Role struct {
	Name     string `json:"name"`
	TenantId string `json:"tenantId"`
}

type Permission struct {
	PermissionId int64    `json:"permissionid"`
	TenantId     int64        `json:"tenantid"`
	Name         string        `json:"name"`
	Action       string        `json:"action"`
}

type AuthUser struct {
	Username string `json:"username"`
	Role     Role   `json:"role"`
}

type Constrains struct {
	TenantId   int               `json:"tenantid"`
	From       string                   `json:"from"`
	To         string               `json:"to"`
	GroupNames []string           `json:"groupnames"`
}

type ApLocation struct {
	TenantId   int                  `db:"tenantid"json:"tenantid"`
	LocationId int64              `db:"locationid"json:"locationid"`
	SSID       string                  `db:"ssid"json:"ssid"`
	BSSID      string                  `db:"bssid"json:"bssid"`
	MAC        string                  `db:"mac"json:"mac"`
	Longitude  utils.NullFloat64   `db:"longitude"json:"longitude"`
	Latitude   utils.NullFloat64    `db:"latitude"json:"latitude"`
	GroupName  string              `db:"groupname"json:"groupname"`
}

type DashboardMetric struct {
	TenantId int                  `db:"tenantid"json:"tenantid"`
	MetricId int                  `db:"metricid"json:"metricid"`
	Name     string                      `db:"name"json:"name"`
}

type DashboardAppInfo struct {
	AppId     int                      `db:"appid"json:"appid"`
	TenantId  int                  `db:"tenantid"json:"tenantid"`
	Aggregate string              `db:"aggregate"json:"aggregate"`
	Name      string                      `db:"name"json:"name"`
	Users     []DashboardAppUser      `db:"users"json:"users"`
	Groups    []DashboardAppGroup      `db:"groups"json:"groups"`
	Metrics   []DashboardAppMetric  `db:"metrics"json:"metrics"`
}

type DashboardApp struct {
	AppId     int                      `db:"appid"json:"appid"`
	TenantId  int                  `db:"tenantid"json:"tenantid"`
	Name      string                      `db:"name"json:"name"`
	Aggregate string              `db:"aggregate"json:"aggregate"`
}

type DashboardAppUser struct {
	TenantId int                  `db:"tenantid"json:"tenantid"`
	AppId    int                      `db:"appid"json:"appid"`
	UserName string                  `db:"username"json:"username"`
}

type DashboardAppMetric struct {
	AppId    int                      `db:"appid"json:"appid"`
	MetricId int                  `db:"metricid"json:"metricid"`
}

type DashboardAppGroup struct {
	AppId     int                      `db:"appid"json:"appid"`
	GroupName string              `db:"groupname"json:"groupname"`
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type DBConfigs struct {
	Username string
	Password string
	DBName   string
	Host     string
	Port     int
}

type ServerConfigs struct {
	Sample       string
	HttpPort     int
	HttpsPort    int
	ReadTimeOut  int
	WriteTimeOut int
}