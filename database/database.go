package database

import (
	"database/sql"
	"github.com/Calaghan1/wb_tech_order_service.git/helpers"
	_ "github.com/lib/pq"
)

func Init_database(acces_data string) *sql.DB {
	db, err := sql.Open("postgres", acces_data)
	helpers.CheckError(err)
	_, err = db.Exec("SET NAMES 'UTF8'")
	helpers.CheckError(err)
	err = db.Ping()
	helpers.CheckError(err)
	return db
}