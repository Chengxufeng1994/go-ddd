package dto

type TransferRequest struct {
	FromAccountId uint  `json:"from_account_id"`
	ToAccountId   uint  `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

type TransferResponse struct {
	FromAccountId uint `json:"from_account_id"`
	ToAccountId   uint `json:"to_account_id"`
}
