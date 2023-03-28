package main

import (
	"go_private_chain/contracts/accountsFactory"
	"go_private_chain/internal/deploy"
	"go_private_chain/utility"
	"log"
)

func main() {
	private := "web3.privateKey0"
	loading, _ := utility.ReadConfigFile([]string{"web3.accountsFactory", private})

	createBox := deploy.LoadWithAddress(loading["web3.accountsFactory"], "accountsFactory", loading[private]).(*accountsFactory.AccountsFactory)

	a, b, c := deploy.InteractiveAccountContract(createBox, "fadsfadfa", loading[private])
	log.Println(a, b, c)

}
