package entity

import (
	"time"
)

type SavingAccount struct {
	ID           string    `gorm:"primaryKey;column:id"`
	UserId       string    `gorm:"column:user_id"`
	TermId       string    `gorm:"column:term_id"`
	Note         string    `gorm:"column:note"`
	StatusId     string    `gorm:"column:status_id"`
	DeleteYN     bool      `gorm:"column:delete_yn"`
	CreatedDate  time.Time `gorm:"column:created_date"`
	CreateBy     string    `gorm:"column:created_by"`
	ModifiedDate time.Time `gorm:"column:modified_date"`
	ModifiedBy   string    `gorm:"column:modified_by"`
	Balance      float64   `gorm:"column:balance"`
	InterestRate float32   `gorm:"column:interest_rate"`
	MaturityDate time.Time `gorm:"column:maturity_date"`
	GGDriveURL   string    `gorm:"column:gg_drive_url"`
	HeirStatus   bool      `gorm:"column:heir_status"`
	Name         string    `gorm:"column:name"`
	UUID         string    `gorm:"column:uuid_id"`
}

func (s SavingAccount) TableName() string {
	return "saving_account"
}
