#!/bin/bash
read -p "Enter value for SERVER_PORT: " SERVER_PORT
read -p "Enter value for JWT_SECRET: " JWT_SECRET
read -p "Enter value for MYSQL_HOST: " MYSQL_HOST
read -p "Enter value for MYSQL_PORT: " MYSQL_PORT
read -p "Enter value for MYSQL_USER: " MYSQL_USER
read -p "Enter value for MYSQL_PASSWORD: " MYSQL_PASSWORD

cat > ".env" <<EOL
SERVER_PORT=$SERVER_PORT
JWT_SECRET=$JWT_SECRET
MYSQL_HOST=$MYSQL_HOST
MYSQL_PORT=$MYSQL_PORT
MYSQL_USER=$MYSQL_USER
MYSQL_PASSWORD=$MYSQL_PASSWORD
EOL

mysql -uroot -ppass -e "CREATE DATABASE IF NOT EXISTS ChefDB;"
migrate -path ./database/migrations -database "mysql://root:pass@tcp(localhost:3306)/ChefDB?parseTime=true" up
