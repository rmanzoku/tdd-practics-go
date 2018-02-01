package money

import (
	"testing"
)

type dummyDollar struct {
	Money  int64
	amount int64
}

func TestMultiplication(t *testing.T) {
	var five = staticMoney.Dollar(5)

	if staticMoney.Dollar(10).Equals(five.Times(2)) {
		t.Errorf("ten != five.Times(2)")
	}

	if staticMoney.Dollar(15).Equals(five.Times(3)) {
		t.Errorf("fifteen != five.Times(3)")
	}
}

func TestEquality(t *testing.T) {
	var product = staticMoney.Dollar(5)

	if product.Equals(5) {
		t.Errorf("Dollar equals Int")
	}

	if product.Equals(dummyDollar{Money: 1, amount: 1}) {
		t.Errorf("Money is not struct")
	}

	if !product.Equals(*staticMoney.Dollar(5)) {
		t.Errorf("Dollar equals Dollar")
	}

	if product.Equals(*staticMoney.Dollar(6)) {
		t.Errorf("Dollar equals Dollar")
	}
	var productf = staticMoney.Franc(5)
	if !productf.Equals(*staticMoney.Franc(5)) {
		t.Errorf("Franc equals Franc")
	}
	if productf.Equals(*staticMoney.Franc(6)) {
		t.Errorf("Franc equals Franc")
	}

	// if staticMoney.Franc(5).Equals(*staticMoney.Dollar(5)) {
	// 	t.Errorf("Franc not equals Dollar")
	// }
}

func TestFrancMultiplication(t *testing.T) {
	var five = staticMoney.Franc(5)

	if staticMoney.Franc(10).Equals(five.Times(2)) {
		t.Errorf("ten != five.Times(2)")
	}

	if staticMoney.Franc(15).Equals(five.Times(3)) {
		t.Errorf("fifteen != five.Times(3)")
	}
}
