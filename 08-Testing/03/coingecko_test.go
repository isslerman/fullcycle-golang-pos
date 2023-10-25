package main

import (
	"testing"
)

func TestCoingeckoPing(t *testing.T) {
	expected := "(V3) To the Moon!"

	response, _ := CoingeckoPing()
	result := response.Geckosays

	if result != expected {
		t.Errorf("FAIL: Expected %s but got %s", expected, result)
	} else {
		t.Logf("PASS: Expected %s and got %s", expected, result)
	}
}
