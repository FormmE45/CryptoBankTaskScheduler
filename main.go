package main

import (
	"fmt"
	"time"

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
	saving_accounts := accountDAO.GetAllSavingAccountAndPreloadTerm()
	//Calculate on given SavingAccount
	for _, account := range *saving_accounts {
		accruedInterestAmount := service.AccruedInterestCalculate(&account)
		startDate := account.CreatedDate
		fmt.Println(startDate)
		fmt.Println(account.Term.AmountMonth)
		endDate := account.CreatedDate.AddDate(0, int(account.Term.AmountMonth), 0)
		fmt.Println(endDate)
		dayLeft := fromDateToDate(startDate, endDate)
		//Instantiate new AccruedInterest Object
		accruedInterest := entity.NewAccruedInterest(account.ID, accruedInterestAmount, dayLeft)
		//Save on DB
		accruedInterestDAO.SaveOnDB(accruedInterest)
	}
	managedIfDueDate(*accruedInterestDAO)
}

func fromDateToDate(start time.Time, end time.Time) uint32 {
	days := uint32(end.Sub(start).Hours() / 24)
	return days
}

func managedIfDueDate(accruedInterestDAO repository.AccruedInterestDAO) {
	accruedInterestDAO.CheckAccruedInterestDayLeftIfZero()
}
