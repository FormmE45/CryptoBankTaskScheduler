package main

import (
	"fmt"
	"strconv"

	service "github.com/FormmE45/CryptoBankTaskScheduler/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBconfig struct {
	host     string
	port     uint64
	user     string
	password string
	dbname   string
}

func main() {

	//Read file that hold DB configuration
	file := service.OpenFile("E:\\DBconfig\\CryptoBankDBConfig.txt")
	info := service.GetDBConfig(file)
	u, err := strconv.ParseUint(info[1].(string), 10, 64)

	service.Check(err)
	//Instantiate DBConfig
	dbConfig := DBconfig{info[0].(string), u, info[2].(string), info[3].(string), info[4].(string)}
	fmt.Println(dbConfig)
	//Open a connection to DB
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbConfig.host, dbConfig.port, dbConfig.user, dbConfig.password, dbConfig.dbname)
	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	service.Check(err)

	// Ping DB
	if err != nil {
		panic(err)
	}
	// Check if ur in the correct database
	// var tables []string
	// db.Raw("SELECT table_name FROM information_schema.tables WHERE table_schema='public'").Scan(&tables)
	// fmt.Println("Tables in database:", tables)

	// var rawID string
	// err = db.Table("users").First("SELECT id FROM users LIMIT 1").Error
	// if err != nil {
	// 	fmt.Println("Raw SQL error:", err)
	// } else {
	// 	fmt.Println("Raw SQL found ID:", rawID)
	// }

	//Count if there is any data in the table mapping to the struct
	// var count int64
	// db.Model(&entity.Term{}).Count(&count)
	// fmt.Println("Total records:", count)

	// var term entity.Term
	// db.Raw("SELECT id,amount_month,type FROM term WHERE id= ? ", "term001").Scan(&term)
	// fmt.Println(term)

	// db.Debug().AutoMigrate(&entity.Term{})

	var schemaName string
	db.Raw("SELECT current_database()").Scan(&schemaName)
	fmt.Println("Current Schema:", schemaName)
	// var saving_account entity.SavingAccount
	// saving_account.GetFirstRow(db, saving_account)
	// fmt.Println("")

}
