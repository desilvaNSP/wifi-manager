#!/bin/bash

DATE_COMMAND=$(which date)
TIME_STAMP=`${DATE_COMMAND} '+%Y-%m-%d.%H:%M:%S'`
CURRENT_DIR=`pwd`
SERVER_HOME=`cd ..;pwd`
export JWT_PRIVATE_KEY_PATH=/home/anuruddha/git/wifi-manager/server/resources/security/private.key
export JWT_PUBLIC_KEY_PATH=/home/anuruddha/git/wifi-manager/server/resources/security/public.key
export JWT_EXPIRATION_DELTA=72

function default_(){
  echo "server started successfully!"
  ./server.bin $SERVER_HOME
  echo $! > server.pid
}

function start_(){
    nohup ./server.bin $SERVER_HOME > ../logs/nohup.log 2>&1&
    echo $! > server.pid
    echo "server started successfully!"
}

function stop_(){
    kill -9 `cat server.pid`
    echo "server stoped successfully!"
    rm -rf server.pid
}

case "$1" in
        "")
           default_
           ;;

        start)
            start_
            ;;

        stop)
            stop_
            ;;

        status)
            process=$(ps -ef | grep server.bin | grep -v grep)
            if [ "$process" ]; then
             echo "server is up and running."
            else
             echo "server is not running at the moment."
            fi
            ;;
        restart)
            stop_
            start_
            ;;
        *)
            echo $"Usage: $0 {start|stop|restart|status}"
            exit 1
esac

