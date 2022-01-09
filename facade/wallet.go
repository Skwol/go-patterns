package main

import "fmt"

type wallet struct {
	balance int
}

func newWallet() *wallet {
	return &wallet{
		balance: 0,
	}
}

func (w *wallet) add(amount int) {
	w.balance += amount
	fmt.Println("adding money to wallet")
}

func (w *wallet) subtrcact(amount int) error {
	if w.balance < amount {
		return fmt.Errorf("not enough money to make a transaction")
	}
	fmt.Println("substructing money from wallet")
	w.balance -= -amount
	return nil
}
