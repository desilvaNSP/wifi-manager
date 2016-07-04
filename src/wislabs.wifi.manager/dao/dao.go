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
    ACL                  utils.NullString                 `db:"acl"json:"acl"`
}

type AccessPoint struct {
    TotalSessions         int                 `db:"totalsessions"json:"totalsessions"`
    TotalUsers            int                     `db:"totalusers"json:"totalusers"`
    AvgdataperUser        utils.NullString      `db:"avgdataperuser"json:"avgdataperuser"`
    Avgdatapersessiontime utils.NullString   `db:"avgdatapersessiontime"json:"avgdatapersessiontime"`
    Totalinputoctets      int64            `db:"totalinputoctets"json:"totalinputoctets"`
    Totaloutputoctets     int64            `db:"totaloutputoctets"json:"totaloutputoctets"`
    Calledstationmac      utils.NullString         `db:"calledstationmac"json:"calledstationmac"`
    APName                utils.NullString           `db:"apname"json:"apname"`
}


type APSummaryDetails struct {
	CalledStationMac     utils.NullString         `db:"calledstationmac"json:"calledstationmac"`
	Value                int64                     `db:"summaryvalue"json:"summaryvalue"`
}


type LocationAccessPoint struct {
    AccessPointData AccessPoint
    LongLatMacData  LongLatMac
}

type LongLatMac struct {
    MAC       string                  `db:"mac"json:"mac"`
    APName    utils.NullString                `db:"apname"json:"apname"`
    Longitude utils.NullFloat64   `db:"longitude"json:"longitude"`
    Latitude  utils.NullFloat64    `db:"latitude"json:"latitude"`
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
    Permissions []Permission   `json:"permissions"`
    ApGroups    []string  `json:"apgroups"`
    SSIDs       []string   `json:"ssids"`
}

type UserInfo struct {
    TenantId    int       `db:"tenantid"json:"tenantid"`
    Username    string    `db:"username"json:"username"`
    Email       string    `db:"email"json:"email"`
    Status      string    `db:"status"json:"status"`
    Permissions []Permission   `json:"permissions"`
    ApGroups    []string  `json:"apgroups"`
}

type DashboardUserDetails struct {
    TenantId  int       `json:"tenantid"`
    Username  string    `json:"username"`
    Email     string    `json:"email"`
    ContactNo string  `json:"contactno"`
}

type DashboardUserResetPassword struct {
    Username    string  `json:"username"`
    OldPassword string  `json:"oldpassword"`
    NewPassword string  `json:"newpassword"`
}

type PortalUser struct {
    TenantId           int                    `db:"tenantid"json:"tenantid"`
    Username           string                 `db:"username"json:"username"`
    Password           string                 `json:"password"`
    AcctStartTime      utils.NullString       `db:"acctstarttime"json:"acctstarttime"`
    AcctActivationTime utils.NullString       `db:"acctactivationtime"json:"acctactivationtime"`
    MaxSessionDuration utils.NullInt64        `db:"maxsessionduration"json:"maxsessionduration"`
    GroupName          utils.NullString       `db:"groupname"json:"groupname"`
    ACL                utils.NullString       `db:"acl"json:"acl"`
    Visits             int64                  `db:"visits"json:"visits"`
    Accounting         string                  `db:"accounting"json:"accounting"`
}

type DataTablesResponse struct {
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
    PreFrom    string            `json:"prefrom"`
    PreTo      string            `json:"preto"`
    ACL        string            `json:"acl"`
    Criteria   string            `json:"criteria"`
    GroupNames []string          `json:"groupnames"`
	Parameters []string 		 `json:"parameters"`
}

type ApLocationSSIDs struct {
    TenantId   int                  `db:"tenantid"json:"tenantid"`
    SSID       []string                  `json:"ssids"`
    BSSID      string                  `db:"bssid"json:"bssid"`
    MAC        string                  `db:"mac"json:"mac"`
    GroupName  string              `db:"groupname"json:"groupname"`
	APs     	APs             `json:"aps"`
}

type ApLocation struct {
	TenantId   int                  `db:"tenantid"json:"tenantid"`
	SSID       string                  `db:"ssid"json:"ssid"`
	BSSID      string                  `db:"bssid"json:"bssid"`
	MAC        string                  `db:"mac"json:"mac"`
	GroupName  string              `db:"groupname"json:"groupname"`
	APName     utils.NullString     `db:"apname"json:"apname"`
	Address	   utils.NullString     `db:"address"json:"address"`
	APs     	APs             `json:"aps"`
}

