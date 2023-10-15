package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	producer "github.com/Calaghan1/wb_tech_order_service.git/Poducer"
	"github.com/Calaghan1/wb_tech_order_service.git/database"
	"github.com/Calaghan1/wb_tech_order_service.git/helpers"
	"github.com/Calaghan1/wb_tech_order_service.git/schemas"
	"github.com/nats-io/stan.go"
	"gorm.io/gorm"
	// "github.com/Calaghan1/wb_tech_order_service.git/migrations"
	// "github.com/nats-io/stan.go"
	// "github.com/nats-io/nats.go"
)

func SendDataById(w http.ResponseWriter, r *http.Request, db *gorm.DB, cache *map[int]schemas.Order) {
	w.Header().Set("Content-Type", "application/json")
	idStr := r.URL.Path[len("/increment/"):]
	id, err := strconv.Atoi(idStr)
	helpers.CheckError(err)
	order, ok := (*cache)[id]
	if ok {
	response, err := json.Marshal(order)
	helpers.CheckError(err)
	w.Write(response)
	} else {
		json.NewEncoder(w).Encode(struct{}{})
	}
	
	
}

func main() {

	//INIT DATABASE AND CACHE
	dsn := "user=WB_TECH dbname=WB_TECH password=111 sslmode=disable"
	db := database.Init_database(dsn)
	cache := database.MakeCache(db)
	cahe_pointer := &cache
	//INIT SUBSCRIBER AND RODUCER NATS STRIMING
	sc, err := stan.Connect("test-cluster", "client-2")
	helpers.CheckError(err)
	defer sc.Close()
	ch := make(chan []byte, 10)
	sub, _ := sc.Subscribe("Orders", func(m *stan.Msg,) {
		ch <- m.Data
		log.Println("SUBSCRIBER RECIVE MASSAGE")
	}, stan.DurableName("Test"))
	defer sub.Unsubscribe()	
	go producer.SendRandomData()
	go database.UpdateDb(ch, db, cache)

	//START SERVER SEE DATA ON http://localhost:8080/get_order/1 WHERE 1 is ID of order
	http.HandleFunc("/get_order/", func(w http.ResponseWriter, r *http.Request){ 
		SendDataById(w, r, db, cahe_pointer)})
	err = http.ListenAndServe(":8080", nil)
	helpers.CheckError(err)
}
