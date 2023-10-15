package helpers

import (
	"fmt"
	"log"
	"math/rand"
	"time"
	"github.com/Calaghan1/wb_tech_order_service.git/schemas"
)

func CheckError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}


func GenerateRandomOrder() schemas.Order {
	return schemas.Order{
		OrderUID:    generateRandomString(19),
		TrackNumber: generateRandomString(10),
		Entry:       generateRandomString(5),
		Locale:      generateRandomString(2),
		CustomerID:  generateRandomString(10),
		Delivery: schemas.Delivery{
			Name:    generateRandomString(10),
			Phone:   generateRandomPhoneNumber(),
			Zip:     generateRandomString(7),
			City:    generateRandomString(15),
			Address: generateRandomString(20),
			Region:  generateRandomString(10),
			Email:   generateRandomEmail(),
		},
		Payment: schemas.Payment{
			Transaction:  generateRandomString(21),
			Currency:      generateRandomString(3),
			Provider:      generateRandomString(5),
			Amount:        rand.Intn(10000),
			PaymentDt:     int(time.Now().Unix()),
			Bank:          generateRandomString(5),
			DeliveryCost:  rand.Intn(2000),
			GoodsTotal:    rand.Intn(5000),
			CustomFee:     rand.Intn(100),
		},
		Items: []schemas.Item{
			{
				ChrtID:     rand.Intn(100000),
				TrackNumber: generateRandomString(10),
				Price:      rand.Intn(1000),
				RID:       generateRandomString(21),
				Name:       generateRandomString(10),
				Sale:       rand.Intn(50),
				Size:       generateRandomString(1),
				TotalPrice: rand.Intn(1000),
				NmID:       rand.Intn(1000000),
				Brand:      generateRandomString(10),
				Status:     rand.Intn(300),
			},
		},
		DeliveryService: generateRandomString(5),
		ShardKey:        generateRandomString(1),
		SmID:            rand.Intn(100),
		DateCreated:     time.Now(),
		OofShard:        generateRandomString(1),
	}
}

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func generateRandomEmail() string {
	return fmt.Sprintf("%s@example.com", generateRandomString(10))
}


func generateRandomPhoneNumber() string {
	return fmt.Sprintf("+%d", rand.Intn(1000000000))
}


