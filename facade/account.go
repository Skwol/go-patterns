package main

import "fmt"

type account struct {
	name string
}

func newAccount(name string) *account {
	return &account{
		name: name,
	}
}

func (a *account) validate(accountName string) error {
	if a.name != accountName {
		return fmt.Errorf("name is incorrect")
	}
	fmt.Println("account validated")
	return nil
}
