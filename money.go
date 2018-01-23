package money

// Dollar is dollar struct
type Dollar struct {
	Amount int64
}

func (d *Dollar) Times(multiplier int64) Dollar {
	return Dollar{Amount: d.Amount * multiplier}
}
