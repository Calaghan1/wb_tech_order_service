package database

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Calaghan1/wb_tech_order_service.git/helpers"
	"github.com/Calaghan1/wb_tech_order_service.git/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init_database(acces_data string) *gorm.DB {
	// db, err := sql.Open("postgres", acces_data)
	// helpers.CheckError(err)
	// _, err = db.Exec("SET NAMES 'UTF8'")
	// helpers.CheckError(err)
	// err = db.Ping()
	// helpers.CheckError(err)
	db, err := gorm.Open(postgres.Open(acces_data), &gorm.Config{})
	helpers.CheckError(err)
	db.AutoMigrate(&schemas.Order{}, &schemas.Delivery{}, &schemas.Payment{}, &schemas.Item{})
	return db
}
func CreateNewOrder(db *gorm.DB, order schemas.Order) bool {
	result := db.Create(&order)
	helpers.CheckError(result.Error)
	return true
}
func GetOrderById(db * gorm.DB, id int) schemas.Order {
	var retrievedOrder schemas.Order
	res := db.Preload("Delivery").Preload("Payment").Preload("Items").Where("ID = ?", "3").First(&retrievedOrder)
	helpers.CheckError(res.Error)
	return retrievedOrder
}

func GetAllOrders(db * gorm.DB) []schemas.Order {
	var orders []schemas.Order
	res := db.Preload("Delivery").Preload("Payment").Preload("Items").Find(&orders)
	helpers.CheckError(res.Error)
	return orders
}

func PringMSg(ch chan []byte) {
	for data := range ch {
		fmt.Println(string(data))
	}
}

func UpdateDb(ch chan []byte, db *gorm.DB, cahe map[int]schemas.Order) {
	for data := range ch {
		var order schemas.Order
		json.Unmarshal(data, &order)
		psdata := db.Create(&order)
		helpers.CheckError(psdata.Error)
		cahe[order.ID] = order
		log.Println("DATABASE UPDATED")
	}
}

func MakeCache(db *gorm.DB) map[int]schemas.Order {
	orders := GetAllOrders(db)
	cache := make(map[int]schemas.Order)
	for _, order := range orders {
		cache[order.ID] = order
	}
	return cache
}