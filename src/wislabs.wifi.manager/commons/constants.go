package commons

const SERVER_HOME string = "SERVER_HOME"


const RADIUS_DB_NAME string    = "radius"
const DASHBOARD_DB_NAME string = "dashboard"
const PORTAL_DB_NAME string    = "portal"
const SUMMARY_DB_NAME string   = "summary"

/* common queries */
const GET_RECORDS_COUNT = "SELECT COUNT(*) from accounting WHERE tenantid=?"

/* Analytics */
const GET_USER_COUNT_OF_DOWNLOADS_OVER_LOCATION string = "SELECT count(DISTINCT username) FROM dailyacct where date >= ? AND date < ? AND location = ? AND outputoctets >= ?";
const GET_USER_COUNT_OF_DOWNLOADS_OVER		string = "SELECT count(DISTINCT username) FROM dailyacct where date >= ? AND date < ? AND outputoctets >= ?";


const GET_USER_COUNT_FROM_TO_LOCATION string = "SELECT COUNT(DISTINCT username) FROM accounting where acctstarttime >= ? AND acctstarttime < ? AND locationid = ? AND tenantid=?";
const GET_RETURNING_USERS_LOCATION string    = "SELECT COUNT(DISTINCT username) FROM accounting where acctstarttime >= ? AND acctstarttime < ? AND locationid = ? AND visits > 1 AND tenantid=?";

/* Dashboard Users */
const GET_ALL_DASHBOARD_USERS string 			= "SELECT tenantid, username, email, status FROM users WHERE tenantid=?"
const GET_ALL_PERMISSIONS string	 		= "SELECT permissionid, name, action FROM permissions WHERE tenantid=?"
const GET_PERMISSION_ID string	 			= "SELECT permissionid FROM permissions WHERE name= ? AND action=? AND tenantid=?"
const GET_USER_ID string	 			= "SELECT userid FROM users WHERE username= ? AND tenantid=?"
const GET_AP_GROUP_ID string	 			= "SELECT groupid FROM apgroups WHERE groupname= ? AND tenantid=?"
const GET_DASHBOARD_USER string      			= "SELECT userid, username, password, email, status FROM users WHERE username=? AND tenantid=?"
const CREATE_DASHBOARD_USER string      		= "INSERT INTO users (tenantid, username, password, email, status) VALUES( ?, ?, ?, ?, ?)"
const ADD_DASHBOARD_USER_AP_GROUP string    		= "INSERT IGNORE INTO userapgroups (groupid, userid) VALUES( ?, ?)"
const ADD_DASHBOARD_USER_PERMISSIONS string   		= "INSERT INTO userpermissions (permissionid, userid) VALUES( ?, ?)"
const GET_DASHBOARD_USER_PERMISSIONS string     	= "SELECT name, action  FROM permissions WHERE permissionid IN (SELECT permissionid FROM userpermissions WHERE  userid IN (SELECT  userid from users WHERE username=? AND tenantid=?)) GROUP BY name, action"
const GET_DASHBOARD_USER_AP_GROUPS string     		= "SELECT groupname  FROM apgroups WHERE groupid IN (SELECT groupid FROM userapgroups WHERE  userid IN (SELECT  userid from users WHERE username=? AND tenantid=?)) GROUP BY groupname"
const UPDATE_DASHBOARD_USER string      		= "UPDATE users SET email=?, status=? WHERE username=? and tenantid=?"
const UPDATE_DASHBORD_USER_PROFILE string               = "UPDATE users SET email=? WHERE username=? and tenantid=?"
const UPDATE_DASHBOARD_USER_PASSWORD string  		= "UPDATE users SET password=? WHERE username=? and tenantid=?"
const DELETE_DASHBOARD_USER string   			= "DELETE FROM users WHERE tenantid=? AND username=?"
const DELETE_DASHBOARD_USER_PERMISSIONS string		= "DELETE FROM userpermissions WHERE userid=?"
const DELETE_DASHBOARD_USER_APPGROUPS string 		= "DELETE FROM userapgroups WHERE userid=?"
const CHECK_EXISTS_USER_NAME	string				= "SELECT EXISTS(SELECT username FROM users WHERE username = ? and tenantid = ?) as checkuser"
/* WIFI users */
const ADD_WIFI_USER_SQL string  = "INSERT INTO accounting (tenantid, username, acctactivationtime, acctstarttime, maxsessionduration, groupname, acl) VALUES( ?, ?, NOW(),NOW(), 3600, ?, ? )";
const UPDATE_WIFI_USER string   = "UPDATE accounting SET acl=? WHERE username=? AND tenantid=?";
const GET_ALL_WIFI_USERS string = "SELECT tenantid, username, acctstarttime, acctlastupdatedtime, acctstoptime, groupname, visits, acl FROM accounting WHERE tenantid=? order by username LIMIT ?, ?";
const SEARCH_WIFI_USERS string  = "SELECT tenantid, username, acctstarttime, acctlastupdatedtime, acctstoptime, groupname, visits, acl FROM accounting WHERE tenantid=? AND username LIKE ? LIMIT ?,?";

const DELETE_WIFI_USER string     = "DELETE FROM accounting where username=? AND tenantid=?";
const DELETE_RADCHECk_USER string = "DELETE FROM radcheck WHERE username = ?";
const DELETE_RADACCT_USER string  = "DELETE FROM radacct WHERE username = ?";

