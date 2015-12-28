package common

/* Analytics */
const GET_USER_COUNT_OF_DOWNLOADS_OVER_LOCATION string = "SELECT count(DISTINCT username) FROM dailyacct where date >= ? AND date < ? AND location = ? AND outputoctets >= ?";
const GET_USER_COUNT_OF_DOWNLOADS_OVER			string = "SELECT count(DISTINCT username) FROM dailyacct where date >= ? AND date < ? AND outputoctets >= ?";

const GET_USER_COUNT_FROM_TO_LOCATION string = "SELECT COUNT(DISTINCT username) FROM accounting where acctstarttime >= ? AND acctstarttime < ? AND locationid = ? AND tenantid=?";
const GET_RETURNING_USERS_LOCATION string = "SELECT COUNT(DISTINCT username) FROM accounting where acctstarttime >= ? AND acctstarttime < ? AND locationid = ? AND visits > 1 AND tenantid=?";

/* Dashboard Users*/

const GET_ALL_DASHBOARD_USERS string = "SELECT tenantid, username, email, status FROM users WHERE tenantid=?"
const GET_DASHBOARD_USER string = "SELECT username,password,email FROM users WHERE username=? AND tenantid=?"
const DELETE_DASHBOARD_USER string = "DELETE FROM users WHERE tenantid=? AND username=?"

/* WIFI users*/
const ADD_WIFI_USER_SQL string = "INSERT INTO accounting (tenantid, username, acctstarttime, acctlastupdatedtime, acctstoptime, locationid) VALUES( ?, ?, NOW(),NOW(),NOW()+ INTERVAL 1 HOUR, ? )";
const UPDATE_WIFI_USER string = "UPDATE accounting SET acl=? WHERE username=? AND tenantid=?";
const GET_ALL_WIFI_USERS string = "SELECT tenantid, username, acctstarttime, acctlastupdatedtime, acctstoptime, groupname, visits, acl FROM accounting WHERE tenantid=? order by username";

const DELETE_WIFI_USER string = "DELETE FROM accounting where username=? AND tenantid=?";
const DELETE_RADCHECk_USER string = "DELETE FROM radcheck WHERE username = ? AND tenantid=?";
const DELETE_RADACCT_USER string = "DELETE FROM radacct WHERE username = ? AND tenantid=?";

/* AP locations */
const ADD_AP_LOCATION string = "INSERT INTO aplocations (tenantid, ssid, mac, longitude, latitude, groupname) VALUES( ?, ?, ?, ?, ?, ? )";
const GET_ALL_AP_LOCATIONS string = "SELECT tenantid, locationid, ssid, mac, longitude, latitude, groupname FROM aplocations WHERE tenantid=?"
const DELETE_AP_LOCATION string = "DELETE FROM aplocations WHERE ssid=? AND mac=? AND groupname=? AND tenantid=?"
const DELETE_AP_GROUP string = "DELETE FROM aplocations WHERE groupname=? AND tenantid=?"
const DELETE_AP string = "DELETE FROM aplocations WHERE mac=? AND tenantid=?"


///* AP Groups */
//const ADD_AP_GROUP string = "INSERT INTO apgroups (locationid, groupname) VALUES( ?, ? )";
//const GET_AP_GROUPS string = "SELECT groupid, locationid, groupname FROM apgroups"
//const DELETE_AP_GROUPS string = "DELETE FROM apgroups WHERE groupname = ?"
//const DELETE_AP_GROUP string = "DELETE FROM apgroup WHERE groupname = ?"