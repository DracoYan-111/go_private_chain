package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func main() {
	keyJSON, err := ioutil.ReadFile("./account.json")
	if err != nil {
		log.Fatal(err)
	}

	password := "abc123"

	key, err := keystore.DecryptKey(keyJSON, password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(key.Address.Hex())
}
