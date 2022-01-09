package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println()
	walletFacade := newWalletFacade("account_id", 1234)
	fmt.Println()

	err := walletFacade.addMoneyToWallet("account_id", 1234, 10)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}

	fmt.Println()
	err = walletFacade.subtractMoneyFromWallet("account_id", 1234, 5)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
}
