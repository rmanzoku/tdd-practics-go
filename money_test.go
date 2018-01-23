package money

import (
	"testing"
)

func TestMultiplication(t *testing.T) {
	var five = Dollar{Amount: 5}

	five.Times(2)
	if 10 != five.Amount {
		t.Errorf("five.amount should be 10")
	}

	five.Times(3)
	if 15 != five.Amount {
		t.Errorf("five.amount should be 15")
	}
}
