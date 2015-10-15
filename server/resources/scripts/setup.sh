#!/bin/bash
set -e

source setup_configs.sh

echo "Installing portal databse..."
mysql -u $DASHBOARD_DB_USERNAME -p$DASHBOARD_DB_PASSWORD -h $DASHBOARD_DB_HOST < ../sql/dashboard.sql
echo "Portal DB installed successfully."
