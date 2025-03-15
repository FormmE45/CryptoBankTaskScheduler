package main

import (
	"fmt"
	"strconv"

	SavingAccount "github.com/FormmE45/CryptoBankTaskScheduler/entity"
	service "github.com/FormmE45/CryptoBankTaskScheduler/service"
	_ "github.com/lib/pq"
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

	//Open a connection to DB
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbConfig.host, dbConfig.port, dbConfig.user, dbConfig.password, dbConfig.dbname)
	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	service.Check(err)

	sqlDB, err := db.DB()
	service.Check(err)

	defer sqlDB.Close()

	//Ping DB
	err = sqlDB.Ping()
	if err != nil {
		panic(err)
	}

	var saving_account SavingAccount

	fmt.Println("Successfully Connected")

}
