#!/bin/bash
read -p "Enter value for JWT_SECRET: " JWT_SECRET
read -p "Enter value for MYSQL_PORT: " MYSQL_PORT
read -p "Enter value for MYSQL_USER: " MYSQL_USER
read -p "Enter value for MYSQL_HOST: " MYSQL_HOST
read -p "Enter value for MYSQL_PASSWORD: " MYSQL_PASSWORD

cat > ".env" <<EOL
JWT_SECRET=$JWT_SECRET
MYSQL_PORT=$MYSQL_PORT
MYSQL_USER=$MYSQL_USER
MYSQL_HOST=$MYSQL_HOST
MYSQL_PASSWORD=$MYSQL_PASSWORD
EOL