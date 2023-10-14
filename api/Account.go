package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	model "moneytrans/models"
	store "moneytrans/store"
)

func ListAccounts(w http.ResponseWriter, r *http.Request) {
	var accounts []model.Account
	for _, account := range store.AccountStore {
		accounts = append(accounts, account)
	}
	json.NewEncoder(w).Encode(accounts)
}

func Transfer(w http.ResponseWriter, r *http.Request) {
	var request model.TransferRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fromAccount, ok := store.AccountStore[request.FromAccount]
	if !ok {
		http.Error(w, "Sender account is not found", http.StatusNotFound)
		return
	}

	toAccount, ok := store.AccountStore[request.ToAccount]
	if !ok {
		http.Error(w, "the Receiver data is wrong", http.StatusNotFound)
		return
	}

	fromBalance, _ := strconv.Atoi(fromAccount.Balance)
	toBalance, _ := strconv.Atoi(toAccount.Balance)
	requestAmount, _ := strconv.Atoi(request.Amount)

	if fromBalance < requestAmount {
		http.Error(w, "Insufficient balance", http.StatusForbidden)
		return
	}

	fromBalance -= requestAmount
	toBalance += requestAmount

	fromAccount.Balance = strconv.Itoa(fromBalance)
	toAccount.Balance = strconv.Itoa(toBalance)

	store.AccountStore[request.FromAccount] = fromAccount
	store.AccountStore[request.ToAccount] = toAccount
	store.SaveAccountsToJSON("./accounts.json")

	fmt.Fprintf(w, "Transfer successful: %s \t from %s to %s\n", request.Amount, request.FromAccount, request.ToAccount)
}
