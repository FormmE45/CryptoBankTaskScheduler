package entity

import "time"

type AccruedInterest struct {
	ID          string        `gorm:"primaryKey;column=id"`
	AccountID   string        `gorm:"column:account"`
	Account     SavingAccount `gorm:"foreignKey:AccountID"`
	UpdatedDate *time.Time    `gorm:"column:updated_date;default:NOW()"`
	Amount      float64       `gorm:"column:amount"`
	CreatedDate time.Time     `gorm:"column:created_date;default:NOW()"`
}

func (AccruedInterest) TableName() string {
	return "accrued_interest"
}

func NewAccruedInterest(id string, accountID string, amount float64) *AccruedInterest {
	return &AccruedInterest{ID: id, AccountID: accountID, Amount: amount}
}
