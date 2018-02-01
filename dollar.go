package money

// Dollar is dollar struct
type Dollar struct {
	Money
	currency string
}

func NewDollar(amount int64) *Dollar {
	d := new(Dollar)
	d.amount = amount
	d.currency = "USD"

	return d
}

func (d Dollar) Times(multiplier int64) *Money {
	return NewMoney(d.amount * multiplier)
}

func (d Dollar) Currency() string {
	return d.currency
}
