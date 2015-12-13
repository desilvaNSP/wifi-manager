#!/bin/bash
set -e

source setup_configs.sh

function default_(){
 echo "Installing dashboard databse..."
 mysql -u $DASHBOARD_DB_USERNAME -p$DASHBOARD_DB_PASSWORD -h $DASHBOARD_DB_HOST < ../sql/dashboard.sql
 echo "Portal DB installed successfully."
}

function clean_(){
    echo "Cleaning databses..."
    mysql -u $DASHBOARD_DB_USERNAME -p$DASHBOARD_DB_PASSWORD -h $DASHBOARD_DB_HOST < ../sql/cleanall.sql
    echo "Cleaning complete"
}

case "$1" in
        "")
           default_
           ;;

        clean)
            clean_
            ;;

        *)
            echo $"Usage: $0 {start|stop|restart|status}"
            exit 1
esac

