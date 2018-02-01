package money

// Dollar is dollar struct
type Dollar struct {
	Money
}

func NewDollar(amount int64) *Dollar {
	d := new(Dollar)
	d.amount = amount

	return d
}

func (d Dollar) Times(multiplier int64) *Money {
	return NewMoney(d.amount * multiplier)
}
