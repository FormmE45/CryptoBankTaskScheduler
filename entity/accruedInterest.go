package entity

import (
	"time"

	"github.com/google/uuid"
)

type AccruedInterest struct {
	ID          uuid.UUID     `gorm:"primaryKey;column=id;default:uuid_generate_v4()"`
	AccountID   string        `gorm:"column:account"`
	Account     SavingAccount `gorm:"foreignKey:AccountID"`
	UpdatedDate *time.Time    `gorm:"column:updated_date;default:NOW()"`
	Amount      float64       `gorm:"column:amount"`
	CreatedDate time.Time     `gorm:"column:created_date;default:NOW()"`
	DeletedYN   bool          `gorm:"column:deleted_yn;default:false"`
	DayLeft     uint32        `gorm:"column:day_left"`
}

func (AccruedInterest) TableName() string {
	return "accrued_interest"
}

func NewAccruedInterest(accountID string, amount float64, dayLeft uint32) *AccruedInterest {

	return &AccruedInterest{ID: uuid.New(), AccountID: accountID, Amount: amount, DayLeft: dayLeft}
}
