package main

import (
	"fmt"

	"github.com/ttochi/tutorials/go/accounts"
)

func structMain() {
	// 1. To export structure, we need to use uppercase (member variables too)
	accountPub := accounts.AccountPub{Owner: "ttochi", Balance: 1000}
	fmt.Println(accountPub)

	// 2. Constructor on struct
	account := accounts.NewAccount("ttochi")

	// 3. Method
	account.Deposit(10)
	fmt.Println(account.Balance())

	// 4. Error handling
	account.Withdraw(20)
	fmt.Println(account.Balance()) // nothing happen - go dose not handle error

	err := account.Withdraw(20)
	if err != nil { // we need to handle error manually
		fmt.Println(err) // instead, log.Fatalln(err)
	}
	fmt.Println(account.Balance())

	// 5. Override internal method
	fmt.Println(account)
}
