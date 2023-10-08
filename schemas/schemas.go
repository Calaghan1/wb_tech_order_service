package schemas

import (
	"time"

	"github.com/google/uuid"
)
type Order struct {
	Id uuid.UUID
	Order_uid string
	Track_number string
	Entry string
	Delivery Delivery
	Payment Payment
	Items []Items
	Locale string
	Internal_signature string
	Customer_id string
	Delivery_service string
	Shardkey string
	Sm_id int
	Date_created time.Time
	Oof_shard string
}

type Delivery struct {
	name string
	phone string
	zip string
	city string
	address string
	region string
	email string
}

type Payment struct {
	Transaction string
	Request_id 	string
	Currency	string
	Provider	string
	Amount		int
	Payment_dt 	int 
	Bank		string
	Delivery_cost	int
	Goods_total	int
	Custom_fee int
}

type Items struct {
	Chrt_id			int
	Track_number	string
	Price 			int
	Rid				string
	Name 			string
	Sale			int
	Size			string
	Total_price		int
	Nm_id 			string
	Brand			string
	Status			int
}
sdsd