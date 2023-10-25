package tax

import "time"

func CalculateTax(amount float64) float64 {
	if amount == 0 {
		return 0
	}
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