type APs struct {
	TenantId   int                  `db:"tenantid"json:"tenantid"`
	MAC        string                  `db:"mac"json:"mac"`
	APName     utils.NullString     `db:"apname"json:"apname"`
	Address	   utils.NullString		`db:"address"json:"address"`
	Longitude  utils.NullFloat64   	`db:"longitude"json:"longitude"`
	Latitude   utils.NullFloat64    `db:"latitude"json:"latitude"`
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
	AppId          int64                   `db:"appid"json:"appid"`
	TenantId       int                   `db:"tenantid"json:"tenantid"`
	Aggregate      string                `db:"aggregate"json:"aggregate"`
	Name           string                `db:"name"json:"name"`
	FilterCriteria string           	 `db:"filtercriteria"json:"filtercriteria"`
	Parameters     []string				 `json:"parameters"`
	Users          []DashboardAppUser    `db:"users"json:"users"`
	Metrics        []DashboardAppMetric  `db:"metrics"json:"metrics"`
	Acls           string                `db:"acl"json:"acls"`
}

type DashboardGroups struct {
    TenantId int                  `db:"tenantid"json:"tenantid"`
    Groups   []DashboardAppGroup      `db:"groups"json:"groups"`
}

type DashboardApp struct {
    AppId     int                 `db:"appid"json:"appid"`
    TenantId  int                 `db:"tenantid"json:"tenantid"`
    Name      string              `db:"name"json:"name"`
    Aggregate string              `db:"aggregate"json:"aggregate"`
    FilterCriteria string         `db:"filtercriteria"json:"filtercriteria"`
}

type DashboardAppUser struct {
    TenantId int                  `db:"tenantid"json:"tenantid"`
    AppId    int                      `db:"appid"json:"appid"`
    UserName string                  `db:"username"json:"username"`
}

type DashboardAppMetric struct {
    MetricId int                  `db:"metricid"json:"metricid"`
    Name     string                      `db:"name"json:"name"`
}

type DashboardAppGroup struct {
    AppId     int                      `db:"appid"json:"appid"`
    GroupName string              `db:"groupname"json:"groupname"`
}

type DashboardAppAcls struct {
    AppId int                      `db:"appid"json:"appid"`
    Acls  string                     `db:"acl"json:"acls"`
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

type RadiusServer struct {
	InstanceId   int                `db:"InsId"json:"InsId"`
	TenantId     int             `db:"tenantid"json:"tenantid"`
	DBHostName   string            `db:"dbhostname"json:"dbhostname"`
	DBHostIp     string        `db:"dbhostip"json:"dbhostip"`
	DBSchemaName string            `db:"dbschemaname"json:"dbschemaname"`
	DBHostPort   int            `db:"dbport"json:"dbhostport"`
	DBUserName   string        `db:"dbusername"json:"dbusername"`
	DBPassword   string        `db:"dbpassword"json:"dbpassword"`
	Status       string            `db:"status"json:"status"`
}

type NasClient struct {
	NasClientID int                    `db:"id"json:"nasid"`
	NasName     string                `db:"nasname"json:"nasname"`
	ShortName   string                `db:"shortname"json:"shortname"`
	NasType     string                `db:"type"json:"nastype"`
	NasPorts    utils.NullInt64                `db:"ports"json:"nasport"`
	Secret      utils.NullString                `db:"secret"json:"secret"`
}

type  NasClientTestInfo struct {
	ServerIp      string  `json:"dbhostip"`
	NASClientName string    `json:"nasname"`
	Secret        string  `json:"secret"`
	AuthPort      string  `json:"authport"`
	UserName      string    `json:"username"`
	Password      string    `json:"password"`
}

type NASClientDBServer  struct {
	RadiusServerInfo RadiusServer  `json:"dbServer"`
	NASClientInfo    NasClient        `json:"nasClient"`
}

type SummaryChangePercentage struct {
	Value 	   int64 		`json:"value"`
	PreValue   int64		`json:"prevalue"`
	Percentage float64	        `json:"percentage"`
	Status 	   string		`json:"status"`
}

type AccessPointConstraints struct {
	TenantId     	int
	From     		string
	To     			string
	Threshold       int
	Query 			string
}
