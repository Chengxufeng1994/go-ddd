package po

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	CustomerID uint     `gorm:"uniqueIndex:uni_accounts_customer_currency;not null"`
	Customer   Customer `gorm:"foreignKey:CustomerID"`
	Amount     int64    `gorm:"type:bigint;not null"`
	Currency   string   `gorm:"type:varchar(50);uniqueIndex:uni_accounts_customer_currency;not null"`
}
