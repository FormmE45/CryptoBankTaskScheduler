package entity

import (
	"time"

	"gorm.io/gorm"
)

type SavingAccount struct {
	id           string    `gorm:"primaryKey;column:id"`
	userId       string    `gorm:"column:user_id"`
	termId       string    `gorm:"column:term_id"`
	note         string    `gorm:"column:note"`
	statusId     string    `gorm:"column:status_id"`
	deleteYN     bool      `gorm:"column:delete_yn"`
	createdDate  time.Time `gorm:"column:created_date"`
	createBy     string    `gorm:"column:created_by"`
	modifiedDate time.Time `gorm:"column:modified_date"`
	modifiedBy   string    `gorm:"column:modified_by"`
	balance      float64   `gorm:"column:balance"`
	interestRate float32   `gorm:"column:interest_rate"`
	maturityDate time.Time `gorm:"column:maturity_date"`
	ggDriveURL   string    `gorm:"column:gg_drive_url"`
	heirStatus   bool      `gorm:"column:heir_status"`
	name         string    `gorm:"column:name"`
	uuid         uint      `gorm:"column:uuid_id"`
}

func (s SavingAccount) TableName() string {
	return "saving_account"
}

func (s SavingAccount) GetFirstRow(db *gorm.DB) *SavingAccount {
	var saving_account SavingAccount
	db.First(&saving_account)
	return &saving_account
}
