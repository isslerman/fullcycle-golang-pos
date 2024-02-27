package main

import (
	"errors"
	"fmt"
)

type NuBank struct {
	balance int
}

func NewNuBank() *NuBank {
	return &NuBank{
		balance: 0,
	}

}

func (n *NuBank) GetBalance() int {
	return n.balance
}

func (n *NuBank) Deposit(amount int) {
	fmt.Printf("New deposit at Nubank: %d\n", amount)
	n.balance += amount
}

func (n *NuBank) WithDraw(amount int) error {
	newBalance := n.balance - amount
	if newBalance < 0 {
		return errors.New("insufficient funds")
	}

	n.balance = newBalance
	return nil
}
