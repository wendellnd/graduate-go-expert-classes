package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Account struct {
	Number  int `json:"n"`
	Balance int `json:"b"`
}

func main() {
	account := Account{Number: 1, Balance: 2000}

	res, err := json.Marshal(account)
	if err != nil {
		panic(err)
	}
	println(string(res))

	err = json.NewEncoder(os.Stdout).Encode(account)
	if err != nil {
		panic(err)
	}

	rawJson := []byte(`{"n": 2, "b": 2000}`)
	var accountX Account
	err = json.Unmarshal(rawJson, &accountX)
	if err != nil {
		panic(err)
	}
	fmt.Println(accountX.Balance)
}
