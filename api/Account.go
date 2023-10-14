package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	model "moneytrans/models"
	store "moneytrans/store"
)

// ListAccounts is an HTTP handler function that lists all accounts.
// It responds to GET requests at the /accounts endpoint.
func ListAccounts(w http.ResponseWriter, r *http.Request) {
	var accounts []model.Account
	// Loop over all accounts in the store and append them to the accounts slice
	for _, account := range store.AccountStore {
		accounts = append(accounts, account)
	}
	// Encode the accounts slice to JSON and send it as a response
	json.NewEncoder(w).Encode(accounts)
}

// Transfer is an HTTP handler function that transfers money from one account to another.
// It responds to POST requests at the /transfer endpoint.
func Transfer(w http.ResponseWriter, r *http.Request) {
	var request model.TransferRequest
	// Decode the JSON request body into the request variable
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate the existence of sender and receiver
	fromAccount, ok := store.AccountStore[request.FromAccount]
	if !ok {
		http.Error(w, "Sender account is not found", http.StatusNotFound)
		return
	}

	toAccount, ok := store.AccountStore[request.ToAccount]
	if !ok {
		http.Error(w, "The Receiver data is wrong", http.StatusNotFound)
		return
	}

	// Convert balances from string to integer to avoid floating point precision issues
	fromBalance, _ := strconv.Atoi(fromAccount.Balance)
	toBalance, _ := strconv.Atoi(toAccount.Balance)
	requestAmount, _ := strconv.Atoi(request.Amount)

	// Check if sender has enough balance for the transfer
	if fromBalance < requestAmount {
		http.Error(w, "Insufficient balance", http.StatusForbidden)
		return
	}

	fromBalance -= requestAmount
	toBalance += requestAmount

	// Convert balances back to string format after performing the transfer
	fromAccount.Balance = strconv.Itoa(fromBalance)
	toAccount.Balance = strconv.Itoa(toBalance)

	store.AccountStore[request.FromAccount] = fromAccount
	store.AccountStore[request.ToAccount] = toAccount

	fmt.Fprintf(w, "Transfer successful: %s  from %s (%s) \n to %s (%s)\n",
		request.Amount,
		fromAccount.Name, request.FromAccount,
		toAccount.Name, request.ToAccount)
}
