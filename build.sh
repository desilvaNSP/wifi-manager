#!/bin/bash
set -e
echo 'Exporting GO variables.'
DEPENDENCY_DIR='dependencies'
BUILD_DIR='target'
GOPATH_=/home/anuruddha/go-workspace

cd ..
PROJECT_ROOT_DIR=`pwd`

echo 'Getting the required dependencies...'
go get github.com/gorilla/mux


if [ ! -d $GOPATH_/src/core ] ; then
 ln -s $PROJECT_ROOT_DIR/hotspot-manager/core $GOPATH_/src
fi

echo 'Running GO build'
export GOBIN=$GOPATH_/bin
cd $GOPATH_
rm -f $GOPATH_/bin/core/main

echo `pwd`
go test -v core/main
go install core/dao
go install core/utils
go install core/controllers/location
go install core/controllers/wifi
go install core/handlers
go install core/routes
go install core/common
go install core/main

echo `pwd`
cp -f $GOPATH_/bin/main $PROJECT_ROOT_DIR/dashboard-web/bin/server.bin

echo 'GO build complete.'