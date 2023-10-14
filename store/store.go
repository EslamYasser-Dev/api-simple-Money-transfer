package store

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	model "moneytrans/models"
	"os"
)

var AccountStore = make(map[string]model.Account)

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
		AccountStore[account.ID] = account
	}
	fmt.Println("System is ready to make a transfer.")
}

// this to save changes to the json file
func SaveAccountsToJSON(filepath string) {
	accounts := make([]model.Account, 0, len(AccountStore))
	for _, account := range AccountStore {
		accounts = append(accounts, account)
	}

	jsonData, err := json.Marshal(accounts)
	if err != nil {
		log.Println(err)
		return
	}

	err = ioutil.WriteFile(filepath, jsonData, 0644)
	if err != nil {
		fmt.Println(err)
	}
}
