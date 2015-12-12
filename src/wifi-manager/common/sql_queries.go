package common

const GET_USER_COUNT_OF_DOWNLOADS_OVER_LOCATION string = "SELECT count(DISTINCT username) FROM dailyacct where date >= ? AND date < ? AND location = ? AND outputoctets >= ?";
const GET_USER_COUNT_OF_DOWNLOADS_OVER			string = "SELECT count(DISTINCT username) FROM dailyacct where date >= ? AND date < ? AND outputoctets >= ?";

const ADD_WIFI_USER_SQL string = "INSERT INTO accounting (username, acctstarttime, acctlastupdatedtime, acctstoptime, locationid) VALUES( ?, NOW(),NOW(),NOW()+ INTERVAL 1 HOUR, ? )";
const UPDATE_WIFI_USER_SQL string = "UPDATE accounting SET acl=? WHERE username=?";
const GET_ALL_WIFI_USER_SQL string = "SELECT username, acctstarttime, acctlastupdatedtime, acctstoptime, locationid, visits, acl FROM accounting order by username";

const DELETE_WIFI_USER string = "DELETE FROM accounting where username=?";
const DELETE_RADCHECk_USER string = "DELETE FROM radcheck WHERE username = ?";
const DELETE_RADACCT_USER string = "DELETE FROM radacct WHERE username = ?";

const GET_USER_COUNT_FROM_TO_LOCATION string = "SELECT COUNT(DISTINCT username) FROM accounting where acctstarttime >= ? AND acctstarttime < ? AND locationid = ?";
const GET_RETURNING_USERS_LOCATION string = "SELECT COUNT(DISTINCT username) FROM accounting where acctstarttime >= ? AND acctstarttime < ? AND locationid = ? AND visits > 1";

