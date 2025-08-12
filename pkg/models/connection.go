package models

import (
	"database/sql"
	"fmt"
	"mvc/pkg/config"
	"mvc/pkg/utils"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var DB *sql.DB

func InitDatabase() *sql.DB {

	var err error

	DB, err = sql.Open("mysql", config.GetConnectionString(""))

	_, err = DB.Exec("Create database if not exists ChefDB")
	utils.PanicIfErr(err, "Cant connect to Database")

	DB, err = sql.Open("mysql", config.GetConnectionString("ChefDB"))
	utils.PanicIfErr(err, "Can't connect to Database")

	driver, _ := mysql.WithInstance(DB, &mysql.Config{})
	m, err := migrate.NewWithDatabaseInstance("file://database/migrations", "mysql", driver)
	utils.PanicIfErr(err, "Can't Connect to Database")

	err = m.Up()
	utils.PanicIfErr(err, "Can't Run Migrations")

	DB.SetMaxOpenConns(25)
	DB.SetMaxIdleConns(5)
	DB.SetConnMaxLifetime(5 * time.Minute)

	err = DB.Ping()
	utils.PanicIfErr(err, "Database not responding")

	return DB
}

func CloseDatabase() {
	if DB != nil {
		fmt.Println("Closing database connection...")
		var err = DB.Close()
		utils.PanicIfErr(err, "Can't close database connection")
	}
}
