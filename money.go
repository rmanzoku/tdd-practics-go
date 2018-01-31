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
	return rv.FieldByName("Money").IsValid()
}

func (m Money) Equals(input interface{}) bool {
	if !isMoneyObj(input) {
		return false
	}

	rv := reflect.ValueOf(input)
	amount := rv.FieldByName("amount").Int()
	return (m.amount == amount)
}
