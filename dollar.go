package money

// Dollar is dollar struct
type Dollar struct {
	Money
}

func NewDollar(amount int64, currency string) *Dollar {
	d := new(Dollar)
	d.amount = amount
	d.currency = currency

	return d
}

func (d Dollar) Times(multiplier int64) *Dollar {
	return staticMoney.Dollar(d.amount * multiplier)
}
