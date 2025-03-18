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

func (savingAccountDAO SavingAccountDAO) GetSavingAccountAndPreloadTerm() *entity.SavingAccount {
	var saving_account entity.SavingAccount
	savingAccountDAO.Db.Preload("Term").First(&saving_account)
	return &saving_account
}
