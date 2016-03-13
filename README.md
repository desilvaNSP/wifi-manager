####How to setup the  Dashboard development environment

* Install go-lang https://golang.org/
* Make sure you have set GOPATH  variable. (This variable must not point to the project folder)
* checkout the source code available at https://apremalal@bitbucket.org/apremalal/wifi-manager.git	
* Execute wifi-manager/build.sh 
    * This will create a distribution pack (wifi-manager.zip)
* Installing the DataBase
    * Create a database with name ‘dashboard’
    * Update the wifi-manager/server/resources/scripts/setup_config.sh with your mysql server configurations
    * Execute wifi-manager/server/resources/scripts/setup.sh
* Default configuration files are config.default.yaml and redis.default.conf, to change and override the defaults, simply
create new files config.yaml and redis.conf respectively and have your preferred configs. Do not change the default configuration files.

* To run the server execute wifi-manager/server/server.sh
* To run the server in daemon mode run server.sh start

* Point your browser to https://localhost:8081/dashboard/
* Username : admin@isl.com Password: admin

#### Adding  a dummy data set

* A dummy data set is located under wifi-manager/server/resources/sql/dummydata folder

1. unzip portaldump.sql.zip and source the file to portal database
 > source portaldump.sql
2. unzip sumarydump.sql.zip and source the file to portal database
 > source sumarydump.sql

#### IDE support for go-lang

* https://github.com/golang/go/wiki/IDEsAndTextEditorPlugins

#### Git & BitBucket
* Please refer https://bitbucket.org/apremalal/wifi-manager/wiki/Home
#### Configure Redis
* Redis[1] is an open source (BSD licensed), in-memory data structure store, used as database, cache and message broker. This 
 app use redis primarily as a JWT token storage. 
* This distribution contains an embedded redis instance compiled for Ubuntu 14.04 LTS. You have to replace thse redis-server 
with the matching redis server for the OS.

[1] http://redis.io/