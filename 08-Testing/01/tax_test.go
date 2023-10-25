package tax

import (
	"testing"
)

func TestCalculateTax(t *testing.T) {
	amount := 500.0
	// passed
	expected := 5.0
	// fail
	// expected := 6.0

	result := CalculateTax(amount)

	if result != expected {
		t.Errorf("Expected %f but got %f", expected, result)
	}
}

func TestCalculateTaxBatch(t *testing.T) {
	type calcTax struct {
		amount, expect float64
	}

	table := []calcTax{
		{500.0, 5.0},
		{1000.0, 10.0},
		{1500.0, 10.0},
	}

	for _, item := range table {
		result := CalculateTax(item.amount)
		if result != item.expect {
			t.Errorf("Expected %f but got %f", item.expect, result)
		}
	}
}

func TestCalculateTax3(t *testing.T) { // T Testing
	type calcTax struct {
		amount, expect float64
	}

	// nessa table não cobrimos todos os casos, para justamente ver algum ponto não coberto.
	table := []calcTax{
		{500.0, 5.0},
		{1000.0, 10.0},
		{1500.0, 10.0},
		//{0.0, 0.0}, // esse valor não foi testado para que seja acusado no teste de coverage.
	}

	for _, item := range table {
		result := CalculateTax(item.amount)
		if result != item.expect {
			t.Errorf("Expected %f but got %f", item.expect, result)
		}
	}
}

// BENCHMARK FAST
func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(500.0)
	}
}

// BENCHMARK SLOW
func BenchmarkCalculateTax2(b *testing.B) { // B = Benchmark
	for i := 0; i < b.N; i++ {
		CalculateTax2(500.0)
	}
}

// TO RUN FUZZ TEST -> go test -fuzz=. -run=^#
// FUZZ TEST
// FUZZ TEST GENERATE A testdata DIR WITH THE VALUE THAT GENERATED THE ERROR
func FuzzCalculateTax(f *testing.F) { // F = FUZZ
	//  some examples to give the range to fuzz work
	seed := []float64{-1, -2, -2.5, 500.0, 1000.0, 1501.0}
	for _, amount := range seed {
		f.Add(amount)
	}
	f.Fuzz(func(t *testing.T, amount float64) {
		result := CalculateTax(amount)
		if amount <= 0 && result != 0 {
			t.Errorf("Received %f but expected 0", result)
		}
		// here we have a condition problem at the case.
		if amount >= 20000 && result != 20 {
			t.Errorf("Received %f but expected 0", result)
		}
	})
}
