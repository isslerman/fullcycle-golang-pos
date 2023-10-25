package tax

import (
	"errors"
	"time"
)

type Repository interface {
	SaveTax(amount float64) error
}

func CalculateTax(amount float64) (float64, error) {
	if amount <= 0 {
		return 0.0, errors.New("amount must be grater than 0")
	}
	if amount >= 1000 && amount < 20000 {
		return 10.0, nil
	}
	if amount >= 20000 {
		return 20.0, nil
	}
	return 5.0, nil
}

// second function with sleep, this is an example to run the benchmark
// and compare the performance between 1 and 2
func CalculateTax2(amount float64) float64 {
	time.Sleep(time.Millisecond)
	if amount < 0 {
		return 0
	}
	if amount >= 1000 && amount < 20000 {
		return 10.0
	}
	if amount >= 20000 {
		return 20.0
	}
	return 5.0
}

// Mocking

// save at database
func CalculateTaxAndSave(amount float64, repository Repository) error {
	tax := CalculateTax3(amount)
	return repository.SaveTax(tax)
}
func CalculateTax3(amount float64) float64 {
	if amount <= 0 {
		return 0.0
	}
	if amount >= 1000 && amount < 20000 {
		return 10.0
	}
	if amount >= 20000 {
		return 20.0
	}
	return 5.0
}
