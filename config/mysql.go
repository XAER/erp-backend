package config

import (
	"backend/helpers"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	gorm_sql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetConnection() *gorm.DB {
	sql_host := helpers.GetEnv("SQL_HOST", "localhost")
	sql_port := helpers.GetEnv("SQL_PORT", "3306")
	sql_user := helpers.GetEnv("SQL_USER", "error")
	sql_pass := helpers.GetEnv("SQL_PASSWORD", "error")
	sql_db := helpers.GetEnv("SQL_DATABASE", "error")

	// Format for the SQL CONNECTION: DNS
	// <user>:<password>@tcp(<host>:<port>)/<dbname>
	sqlConnectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", sql_user, sql_pass, sql_host, sql_port, sql_db)
	db, err := gorm.Open(gorm_sql.Open(sqlConnectionString), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	return db
}
