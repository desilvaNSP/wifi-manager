#!/bin/bash
set -e

PROJECT_NAME='wifi-manager'
PROJECT_ROOT=`pwd`

echo 'Installing Gom'
go get github.com/mattn/gom

echo 'Exporting GO variables.'
export GOPATH=$GOPATH:$PROJECT_ROOT

cd src/wifi-manager/main
echo 'Gom install dependencies. This might take some time...'
gom install

gom build

echo "Executing test"
gom test -v

mv main $PROJECT_ROOT/server/bin/server.bin

echo 'GO build complete.'

mkdir -p $PROJECT_ROOT/target
cd $PROJECT_ROOT/target

echo "Removing existing distribution"
rm -rf $PROJECT_NAME.zip

if [ "$1" = "--release" ];then
 echo "Writing version information to versioninfo.md"
 DATE_COMMAND=$(which date)
 TIME_STAMP=`${DATE_COMMAND} '+%Y-%m-%d.%H:%M:%S'`

 echo "Time Stamp : ${TIME_STAMP}" > ../server/versioninfo.md
 LAST_COMMIT_ID=$(git log | head -1 | sed s/'commit '//)
 echo "Last Commit ID : ${LAST_COMMIT_ID}" >> ../server/versioninfo.md
 GIT_BRANCH=$(git branch)
 echo "Branch : ${GIT_BRANCH}" >> ../server/versioninfo.md
fi

echo "Start creating new distribution"
mkdir $PROJECT_NAME
cp -r ../server/* $PROJECT_NAME/
zip -rq $PROJECT_NAME.zip ./$PROJECT_NAME/* -x ./$PROJECT_NAME/logs/*
rm -rf $PROJECT_NAME
echo "Distribution creation complete."