package model

type Account struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Balance string `json:"balance"`
	//you can add or change ex.: token, currency, status, date joined,etc
}
