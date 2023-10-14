package model

type TransferRequest struct {
	FromAccount string `json:"from"`
	ToAccount   string `json:"to"`
	Amount      string `json:"balance"`
}
