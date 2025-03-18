package entity

type Term struct {
	ID           string  `gorm:"column:id;primaryKey"`
	AmountMonth  int64   `gorm:"column:amount_month"`
	Type         string  `gorm:"column:type"`
	InterestRate float64 `gorm:"column:interest_rate_of_month"`
}

func (Term) TableName() string {
	return "term"
}
