package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
	"time"

	producer "github.com/Calaghan1/wb_tech_order_service.git/Poducer"
	"github.com/Calaghan1/wb_tech_order_service.git/database"
	"github.com/Calaghan1/wb_tech_order_service.git/helpers"
	"github.com/Calaghan1/wb_tech_order_service.git/schemas"
	"github.com/joho/godotenv"
	"github.com/nats-io/stan.go"
)

func Test_postres(t *testing.T) {
	err := godotenv.Load()
	var order schemas.Order
	data, err := ioutil.ReadFile("order.json")
	helpers.CheckError(err)
	err = json.Unmarshal(data, &order)
	helpers.CheckError(err)
	db := database.Init_database(os.Getenv("POSTRGRES__TEST_CONNECTION"))
	db.Migrator().DropTable(&schemas.Order{}, &schemas.Delivery{}, &schemas.Payment{}, &schemas.Item{})
	db.AutoMigrate(&schemas.Order{}, &schemas.Delivery{}, &schemas.Payment{}, &schemas.Item{})
	database.CreateNewOrder(db, order)
	orders := database.GetAllOrders(db)
	if len(orders) != 1 {
		t.Fatalf("expected %d, got %d", 1, len(orders))
	}
	if orders[0].OrderUID != order.OrderUID {
		t.Fatalf("expected %d, got %d", 1, len(orders))
	}
}

func Test_cache(t * testing.T) {
	var order schemas.Order
	err := godotenv.Load()
	helpers.CheckError(err)
	db := database.Init_database(os.Getenv("POSTRGRES__TEST_CONNECTION"))
	Cache := database.MakeCache(db)
	
	data, err := ioutil.ReadFile("order.json")
	helpers.CheckError(err)

	err = json.Unmarshal(data, &order)
	helpers.CheckError(err)
	if Cache[1].OrderUID != order.OrderUID {
		t.Fatal("Cache have not order")
	}
}


func Test_producer(t *testing.T) {
	go producer.SendRandomData()
	sc, err := stan.Connect(os.Getenv("NATS_STRIMNG_CLASTER"), os.Getenv("NATS_STRIMNG_CLIENT_ID_SUB"))
	helpers.CheckError(err)
	defer sc.Close()
	i := 0
	sub, _ := sc.Subscribe("Orders", func(m *stan.Msg,) {
		if len(m.Data) == 0 {
			t.Fatal("Dont Recive massage correctly")
		} 
		i ++
		if i > 4 {
			return
		}

	}, stan.DurableName("Test"))
	defer sub.Unsubscribe()
		time.Sleep(10 * time.Second)
}