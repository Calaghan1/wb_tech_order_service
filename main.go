package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"

	producer "github.com/Calaghan1/wb_tech_order_service.git/Poducer"
	"github.com/Calaghan1/wb_tech_order_service.git/database"
	"github.com/Calaghan1/wb_tech_order_service.git/helpers"
	"github.com/Calaghan1/wb_tech_order_service.git/schemas"
	"github.com/joho/godotenv"
	"github.com/nats-io/stan.go"
	"gorm.io/gorm"
	// "github.com/Calaghan1/wb_tech_order_service.git/migrations"
	// "github.com/nats-io/stan.go"
	// "github.com/nats-io/nats.go"
)

func SendDataById(w http.ResponseWriter, r *http.Request, db *gorm.DB, cache *map[int]schemas.Order, mutex *sync.Mutex) {
	w.Header().Set("Content-Type", "application/json")
	idStr := r.URL.Path[len("/increment/"):]
	id, err := strconv.Atoi(idStr)
	helpers.CheckError(err)
	mutex.Lock()
	order, ok := (*cache)[id]
	if ok {
	response, err := json.Marshal(order)
	helpers.CheckError(err)
	w.Write(response)
	} else {
		json.NewEncoder(w).Encode(struct{}{})
	}
	mutex.Unlock()
}

func main() {
	err := godotenv.Load()
	helpers.CheckError(err)
	var mutex sync.Mutex
	//INIT DATABASE AND CACHE
	fmt.Println(os.Getenv("POSTRGRES_CONNECTION"))
	db := database.Init_database(os.Getenv("POSTRGRES_CONNECTION"))
	cache := database.MakeCache(db)
	cahe_pointer := &cache
	//INIT SUBSCRIBER AND RODUCER NATS STRIMING
	sc, err := stan.Connect(os.Getenv("NATS_STRIMNG_CLASTER"), os.Getenv("NATS_STRIMNG_CLIENT_ID_SUB"))
	helpers.CheckError(err)
	defer sc.Close()
	ch := make(chan []byte, 10)
	sub, _ := sc.Subscribe("Orders", func(m *stan.Msg,) {
		ch <- m.Data
		log.Println("SUBSCRIBER RECIVE MASSAGE")
	}, stan.DurableName("Test"))
	defer sub.Unsubscribe()	
	go producer.SendRandomData()
	go database.UpdateDb(ch, db, cache, &mutex)

	//START SERVER SEE DATA ON http://localhost:8080/get_order/1 WHERE 1 is ID of order
	http.HandleFunc("/get_order/", func(w http.ResponseWriter, r *http.Request){ 
		SendDataById(w, r, db, cahe_pointer, &mutex)})
	err = http.ListenAndServe("0.0.0.0:8080", nil)
	helpers.CheckError(err)
}
