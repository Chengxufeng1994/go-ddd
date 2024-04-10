package valueobject

import (
	"errors"
)

var (
	ErrMustGreaterThanZero = errors.New("amount must greater than zero")
	ErrCurrencyNotEqual    = errors.New("currencies must be identical")
	ErrAmountNotEnough     = errors.New("there is not enough amount to deduct")
)

type Money struct {
	amount   int64
	currency string
}

func NewMoney(amount int64, currency string) (*Money, error) {
	if amount < 0 {
		return nil, ErrMustGreaterThanZero
	}
	return &Money{amount: amount, currency: currency}, nil
}

func (m *Money) Amount() int64 {
	return m.amount
}

func (m *Money) Currency() string {
	return m.currency
}

func (m *Money) IsCurrencyEquals(other Money) bool {
	return m.currency == other.currency
}

func (m Money) Add(other Money) (*Money, error) {
	if !(m.currency == other.currency) {
		return nil, ErrCurrencyNotEqual
	}

	return &Money{
		amount:   m.amount + other.amount,
		currency: m.currency,
	}, nil
}

func (m Money) Deduct(other Money) (*Money, error) {
	if !(m.currency == other.currency) {
		return nil, ErrCurrencyNotEqual
	}

	if other.amount > m.amount {
		return nil, ErrAmountNotEnough
	}

	return &Money{
		amount:   m.amount - other.amount,
		currency: m.currency,
	}, nil
}
