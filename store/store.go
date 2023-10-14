package store

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
