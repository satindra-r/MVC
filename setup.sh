#!/bin/bash
mysql -uroot -ppass -e "CREATE DATABASE IF NOT EXISTS ChefDB;"
migrate -path ./database/migrations -database "mysql://root:pass@tcp(localhost:3306)/ChefDB?parseTime=true" up
