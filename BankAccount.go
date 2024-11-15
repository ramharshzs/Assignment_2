package main

import "fmt"

type BankAccount struct {
	accountNumber int
	balance       float64
	name          string
	ifsc          string
}

func display(bank BankAccount) {

	fmt.Println("ACCOUNT NUMBER --> ", bank.accountNumber, "\n BANK NAME --> ", bank.name,
		"\n BALANCE --> ", bank.balance, "\n IFSC CODE --> ", bank.balance)
}
