package producer

import (
	"encoding/json"
	"log"
	"time"

	"github.com/Calaghan1/wb_tech_order_service.git/helpers"
	"github.com/nats-io/stan.go"
)



func SendRandomData() {
	sc, err := stan.Connect("test-cluster", "client-1")
	helpers.CheckError(err)
	defer sc.Close()
	for true {
		order := helpers.GenerateRandomOrder()
		data, err := json.MarshalIndent(order, "", " ")
		helpers.CheckError(err)
		err = sc.Publish("Orders", data)
		time.Sleep(5 * time.Second)
		log.Println("PRIDUCER SEND MASSAGE")
	}
}