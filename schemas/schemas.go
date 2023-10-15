package schemas

import (
	"time"
)

type Delivery struct {
	ID      int `gorm:"primaryKey" json:"delivery_id"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Zip     string `json:"zip"`
	City    string `json:"city"`
	Address string `json:"address"`
	Region  string `json:"region"`
	Email   string `json:"email"`
}

type Payment struct {
	ID          int `gorm:"primaryKey" json:"payment_id"`
	Transaction string `json:"transaction"`
	RequestId   string `json:"request_id"`
	Currency    string `json:"currency"`
	Provider    string `json:"provider"`
	Amount      int `json:"amount"`
	PaymentDt   int `json:"payment_dt"`
	Bank        string `json:"bank"`
	DeliveryCost int `json:"delivery_cost"`
	GoodsTotal   int `json:"goods_total"`
	CustomFee   int `json:"custom_fee"`
}

type Item struct {
	ID          int `gorm:"primaryKey" json:"item_id"`
	ChrtID      int `json:"chrt_id"`
	TrackNumber string `json:"track_number"`
	Price       int `json:"price"`
	RID         string `json:"rid"`
	Name        string `json:"name"`
	Sale        int `json:"sale"`
	Size        string `json:"size"`
	TotalPrice  int `json:"total_price"`
	NmID        int `json:"nm_id"`
	Brand       string `json:"brand"`
	Status      int `json:"status"`
}

type Order struct {
	ID              int `gorm:"primaryKey" json:"order_id"`
	OrderUID        string `json:"order_uid"`
	TrackNumber     string `json:"track_number"`
	Entry           string `json:"entry"`
	Delivery        Delivery `gorm:"foreignKey:ID" json:"delivery"`
	Payment         Payment  `gorm:"foreignKey:ID" json:"payment"`
	Items           []Item   `gorm:"foreignKey:ID" json:"items"`
	Locale          string `json:"locale"`
	InternalSignature string `json:"internal_signature"`
	CustomerID      string `json:"customer_id"`
	DeliveryService string `json:"delivery_service"`
	ShardKey        string `json:"shard_key"`
	SmID            int `json:"sm_id"`
	DateCreated     time.Time `json:"date_created"`
	OofShard        string `json:"oof_shard"`
}