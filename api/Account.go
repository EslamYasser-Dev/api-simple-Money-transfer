package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	model "moneytrans/models"
	store "moneytrans/store"
)

// to get the data from json file
func LoadAccountsFromJSON(filepath string) {
	jsonFile, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var accounts []model.Account
	json.Unmarshal(byteValue, &accounts)

	for _, account := range accounts {
		store.AccountStore[account.ID] = account
	}

	fmt.Println("System is ready to make a transfer.")
}

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

	//validate the exitance of sender and receiver
	fromAccount, ok := store.AccountStore[request.FromAccount]
	if !ok {
		http.Error(w, "Sender account is not found", http.StatusNotFound)
		return
	}

	toAccount, ok := store.AccountStore[request.ToAccount]
	if !ok {
		http.Error(w, "the Receiver data is worng", http.StatusNotFound)
		return
	}

	//i did it	to avoid floating point while we were calculating the balance
	fromBalance, _ := strconv.Atoi(fromAccount.Balance)
	toBalance, _ := strconv.Atoi(toAccount.Balance)
	requestAmount, _ := strconv.Atoi(request.Amount)

	//this to validate the balance >>> fromBalance === sender balance
	if fromBalance < requestAmount {
		http.Error(w, "Insufficient balance", http.StatusForbidden)
		return
	}

	fromBalance -= requestAmount
	toBalance += requestAmount

	/*convert integers to floats in a string format like this: "16551.54"
	onTransaction(stirng >>> float32 >>> uint)  ====== onStoring the balance (uint >> float32 >>> string)

	*/
	fromAccount.Balance = strconv.Itoa(fromBalance)
	toAccount.Balance = strconv.Itoa(toBalance)

	store.AccountStore[request.FromAccount] = fromAccount
	store.AccountStore[request.ToAccount] = toAccount

	fmt.Fprintf(w, "Transfer successful: %s from %s to %s\n", request.Amount, request.FromAccount, request.ToAccount)
	print("thanks for your trust")
}
