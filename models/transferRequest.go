package model

type TransferRequest struct {
	FromAccount string `json:"FromAccount"`
	ToAccount   string `json:"ToAccount"`
	Amount      string `json:"Amount"`
}
