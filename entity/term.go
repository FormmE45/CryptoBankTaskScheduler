package entity

type Term struct {
	ID          string `gorm:"primaryKey;column:id"`
	AmountMonth int64
	Type        string
}

func (Term) TableName() string {
	return "term"
}
