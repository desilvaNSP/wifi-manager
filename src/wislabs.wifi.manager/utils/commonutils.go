package utils

import (
	"database/sql"
	"gopkg.in/gorp.v1"
	"strconv"
	"github.com/spf13/viper"
	"fmt"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"log"
)

type DBConfigs struct{
	Username string
	Password string
	DBName string
	Host string
	Port int
}

var portalDBConfigs DBConfigs
var radiusDBConfigs DBConfigs
var dashboardDBConfigs DBConfigs
var summaryDBConfigs DBConfigs

var dbConfigs map[string] DBConfigs

func Init_(serverHome string){
	viper.New()
	viper.SetConfigName("config")
	viper.AddConfigPath(serverHome + "/configs")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	radiusDBConfigs = readDBConfigs("radiusDBConfigs")
	dashboardDBConfigs = readDBConfigs("dashboardDBConfigs")
	portalDBConfigs = readDBConfigs("portalDBConfigs")
	summaryDBConfigs = readDBConfigs("summaryDBConfigs")

	dbConfigs = make(map[string] DBConfigs)
	dbConfigs["radius"] = radiusDBConfigs
	dbConfigs["summary"] = summaryDBConfigs
	dbConfigs["portal"] = portalDBConfigs
	dbConfigs["dashboard"] = dashboardDBConfigs
}

func GetDBConnection(dbname string) *gorp.DbMap {
	var connectionUrl string
	if(dbConfigs != nil){
		dbConfig := dbConfigs[dbname]
		connectionUrl = dbConfig.Username + ":" + dbConfig.Password + "@tcp(" + dbConfig.Host + ":" + strconv.Itoa(dbConfig.Port)  + ")/" + dbConfig.DBName
	}
	db, err := sql.Open("mysql", connectionUrl)
	if err != nil {
		panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
	}
	dbmap := &gorp.DbMap{Db: db, Dialect:gorp.MySQLDialect{"InnoDB", "UTF8"}}
	return dbmap

}

func readDBConfigs(configName string) DBConfigs {
	configsMap := viper.GetStringMap(configName)
	dBConfigs := DBConfigs{}
	dBConfigs.Username = configsMap["username"].(string)
	dBConfigs.Password = configsMap["password"].(string)
	dBConfigs.DBName = configsMap["DBName"].(string)
	dBConfigs.Host = configsMap["host"].(string)
	dBConfigs.Port = configsMap["port"].(int)
	return dBConfigs
}


func SetDBConfigs(dbConfigs map[string] DBConfigs){
  dbConfigs = dbConfigs
}

func AddDBConfig(name string, dbConfig DBConfigs){
		dbConfigs[name] = dbConfig
}

func GetTenantId( r *http.Request)int{
	vars := mux.Vars(r)
	tenantId, err := strconv.Atoi(vars["tenantid"])
	if (err != nil) {
		log.Fatalln("Error while reading tenantid", err)
	}
	return tenantId
}

type NullString struct {
	sql.NullString
}

func (r NullString) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String)
}

func (r *NullString) UnmarshalJSON(data []byte) error{
	if string(data) == "null" {
		return nil
	}
	r.Valid = true
	return json.Unmarshal(data, (*string)(&r.String))
}

/* NullInt*/
type NullInt64 struct {
	sql.NullInt64
}

func (r NullInt64) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.Int64)
}

func (r *NullInt64) UnmarshalJSON(data []byte) error{
	if string(data) == "null" {
		return nil
	}
	r.Valid = true
	return json.Unmarshal(data, (*int64)(&r.Int64))
}

/* NullFloat64*/
type NullFloat64 struct {
	sql.NullFloat64
}

func (r NullFloat64) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.Float64)
}

func (r *NullFloat64) UnmarshalJSON(data []byte) error{
	if string(data) == "null" {
		return nil
	}
	r.Valid = true
	return json.Unmarshal(data, (*float64)(&r.Float64))
}