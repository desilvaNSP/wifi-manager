package dao
import (
	"wislabs.wifi.manager/utils"
	"net/textproto"
)

type NameValue struct {
	Name  string  `db:"name"json:"name"`
	Value float64  `db:"value"json:"value"`
}

type SummaryDailyAcctAll struct {
	Tenantid             int            `db:"tenantid"json:"tenantid"`
	Username             string            `db:"username"json:"username"`
	Date                 utils.NullString    `db:"date"json:"date"`
	Noofsessions         int            `db:"noofsessions"json:"noofsessions"`
	Totalsessionduration int            `db:"totalsessionduration"json:"totalsessionduration"`
	Sessionmaxduration   int            `db:"sessionmaxduration"json:"sessionmaxduration"`
	Sessionminduration   int            `db:"sessionminduration"json:"sessionminduration"`
	Sessionavgduration   int            `db:"sessionavgduration"json:"sessionavgduration"`
	Inputoctets          int64            `db:"inputoctets"json:"inputoctets"`
	Outputoctets         int64            `db:"outputoctets"json:"outputoctets"`
	Nasipaddress         string            `db:"nasipaddress"json:"nasipaddress"`
	Framedipaddress      string            `db:"framedipaddress"json:"framedipaddress"`
	Calledstationid      string            `db:"calledstationid"json:"calledstationid"`
	Ssid                 utils.NullString            `db:"ssid"json:"ssid"`
	Calledstationmac     utils.NullString            `db:"calledstationmac"json:"calledstationmac"`
	Groupname            utils.NullString            `db:"groupname"json:"groupname"`
	Locationid           utils.NullString            `db:"locationid"json:"locationid"`
}

type AccessPoint struct {
	TotalSessions         int            `db:"totalsessions"json:"totalsessions"`
	TotalUsers   int            `db:"totalusers"json:"totalusers"`
	AvgdataperUser   utils.NullString            `db:"avgdataperuser"json:"avgdataperuser"`
	Avgdatapersessiontime   utils.NullString           `db:"avgdatapersessiontime"json:"avgdatapersessiontime"`
	Totalinputoctets          int64            `db:"totalinputoctets"json:"totalinputoctets"`
	Totaloutputoctets         int64            `db:"totaloutputoctets"json:"totaloutputoctets"`
	Calledstationmac     utils.NullString            `db:"calledstationmac"json:"calledstationmac"`
}

type Tenant struct {
	TenantId  int       `db:"tenantid"json:"tenantid"`
	Domain    string    `db:"domain"json:"domain"`
	Status    string    `db:"status"json:"status"`
	CreatedOn string    `db:"createdon"json:"createdon"`
}

type DashboardUser struct {
	UserId      int64     `db:"userid"json:"userid"`
	TenantId    int       `db:"tenantid"json:"tenantid"`
	Username    string    `db:"username"json:"username"`
	Password    string    `db:"password"json:"password"`
	Email       string    `db:"email"json:"email"`
	Status      string    `db:"status"json:"status"`
	Roles       []string  `json:"roles"`
	Permissions []string  `json:"permissions"`
	ApGroups    []string  `json:"apgroups"`
}

type DashboardUserDetails struct {
	TenantId    int       `json:"tenantid"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	ContactNo    string  `json:"contactno"`
}

type DashboardUserResetPassword struct {
	Username    string  `json:"username"`
	OldPassword string  `json:"oldpassword"`
	NewPassword string  `json:"newpassword"`
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

type DataTablesResponce struct {
	Draw            int  `json:"draw"`
	RecordsTotal    int64 `json:"recordsTotal"`
	RecordsFiltered int64 `json:"recordsFiltered"`
	Data            []PortalUser `json:"data"`
	Error           string
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
	From       string            `json:"from"`
	To         string            `json:"to"`
	GroupNames []string          `json:"groupnames"`
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

type ApGroup struct {
	TenantId    int                  `db:"tenantid"json:"tenantid"`
	GroupName   string              `db:"groupname"json:"groupname"`
	GroupSymbol string              `db:"groupsymbol"json:"groupsymbol"`
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

type FileHeader struct {
	Filename string
	Header   textproto.MIMEHeader
	// contains filtered or unexported fields
}