package main

import (
	"errors"
	"fmt"
)

//Bitcoin type
type Bitcoin int

//Wallet for storing Bitcoins
type Wallet struct {
	balance Bitcoin
}

//Stringer formats Bitcoin string
type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

//Balance returns the balance of your Bitcoin wallet
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

//Deposit into your Bitcoin wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

//ErrInsufficientFunds is an error for lack of funds
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

//Withdraw from your Bitcoin wallet
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}
