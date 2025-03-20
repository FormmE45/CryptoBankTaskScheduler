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
	// 	dao.DB.Where(entity.AccruedInterest{AccountID: accruedInterest.AccountID}).FirstOrCreate(&accruedInterest)
	// 	dao.DB.Clauses(clause.OnConflict{
	// 		Columns: []clause.Column{{Name: "account"}},
	// 		DoUpdates: clause.Assignments(map[string]any{
	// 			"amount": clause.Expr{SQL: "accrued_interest.amount + EXCLUDED.amount", Vars: []interface{}{}},
	// 		}),
	// 	}).Create(&accruedInterest)
	//SQL approach

	dao.DB.Exec(`INSERT INTO accrued_interest (id,account,updated_date,amount,deleted_yn,day_left)
				VALUES (?,?,?,?,?,?)
				ON CONFLICT (account) WHERE deleted_yn=false
				DO UPDATE SET amount=accrued_interest.amount+ EXCLUDED.amount,day_left=accrued_interest.day_left-1`,
		accruedInterest.ID, accruedInterest.AccountID, accruedInterest.UpdatedDate, accruedInterest.Amount, accruedInterest.DeletedYN, accruedInterest.DayLeft)
	//code approach
}

func (dao AccruedInterestDAO) CheckAccruedInterestDayLeftIfZero() {
	var matchingZeroDay []entity.AccruedInterest
	dao.DB.Model(&entity.AccruedInterest{}).Where("day_left = ?", 0).Preload("Account").Find(&matchingZeroDay)
	for _, accruedInterest := range matchingZeroDay {
		dao.DB.Model(&accruedInterest.Account).Update("status_id", "1")
		dao.DB.Model(&accruedInterest).Update("deleted_yn", true)
	}
}
