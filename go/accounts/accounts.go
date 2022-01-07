package accounts

import (
	"errors"
	"fmt"
)

var errNoMoney = errors.New("Cannot withdraw")

// AccountPub struct
type AccountPub struct {
	// public variable: use uppercase
	Owner   string
	Balance int
}

// Account struct
type Account struct {
	owner   string
	balance int
}

// NewAccount creates Account
func NewAccount(owner string) *Account {
	// Create object and return itself using pointer
	account := Account{owner: owner, balance: 0}
	return &account
}

// Balance of your account
func (a Account) Balance() int {
	// To create method, we need to declare "receiver"
	// convention: first letter of structure

	return a.balance
}

// Deposit - add amount on balance
func (a *Account) Deposit(amount int) {
	// If you use Account in receiver it just copying struct
	// If you want change the value, use *Account to get reference

	a.balance += amount
}

// Withdraw - withdraw amount from balance
func (a *Account) Withdraw(amount int) error {
	// There are no exception in go (like try..catch)
	// So, we need to return error and handle it manually
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

// String is method that go call internally when you print it
func (a Account) String() string {
	// We can override it!
	return fmt.Sprint(a.owner, "'s account.\nHas: ", a.balance)
}
