package money

// Franc is dollar struct
type Franc struct {
	amount int64
}

func NewFranc(amount int64) *Franc {
	d := new(Franc)
	d.amount = amount

	return d
}

func (d Franc) Times(multiplier int64) *Franc {
	return NewFranc(d.amount * multiplier)
}

func (d Franc) Equals(input *Franc) bool {
	return (d.amount == input.amount)
}
