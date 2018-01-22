package money

import "testing"

func testMultiplication(t *testing.T) {
	var five = Dollar(5)
	Dollar.times(2)

	if 10 == five.amount {
		t.Errorf("five.amount should be 10")
	}
}
