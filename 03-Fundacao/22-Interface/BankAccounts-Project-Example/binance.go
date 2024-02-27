package main

import (
	"errors"
	"fmt"
)

type Binance struct {
	balance int
	fee     int
}

func NewBinance() *Binance {
	return &Binance{
		balance: 0,
		fee:     1,
	}

}

func (b *Binance) GetBalance() int {
	return b.balance
}

func (b *Binance) Deposit(amount int) {
	fmt.Printf("New deposit at Binance: %d\n", amount)
	b.balance += amount - b.fee
}

func (b *Binance) WithDraw(amount int) error {
	fee := b.fee
	newBalance := b.balance - amount - fee
	if newBalance < 0 {
		return errors.New("insufficient funds")
	}

	b.balance = newBalance
	return nil
}
