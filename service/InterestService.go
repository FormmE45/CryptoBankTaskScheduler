package service

import (
	"fmt"

	"github.com/FormmE45/CryptoBankTaskScheduler/entity"
)

func AccruedInterestCalculate(s *entity.SavingAccount) float64 {
	balance := s.Balance
	fmt.Println(balance)
	interestRate := s.InterestRate
	accruedInterest := (balance * (float64(interestRate))) / 100
	fmt.Println(accruedInterest)
	return accruedInterest
}
