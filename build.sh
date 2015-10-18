#!/bin/bash
set -e
echo 'Exporting GO variables.'
DEPENDENCY_DIR='dependencies'
PROJECT_NAME='wifi-manager'
GOPATH_=/home/anuruddha/go-workspace

cd ..
PROJECT_ROOT_DIR=`pwd`

echo 'Getting the required dependencies...'
go get github.com/gorilla/mux


if [ ! -d $GOPATH_/src/$PROJECT_NAME/core ] ; then
 ln -s $PROJECT_ROOT_DIR/$PROJECT_NAME $GOPATH_/src
fi

echo 'Running GO build'
export GOBIN=$GOPATH_/bin
cd $GOPATH_
rm -f $GOPATH_/bin/main

#go test -v $PROJECT_NAME/core/main
go build -ldflags "-s" $PROJECT_NAME/core/dao
go build -ldflags "-s" $PROJECT_NAME/core/utils
go build -ldflags "-s" $PROJECT_NAME/core/controllers/location
go build -ldflags "-s" $PROJECT_NAME/core/controllers/wifi
go build -ldflags "-s" $PROJECT_NAME/core/handlers
go build -ldflags "-s" $PROJECT_NAME/core/routes
go build -ldflags "-s" $PROJECT_NAME/core/common
go install -ldflags "-s" $PROJECT_NAME/core/main

cp -f $GOPATH_/bin/main $PROJECT_ROOT_DIR/$PROJECT_NAME/server/bin/server.bin

echo 'GO build complete.'

cd $PROJECT_ROOT_DIR/$PROJECT_NAME/server
echo "Removing existing distribution"
rm -rf $PROJECT_NAME.zip

echo "Start creating new distribution"
zip -rq ../$PROJECT_NAME.zip ./* -x logs/*
echo "Distribution creation complete."