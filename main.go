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
	//Instantiate DAO object
	accountDAO := repository.NewSavingAccountDAO(db)
	accruedInterestDAO := repository.NewAccruedInterestDAO(db)
	//Get Saving Account with Term Preload from database
	s := accountDAO.GetSavingAccountAndPreloadTerm()
	//Calculate on given SavingAccount
	accruedInterestAmount := service.AccruedInterestCalculate(s)
	//Instantiate new AccruedInterest Object
	accruedInterest := entity.NewAccruedInterest("1", s.ID, accruedInterestAmount)
	//Save on DB
	accruedInterestDAO.SaveOnDB(accruedInterest)
}
