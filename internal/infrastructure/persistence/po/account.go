package po

import (
	"fmt"

	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	UserID   uint   `gorm:"uniqueIndex:uni_accounts_customer_currency;not null"`
	User     User   `gorm:"foreignKey:UserID"`
	Amount   int64  `gorm:"type:bigint;not null"`
	Currency string `gorm:"type:varchar(50);uniqueIndex:uni_accounts_customer_currency;not null"`
}

func (a Account) TableName() string {
	return fmt.Sprintf("%s.%s", SCHEMA_PREFIX, "accounts")
}
