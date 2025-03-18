package entity

type Term struct {
	ID          string `gorm:"column:id;primaryKey"`
	AmountMonth int64
	Type        string
}

func (Term) TableName() string {
	return "term"
}
