package main

import (
	"log"
	api "moneytrans/api"
	store "moneytrans/store"
	"net/http"
)

func main() {
	const accountsFilePath = "./accounts.json"
	store.LoadAccountsFromJSON(accountsFilePath)
	http.HandleFunc("/accounts", api.ListAccounts)
	http.HandleFunc("/transfer", api.Transfer)

	print("\t -= welcome to money transfer REST API =-\n\t ***********************************************\n \t  - [POST] to transfer money use: http://localhost:8570/transfer  \n")
	print("\n \t  - [GET] to get all accounts use: http://localhost:8570/accounts  \n\t ***********************************************\n for support call +201062700575")

	log.Fatal(http.ListenAndServe(":8570", nil))
}
