#!/bin/bash
set -e

PROJECT_NAME='wifi-manager'
PROJECT_ROOT=`pwd`

echo 'Exporting GO variables.'
if [ -z "$GOPATH" ]; then
 echo "Build failed due to GOPATH has not been set."
 exit 1
fi
export GOPATH=$GOPATH:$PROJECT_ROOT

echo 'Installing Gom'
go get github.com/mattn/gom

cd src/wislabs.wifi.manager/main

echo 'Gom install dependencies. This might take some time...'
gom install

gom build

echo "Executing test"
#gom test -v

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
zip -rq $PROJECT_NAME.zip ./$PROJECT_NAME/* -x *.log -x *.out
rm -rf $PROJECT_NAME
echo "Distribution creation complete."