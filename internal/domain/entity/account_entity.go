package entity

import (
	"time"

	"github.com/Chengxufeng1994/go-ddd/internal/domain/valueobject"
)

type Account struct {
	ID         uint
	CustomerID uint
	Money      *valueobject.Money
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Accounts []Account

func (a *Account) Add(other valueobject.Money) error {
	result, err := a.Money.Add(other)
	if err != nil {
		return err
	}

	a.Money = result

	return nil
}

func (a *Account) Deduct(other valueobject.Money) error {
	result, err := a.Money.Deduct(other)
	if err != nil {
		return err
	}

	a.Money = result

	return nil
}
