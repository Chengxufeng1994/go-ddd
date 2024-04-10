package po

import "gorm.io/gorm"

type Transfer struct {
	gorm.Model
	FromAccountId uint  `json:"from_account_id"`
	ToAccountId   uint  `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}
