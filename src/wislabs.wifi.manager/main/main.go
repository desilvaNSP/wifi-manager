package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/handlers"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"strconv"
	"time"
	"wislabs.wifi.manager/commons"
	"wislabs.wifi.manager/dao"
	"wislabs.wifi.manager/routes"
	"wislabs.wifi.manager/utils"
)

var ServerHome string
var serverConfigs dao.ServerConfigs
var logHandler http.Handler
var serverLogFile os.File
var httpAccessLogFile os.File

func main() {

	ServerHome = os.Getenv(commons.SERVER_HOME)
	if( len(ServerHome) <=0 ){
		ServerHome = os.Args[1]
	}

	initConfigurations(ServerHome)
	InitConfigs(ServerHome)
	commons.ServerHome = ServerHome
	defer serverLogFile.Close()
	defer httpAccessLogFile.Close()
	router := routes.NewRouter()
	router.PathPrefix("/dashboard/").Handler(http.StripPrefix("/dashboard/", http.FileServer(http.Dir(ServerHome+"/webapps/dashboard/"))))
	router.PathPrefix("/apiexplorer/").Handler(http.StripPrefix("/apiexplorer/", http.FileServer(http.Dir(ServerHome+"/webapps/apiexplorer/"))))
	router.PathPrefix("/apieditor/").Handler(http.StripPrefix("/apieditor/", http.FileServer(http.Dir(ServerHome+"/webapps/apieditor/"))))
	http.Handle("/", router)

	s := &http.Server{
		Addr:           ":" + strconv.Itoa(serverConfigs.HttpPort),
		Handler:        logHandler,
		ReadTimeout:    time.Duration(serverConfigs.ReadTimeOut) * time.Second,
		WriteTimeout:   time.Duration(serverConfigs.WriteTimeOut) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Info("Starting server on port : " + strconv.Itoa(serverConfigs.HttpPort))
	log.Fatal("HTTP Server error: ", s.ListenAndServeTLS(ServerHome+"/resources/security/server.pem", ServerHome+"/resources/security/server.key"))
}

func initConfigurations(serverHome string) {
	viper.New()
	viper.AddConfigPath(serverHome + "/configs")
	viper.SetConfigName("config")
	if _, err := os.Stat("../configs/config.yaml"); os.IsNotExist(err) {
		viper.SetConfigName("config.default")
	}
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	serverLogFile, err := os.OpenFile(serverHome+"/logs/server.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(serverLogFile)

	log.Info("Server home is set to : " + serverHome)

	httpAccessLogFile, err := os.OpenFile(serverHome+"/logs/http-access.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	enableHttpAccessLogs := viper.GetBool("httpAccessLogs")

	if enableHttpAccessLogs {
		logHandler = handlers.LoggingHandler(httpAccessLogFile, http.DefaultServeMux)
	}

	utils.Init_(serverHome)

	configsMap := viper.GetStringMap("serverConfigs")
	serverConfigs.HttpPort = configsMap["httpPort"].(int)
	serverConfigs.ReadTimeOut = configsMap["readTimeOut"].(int)
	serverConfigs.WriteTimeOut = configsMap["writeTimeOut"].(int)
}
