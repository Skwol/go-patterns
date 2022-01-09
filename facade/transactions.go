package main

import "fmt"

type transactions struct{}

func (s *transactions) addTransaction(accountID, txType string, amount int) {
	fmt.Printf("adding transaction for accountId %s with type %s for amount %d\n", accountID, txType, amount)
	return
}
