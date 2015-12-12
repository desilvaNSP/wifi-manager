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
	"wifi-manager/core/common"
	"wifi-manager/core/dao"
	"wifi-manager/core/routes"
	"wifi-manager/core/utils"
)

var ServerHome string
var serverConfigs dao.ServerConfigs
var logHandler http.Handler
var serverLogFile os.File
var httpAccessLogFile os.File

func main() {
	ServerHome = os.Args[1]
	loadConfigs(ServerHome)
	common.ServerHome = ServerHome
	defer serverLogFile.Close()
	defer httpAccessLogFile.Close()
	router := routes.NewRouter()
	router.PathPrefix("/dashboard/").Handler(http.StripPrefix("/dashboard/", http.FileServer(http.Dir(ServerHome+"/webapps/dashboard/"))))
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

func loadConfigs(serverHome string) {
	viper.New()
	viper.SetConfigName("config")
	viper.AddConfigPath(serverHome + "/configs")

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
