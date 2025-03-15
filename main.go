package main

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/lib/pq"

	service "github.com/FormmE45/CryptoBankTaskScheduler/service"
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
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	//Ping DB
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully Connected")

}
