package money

import (
	"reflect"
)

type Money struct {
	amount int64
}

func (m Money) Equals(input interface{}) bool {
	i := reflect.ValueOf(input)
	iAmount := i.FieldByName("amount").Int()
	return (m.amount == iAmount)
}
