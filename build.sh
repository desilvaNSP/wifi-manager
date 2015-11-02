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
go get golang.org/x/crypto/bcrypt


if [ ! -d $GOPATH_/src/$PROJECT_NAME/core ] ; then
 ln -s $PROJECT_ROOT_DIR/$PROJECT_NAME $GOPATH_/src
fi

echo 'Running GO build'
export GOBIN=$GOPATH_/bin
cd $GOPATH_
rm -f $GOPATH_/bin/main

#go test -v $PROJECT_NAME/core/main
go install $PROJECT_NAME/core/dao
go install $PROJECT_NAME/core/utils
go install $PROJECT_NAME/core/controllers/location
go install $PROJECT_NAME/core/controllers/wifi
go install $PROJECT_NAME/core/handlers
go install $PROJECT_NAME/core/routes
go install $PROJECT_NAME/core/common
go install $PROJECT_NAME/core/main

cp -f $GOPATH_/bin/main $PROJECT_ROOT_DIR/$PROJECT_NAME/server/bin/server.bin

echo 'GO build complete.'

mkdir -p $PROJECT_ROOT_DIR/$PROJECT_NAME/build
cd $PROJECT_ROOT_DIR/$PROJECT_NAME/build

echo "Removing existing distribution"
rm -rf build/$PROJECT_NAME.zip

echo "Writing version information to versioninfo.md"
DATE_COMMAND=$(which date)
TIME_STAMP=`${DATE_COMMAND} '+%Y-%m-%d.%H:%M:%S'`

echo "Time Stamp : ${TIME_STAMP}" > ../server/versioninfo.md
LAST_COMMIT_ID=$(git log | head -1 | sed s/'commit '//)
echo "Last Commit ID : ${LAST_COMMIT_ID}" >> ../server/versioninfo.md
GIT_BRANCH=$(git branch)
echo "Branch : ${GIT_BRANCH}" >> ../server/versioninfo.md

echo "Start creating new distribution"
mkdir $PROJECT_NAME
cp -r ../server/* $PROJECT_NAME/
zip -rq $PROJECT_NAME.zip ./$PROJECT_NAME/* -x ./$PROJECT_NAME/logs/*
rm -rf $PROJECT_NAME
echo "Distribution creation complete."