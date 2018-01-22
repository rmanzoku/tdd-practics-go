package money

// Dollar is dollar struct
type Dollar struct {
	Amount int64
}

func (d *Dollar) Times(multiplier int64) {
	d.Amount = d.Amount * multiplier
}
