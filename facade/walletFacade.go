package main

import "fmt"

type walletFacade struct {
	account      *account
	wallet       *wallet
	transactions *transactions
}

func newWalletFacade(accountID string, code int) *walletFacade {
	walletFacacde := &walletFacade{
		account:      newAccount(accountID),
		wallet:       newWallet(),
		transactions: &transactions{},
	}
	fmt.Println("account created")
	return walletFacacde
}

func (w *walletFacade) addMoneyToWallet(accountID string, securityCode int, amount int) error {
	err := w.account.validate(accountID)
	if err != nil {
		return err
	}
	w.wallet.add(amount)
	w.transactions.addTransaction(accountID, "deposit", amount)
	return nil
}

func (w *walletFacade) subtractMoneyFromWallet(accountID string, securityCode int, amount int) error {
	err := w.account.validate(accountID)
	if err != nil {
		return err
	}

	err = w.wallet.subtrcact(amount)
	if err != nil {
		return err
	}
	w.transactions.addTransaction(accountID, "withdraw", amount)
	return nil
}
