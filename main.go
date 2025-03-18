package main

import (
	"github.com/FormmE45/CryptoBankTaskScheduler/repository"
	service "github.com/FormmE45/CryptoBankTaskScheduler/service"
	_ "github.com/lib/pq"
)

type Term struct {
	ID          string `gorm:"column:id;primaryKey"`
	AmountMonth int64
	Type        string
}

func (Term) TableName() string {
	return "term"
}

func main() {
	//Open DB connection
	db := repository.OpenDBConnection()
	Repository := repository.Repository{Db: db}
	//Get sqlDB
	sqlDB, err := db.DB()
	service.Check(err)
	defer sqlDB.Close()

}
