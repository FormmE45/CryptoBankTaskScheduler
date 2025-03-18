package repository

import (
	"fmt"

	"github.com/FormmE45/CryptoBankTaskScheduler/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const DBConfigPathfileServerMain = "E:\\DBconfig\\CryptoBankDBConfigMain.txt"
const DBConfigPathfileServerTestLaptop = "C:\\GoLang\\DBconfig\\CryptoBankDBConfigTest.txt"
const DBConfigPathfileServerTest = "E:\\DBconfig\\CryptoBankDBConfigTest.txt"

type Repository struct {
	Db *gorm.DB
}

func OpenDBConnection() *gorm.DB {
	//Init DBConfig file from file path
	dbConfig := service.InitDBConfig(DBConfigPathfileServerTestLaptop)
	//Open a connection to DB using GORM
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.Dbname)
	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	service.Check(err)
	return db
}

func GetFirstRow(db *gorm.DB, object any) {
	result := db.First(&object)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
}
