package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnDB() {
	dbDrv := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "go_inventaris"

	db, err := sql.Open(dbDrv, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err)
	}
	log.Println("Database connected!")
	DB = db
}
