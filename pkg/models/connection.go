package models

import (
	"database/sql"
	"fmt"
	"mvc/pkg/config"
	"mvc/pkg/utils"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDatabase() *sql.DB {

	var err error
	DB, err = sql.Open("mysql", config.GetConnectionString())

	utils.PanicIfErr(err, "Can't connect to database")

	DB.SetMaxOpenConns(25)
	DB.SetMaxIdleConns(5)
	DB.SetConnMaxLifetime(5 * time.Minute)

	err = DB.Ping()
	utils.PanicIfErr(err, "Database not responding")

	fmt.Println("Database connected successfully!")
	return DB
}

func CloseDatabase() {
	if DB != nil {
		fmt.Println("Closing database connection...")
		var err = DB.Close()
		utils.PanicIfErr(err, "Can't close database connection")
	}
}
