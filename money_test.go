package money

import (
	"testing"
)

func TestMultiplication(t *testing.T) {
	var five = Dollar{Amount: 5}

	ten := Dollar{Amount: 10}
	if ten != five.Times(2) {
		t.Errorf("ten != five.Times(2)")
	}

	fifteen := Dollar{Amount: 15}
	if fifteen != five.Times(3) {
		t.Errorf("fifteen != five.Times(3)")
	}
}

func TestEquality(t *testing.T) {
	var product = Dollar{Amount: 5}
	if !product.Equals(Dollar{Amount: 5}) {
		t.Errorf("Dollar equals Dollar")
	}
	if product.Equals(Dollar{Amount: 6}) {
		t.Errorf("Dollar equals Dollar")
	}
}
