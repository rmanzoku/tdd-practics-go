package money

// Dollar is dollar struct
type Dollar struct {
	amount int64
}

func NewDollar(amount int64) *Dollar {
	d := new(Dollar)
	d.amount = amount

	return d
}

func (d Dollar) Times(multiplier int64) *Dollar {
	return NewDollar(d.amount * multiplier)
}

func (d Dollar) Equals(input *Dollar) bool {
	return (d.amount == input.amount)
}
