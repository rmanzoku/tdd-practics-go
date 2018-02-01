package money

// Franc is franc struct
type Franc struct {
	Money
}

func NewFranc(amount int64, currency string) *Franc {
	f := new(Franc)
	f.amount = amount
	f.currency = "CHF"

	return f
}

func (f Franc) Times(multiplier int64) *Franc {
	return staticMoney.Franc(f.amount * multiplier)
}
