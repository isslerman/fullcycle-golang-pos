package main

import (
	"fmt"
)

type IBankAccount interface {
	GetBalance() int // 100 = 1 unit
	Deposit(amount int)
	WithDraw(amount int) error
}

func main() {

	// setting new account with both "banks"
	myAccounts := []IBankAccount{
		NewBinance(),
		NewNuBank(),
	}
	fmt.Printf("myAccounts: %v %T\n", myAccounts, myAccounts)

	for _, account := range myAccounts {
		account.Deposit(500) // 100 = 1 unit

		currentBalance := account.GetBalance()
		fmt.Printf("NuBank Balance: %d\n", currentBalance)

	}

	nu := NewNuBank()
	nu.Deposit(1500)
	nu.Deposit(10000)
	nu.WithDraw(3500)

	bi := NewBinance()
	bi.Deposit(1500)
	bi.Deposit(10000)
	bi.WithDraw(3500)

	// print balance using the method
	fmt.Println(nu.GetBalance())
	fmt.Println(bi.GetBalance())

	// print balance using the interface
	fmt.Println(nu)
	fmt.Println(bi)
}

func GetBalance(account IBankAccount) int {
	return account.GetBalance()
}
