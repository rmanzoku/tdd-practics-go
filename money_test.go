package money

import (
	"testing"
)

func TestMultiplication(t *testing.T) {
	var five = Dollar{Amount: 5}

	var product = five.Times(2)
	if 10 != product.Amount {
		t.Errorf("five.amount should be 10")
	}

	product = five.Times(3)
	if 15 != product.Amount {
		t.Errorf("five.amount should be 15")
	}
}

func TestEquality(t *testing.T) {
	var product = Dollar{Amount: 5}
	if !product.Equals(Dollar{Amount: 5}) {
		t.Errorf("Dollar equals Dollar")
	}
}
