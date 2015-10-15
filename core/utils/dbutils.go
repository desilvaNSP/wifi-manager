package utils

import (
	"database/sql"
	"gopkg.in/gorp.v1"
	"strconv"
	"github.com/spf13/viper"
	"fmt"
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
var radsummaryDBConfigs DBConfigs

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
	radsummaryDBConfigs = readDBConfigs("radsummaryDBConfigs")

	dbConfigs = make(map[string] DBConfigs)
	dbConfigs["radius"] = radiusDBConfigs
	dbConfigs["radsummary"] = radsummaryDBConfigs
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