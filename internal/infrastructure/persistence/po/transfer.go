package po

import "gorm.io/gorm"

type Transfer struct {
	gorm.Model
	FromAccountId uint  `gorm:"column:from_account_id"`
	ToAccountId   uint  `gorm:"column:to_account_id"`
	Amount        int64 `gorm:"column:amount"`
}
