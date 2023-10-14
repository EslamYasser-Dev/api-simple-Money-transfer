package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"

	model "moneytrans/models"
	store "moneytrans/store"
)

var mutex = &sync.Mutex{}

func ListAccounts(w http.ResponseWriter, r *http.Request) {
	var accounts []model.Account
	for _, account := range store.AccountStore {
		accounts = append(accounts, account)
	}
	json.NewEncoder(w).Encode(accounts)
}

func Transfer(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	var request model.TransferRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fromAccount, ok := store.AccountStore[request.FromAccount]
	if !ok {
		http.Error(w, "Sender ID is wrong", http.StatusNotFound)
		return
	}

	toAccount, ok := store.AccountStore[request.ToAccount]
	if !ok {
		http.Error(w, "The Receiver ID is wrong", http.StatusNotFound)
		return
	}

	fromBalance, err := strconv.ParseFloat(fromAccount.Balance, 64)
	if err != nil {
		http.Error(w, "Invalid balance format for sender", http.StatusBadRequest)
		return
	}

	toBalance, err := strconv.ParseFloat(toAccount.Balance, 64)
	if err != nil {
		http.Error(w, "Invalid balance format for receiver", http.StatusBadRequest)
		return
	}

	requestAmount, err := strconv.ParseFloat(request.Amount, 64)
	if err != nil {
		http.Error(w, "Invalid transfer amount format", http.StatusBadRequest)
		return
	}

	if fromBalance < requestAmount {
		http.Error(w, "Insufficient balance", http.StatusBadRequest)
		return
	}

	fromBalance -= requestAmount
	toBalance += requestAmount
	if fromBalance < 0 || requestAmount < 0 {
		http.Error(w, "Error in calculations", http.StatusBadRequest)
		return
	}

	fromAccount.Balance = fmt.Sprintf("%.2f", fromBalance)
	toAccount.Balance = fmt.Sprintf("%.2f", toBalance)

	store.AccountStore[request.FromAccount] = fromAccount
	store.AccountStore[request.ToAccount] = toAccount
	store.SaveAccountsToJSON("../accounts.json") //save changes

	fmt.Fprintf(w, "\n Transfer Operation has been successfully done\n  %s $ has been tranfered \n from: %s => id: (%s) \n to: %s => id: (%s)\n ============================================================",
		request.Amount,
		fromAccount.Name, request.FromAccount,
		toAccount.Name, request.ToAccount)
	// json.NewEncoder(w).Encode(w)

}
