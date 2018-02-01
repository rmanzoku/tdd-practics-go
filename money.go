package money

import (
	"reflect"
)

type Money struct {
	amount int64
}

func isMoneyObj(input interface{}) bool {
	rv := reflect.ValueOf(input)
	if rv.Kind() != reflect.Struct {
		return false
	}
	if rv.FieldByName("Money").IsValid() {
		if rv.FieldByName("Money").Type() == reflect.TypeOf(Money{}) {
			return true
		}
	}
	return false
}

func NewMoney(amount int64) *Money {
	m := new(Money)
	m.amount = amount

	return m
}

func (m *Money) Equals(input interface{}) bool {
	if !isMoneyObj(input) {
		return false
	}

	//fmt.Println(reflect.TypeOf(m))

	rv := reflect.ValueOf(input)
	amount := rv.FieldByName("amount").Int()
	return (m.amount == amount)
}

var staticMoney = NewMoney(0)

func (m *Money) Dollar(amount int64) *Dollar {
	return NewDollar(amount)
}

func (m *Money) Franc(amount int64) *Franc {
	return NewFranc(amount)
}
