package money

// Franc is franc struct
type Franc struct {
	Money
}

func NewFranc(amount int64) *Franc {
	d := new(Franc)
	d.amount = amount

	return d
}

func (d Franc) Times(multiplier int64) *Franc {
	return NewFranc(d.amount * multiplier)
}
