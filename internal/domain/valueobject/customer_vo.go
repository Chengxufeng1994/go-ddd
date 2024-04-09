package valueobject

import (
	"errors"
	"fmt"
	"regexp"
)

type CustomerInfo struct {
	age       int
	firstName string
	lastName  string
}

func NewCustomerInfo(age int, firstName, lastName string) *CustomerInfo {
	return &CustomerInfo{
		age:       age,
		firstName: firstName,
		lastName:  lastName,
	}
}

func (ci CustomerInfo) FullName() string {
	return fmt.Sprintf("%s %s", ci.firstName, ci.lastName)
}

func (ci CustomerInfo) FirstName() string {
	return ci.firstName
}

func (ci CustomerInfo) LastName() string {
	return ci.lastName
}

func (ci CustomerInfo) Age() int {
	return ci.age
}

var emailReg = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

type Email struct {
	value string
}

func NewEmail(value string) (*Email, error) {
	if value == "" {
		return &Email{}, errors.New("Email.email cannot be empty")
	}

	e := Email{
		value: value,
	}
	if err := e.validate(); err != nil {
		return &Email{}, err
	}

	return &e, nil
}

func (e Email) String() string {
	return e.value
}

func (e Email) validate() error {
	if !emailReg.MatchString(e.value) {
		return fmt.Errorf("%s is not a valid email", e.value)
	}

	return nil
}

func MustNewEmail(email string) *Email {
	e, err := NewEmail(email)
	if err != nil {
		panic(err)
	}
	return e
}

func (e Email) IsZero() bool {
	return e == Email{}
}

func (e Email) MarshalText() (text []byte, err error) {
	return []byte(e.value), nil
}

func (e *Email) UnmarshalText(text []byte) error {
	e.value = string(text)
	return e.validate()
}
