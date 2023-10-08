package main

import (
	"fmt"

	"github.com/Calaghan1/wb_tech_order_service.git/database"
	"github.com/Calaghan1/wb_tech_order_service.git/helpers"
	"github.com/Calaghan1/wb_tech_order_service.git/migrations"
)


const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "111"
	dbname   = "postgres"
)

func main() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)
	db := database.Init_database(connStr)
	err := db.Ping()
	helpers.CheckError(err)
	migrations.CreateTables(db)
	migrations.CheckAndSetDb(db)
	db.Close()
}