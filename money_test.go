package money

import (
	"testing"
)

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
	if !product.Equals(NewDollar(5)) {
		t.Errorf("Dollar equals Dollar")
	}
	if product.Equals(NewDollar(6)) {
		t.Errorf("Dollar equals Dollar")
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
