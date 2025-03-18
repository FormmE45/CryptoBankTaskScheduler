package repository

import (
	"github.com/FormmE45/CryptoBankTaskScheduler/entity"
	"gorm.io/gorm"
)

type AccruedInterestDAO struct {
	DB *gorm.DB
}

func NewAccruedInterestDAO(db *gorm.DB) *AccruedInterestDAO {
	return &AccruedInterestDAO{DB: db}
}

func (dao AccruedInterestDAO) SaveOnDB(accruedInterest *entity.AccruedInterest) {
	dao.DB.Create(&accruedInterest)
}
