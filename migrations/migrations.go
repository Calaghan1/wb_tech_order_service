package migrations

import (
	"database/sql"
	"log"
	"github.com/Calaghan1/wb_tech_order_service.git/helpers"
)

func CheckAndSetDb(db *sql.DB) []string{
	var table_name string 
	ArrayOfTables := make([]string, 0, 0)
	transaction, err := db.Begin()
	helpers.CheckError(err)
	rows, err := transaction.Query("SELECT table_name FROM information_schema.tables WHERE table_schema = 'public'")
	helpers.CheckError(err)
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&table_name)
		helpers.CheckError(err)
		ArrayOfTables = append(ArrayOfTables, table_name)
	}
	err = rows.Err()
	helpers.CheckError(err)
log.Println(ArrayOfTables)
	return ArrayOfTables
}

func CreateTables(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS orders (
		id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
		order_uid VARCHAR(50),
		track_number VARCHAR(50),
		entry VARCHAR(50),
		delivery_id UUID,
		payment_id UUID,
		items_id UUID,
		locale VARCHAR(50),
		internal_signature VARCHAR(50),
		customer_id VARCHAR(50),
		delivery_service VARCHAR(50),
		shardkey VARCHAR(50),
		sm_id INT,
		date_created TIMESTAMP,
		oof_shard VARCHAR(10)
	)`
	_, err := db.Exec(query)
	helpers.CheckError(err)
	query = `CREATE TABLE IF NOT EXISTS delivery (
		id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
		name VARCHAR(100),
		phone VARCHAR(15),
		zip VARCHAR(10),
		city VARCHAR(100),
		address VARCHAR(255),
		region VARCHAR(100),
		email VARCHAR(100),
		order_id UUID,
		FOREIGN KEY (order_id) REFERENCES orders(id)
	)`
	_, err = db.Exec(query)
	helpers.CheckError(err)
	query = `CREATE TABLE IF NOT EXISTS payment  (
		id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
		transaction VARCHAR(50),
		request_id VARCHAR(50),
		currency VARCHAR(5),
		provider VARCHAR(50),
		amount INT,
		payment_dt INT,
		bank VARCHAR(50),
		delivery_cost INT,
		goods_total INT,
		custom_fee INT,
		order_id UUID,
		FOREIGN KEY (order_id) REFERENCES orders(id)
	)`
	_, err = db.Exec(query)
	helpers.CheckError(err)
	query = `CREATE TABLE IF NOT EXISTS items (
		id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
		chrt_id INT,
		track_number VARCHAR(50),
		price INT,
		rid VARCHAR(50),
		name VARCHAR(255),
		sale INT,
		size VARCHAR(10),
		total_price INT,
		nm_id INT,
		brand VARCHAR(100),
		status INT,
		order_id UUID,
		FOREIGN KEY (order_id) REFERENCES orders(id)
	);`
	_, err = db.Exec(query)
	helpers.CheckError(err)
}







