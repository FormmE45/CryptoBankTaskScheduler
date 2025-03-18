package main

import (
	"github.com/FormmE45/CryptoBankTaskScheduler/entity"
	_ "github.com/FormmE45/CryptoBankTaskScheduler/entity"
	"github.com/FormmE45/CryptoBankTaskScheduler/repository"
	service "github.com/FormmE45/CryptoBankTaskScheduler/service"
	_ "github.com/lib/pq"
)

func main() {
	//Open DB connection
	db := repository.OpenDBConnection()
	// Repository := repository.Repository{Db: db}
	//Get sqlDB
	sqlDB, err := db.DB()
	service.Check(err)
	defer sqlDB.Close()

	accountDAO := repository.NewSavingAccountDAO(db)
	accruedInterestDAO := repository.NewAccruedInterestDAO(db)
	s := accountDAO.GetSavingAccountAndPreloadTerm()
	accruedInterestAmount := service.AccruedInterestCalculate(s)
	accruedInterest := entity.NewAccruedInterest("3", s.ID, accruedInterestAmount)
	accruedInterestDAO.SaveOnDB(accruedInterest)
	// var term entity.Term
	// db.First(&term)
	// fmt.Println(term)
}
