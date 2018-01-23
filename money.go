package money

// Dollar is dollar struct
type Dollar struct {
	Amount int64
}

func NewDollar(amount int64) *Dollar {
	d := new(Dollar)
	d.Amount = amount

	return d
}

func (d Dollar) Times(multiplier int64) *Dollar {
	return NewDollar(d.Amount * multiplier)
}

func (d Dollar) Equals(input Dollar) bool {
	return (d.Amount == input.Amount)
}
