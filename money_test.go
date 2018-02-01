package money

import (
	"testing"
)

type dummyDollar struct {
	Money  int64
	amount int64
}

func TestMultiplication(t *testing.T) {
	var five = NewDollar(5)

	if *NewDollar(10) != *five.Times(2) {
		t.Errorf("ten != five.Times(2)")
	}

	if *NewDollar(15) != *five.Times(3) {
		t.Errorf("fifteen != five.Times(3)")
	}
}

func TestEquality(t *testing.T) {
	var product = NewDollar(5)

	if product.Equals(5) {
		t.Errorf("Dollar equals Int")
	}

	if product.Equals(dummyDollar{Money: 1, amount: 1}) {
		t.Errorf("Money is not struct")
	}

	if !product.Equals(*NewDollar(5)) {
		t.Errorf("Dollar equals Dollar")
	}

	if product.Equals(*NewDollar(6)) {
		t.Errorf("Dollar equals Dollar")
	}
	var productf = NewFranc(5)
	if !productf.Equals(*NewFranc(5)) {
		t.Errorf("Franc equals Franc")
	}
	if productf.Equals(*NewFranc(6)) {
		t.Errorf("Franc equals Franc")
	}
}

func TestFrancMultiplication(t *testing.T) {
	var five = NewFranc(5)

	if *NewFranc(10) != *five.Times(2) {
		t.Errorf("ten != five.Times(2)")
	}

	if *NewFranc(15) != *five.Times(3) {
		t.Errorf("fifteen != five.Times(3)")
	}
}