/* AP locations */
const ADD_AP_LOCATION string 	          = "INSERT INTO aplocations (tenantid, ssid, mac, bssid, longitude, latitude, groupid,  groupname) VALUES( ?, ?, ?, ?, ?, ?, ?, ? )"
const UPDATE_AP_LOCATION string           ="UPDATE aplocations SET ssid=?,mac=?,bssid=?,longitude=?,latitude=?,groupid=?,groupname=? WHERE locationid=? and tenantid=? "
const ADD_AP_GROUP string 	 	  = "INSERT INTO apgroups (tenantid, groupname, groupsymbol) VALUES( ?, ?, ?)"
const GET_ALL_AP_LOCATIONS string         = "SELECT tenantid, locationid, ssid, mac, bssid, longitude, latitude, groupname FROM aplocations WHERE tenantid=?"
const GET_ALL_AP_GROUPS string	          = "SELECT distinct(groupname) FROM apgroups WHERE tenantid=?"
const DELETE_AP_LOCATION string           = "DELETE FROM aplocations WHERE ssid=? AND mac=? AND groupname=? AND tenantid=?"
const DELETE_AP_GROUP string 	          = "DELETE FROM aplocations WHERE groupname=? AND tenantid=?"
const DELETE_AP string 			  = "DELETE FROM aplocations WHERE mac=? AND tenantid=?"

/* Dashboard Apps */
const GET_DASHBOARD_APP string 		   = "SELECT appid, tenantid, name, aggregate FROM apps WHERE tenantid=? AND name=?"
const GET_DASHBOARD_APP_GROUPS string      = "SELECT appid, groupname FROM appgroups WHERE appid=?"
const GET_DASHBOARD_APP_ACLS string        = "SELECT appid, acl FROM appacls WHERE appid=?"
const GET_DASHBOARD_APP_AGGREGATE string   = "SELECT aggregate FROM apps WHERE appid=?"
const GET_DASHBOARD_APP_METRICS string     = "SELECT metricid, name FROM metrics WHERE metricid IN (SELECT metricid FROM appmetrics WHERE appid=?)"
const GET_DASHBOARD_APP_USERS string       = "SELECT tenantid, appid, username FROM appusers WHERE appid=?"
const GET_DASHBOARD_USER_APPS string       = "SELECT tenantid, appid, name, aggregate FROM apps WHERE appid IN (SELECT appid FROM appusers WHERE username=? AND tenantid=?)"
const ADD_DASHBOARD_APP string 		   = "INSERT INTO apps (tenantid, name, aggregate) VALUES( ?, ?, ?)"
const ADD_DASHBOARD_APP_USER string        = "INSERT INTO appusers (tenantid, appid, username) VALUES(?, ?, ? )"
const ADD_DASHBOARD_APP_METRIC string      = "INSERT INTO appmetrics (appid, metricid) VALUES( ?, ? )"
const ADD_DASHBOARD_APP_GROUP string       = "INSERT INTO appgroups (appid, groupname) VALUES( ?, ? )"
const ADD_DASHBOARD_ACLS string            = "INSERT INTO appacls (appid,acl) VALUES( ?, ?)"
const DELETE_DASHBOARD_APP string          = "DELETE FROM apps WHERE appid=? AND tenantid=?"
const DELETE_DASHBOARD_APP_USER string     = "DELETE FROM appusers WHERE appid=? AND username=?"

 /* App Settings */
const UPDATE_DB_APP_ACLS string                 = "UPDATE appacls SET acl=? WHERE appid=?"
const DELETE_OLD_DB_APP_GROUPS string           = "DELETE FROM appgroups WHERE appid=? and groupname=?"
const ADD_NEW_DB_APP_GROUPS string              = "INSERT INTO appgroups (appid,groupname) VALUES( ?, ?)"
const GET_EXIST_DASHBOARD_APP_METRICS string    = "SELECT metricid FROM appmetrics WHERE appid=?"
const DELETE_OLD_DB_APP_METRICS string          = "DELETE FROM appmetrics WHERE appid=? and metricid=?"
const ADD_NEW_DB_APP_METRICS string             = "INSERT INTO appmetrics (appid,metricid) VALUES( ?, ?)"
const UPDATE_DB_APP_AGGREGATE_VALUE  string     = "UPDATE apps SET aggregate=? WHERE appid=?"
const DELETE_OLD_DB_APP_USERS  string		= "DELETE FROM appusers WHERE appid=? and username=?"
const ADD_NEW_DB_APP_USERS string 		= "INSERT INTO appusers (tenantid, appid, username) VALUES( ?, ?, ?)"

/* Metrics */
const GET_ALL_DASHBOARD_METRICS string 		= "SELECT tenantid, metricid, name FROM metrics WHERE tenantid=?"
const GET_ALL_DASHBOARD_ACLS string    		= "SELECT DISTINCT acl FROM accounting"

/* ADD_RADIUS_SERVER */
const  ADD_RADIUS_SERVER string  		= "INSERT INTO radiusservers (tenantid,username,servername,serverip,authport,accounting,sharedsecret) VALUES( ?, ?, ?, ?, ?, ?, ?)"
const  GET_ALL_RADIUS_CONFIGS string	 	= "SELECT instid, servername, serverip, authport, accounting, sharedsecret FROM radiusservers WHERE tenantid=? and username=? "
const  DELETE_RADIUS_SERVER_INST string 	= "DELETE FROM radiusservers WHERE tenantid=? AND instid=?"
const  UPDATE_RADIUS_SERVER_INST string 	= "UPDATE radiusservers SET servername =?, serverip =?, authport= ?, accounting=?, sharedsecret=? WHERE tenantid=? and instid=?"

/* RADIUS */
const ADD_RADIUS_USER string 	 	   	= "INSERT INTO radcheck (username,attribute,op,value) VALUES( ?, ?, ?, ?)"
