package repository

import (
	"github.com/FormmE45/CryptoBankTaskScheduler/entity"
	"gorm.io/gorm"
)

type SavingAccountDAO struct {
	Db *gorm.DB
}

func NewSavingAccountDAO(db *gorm.DB) *SavingAccountDAO {
	return &SavingAccountDAO{Db: db}
}

func (savingAccountDAO SavingAccountDAO) GetAllSavingAccountAndPreloadTerm() *[]entity.SavingAccount {
	var saving_accounts []entity.SavingAccount
	savingAccountDAO.Db.Preload("Term").Find(&saving_accounts)
	return &saving_accounts
}
