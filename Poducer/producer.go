package producer

import (
	"log"
	"os"
	"time"

	"github.com/Calaghan1/wb_tech_order_service.git/helpers"
	"github.com/nats-io/stan.go"
)



func SendRandomData() {
	
	sc, err := stan.Connect(os.Getenv("NATS_STRIMNG_CLASTER"), os.Getenv("NATS_STRIMNG_CLIENT_ID_PROD"))
	helpers.CheckError(err)
	defer sc.Close()
	for true {
		data := helpers.GenerateRandomOrder()
		err = sc.Publish("Orders", data)
		log.Println("Producer sends a message.")
		time.Sleep(time.Second * 1)
	}
}