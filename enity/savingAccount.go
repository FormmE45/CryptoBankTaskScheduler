package entity

import "time"

type SavingAccount struct {
	id           string
	userId       string
	termId       string
	note         string
	statusId     string
	deleteYN     bool
	createdDate  time.Time
	createBy     string
	modifiedDate time.Time
	modifiedBy   string
	balance      float64
	interestRate float32
	maturityDate time.Time
	ggDriveURL   string
	heirStatus   bool
	name         string
	uuid         uint
}
