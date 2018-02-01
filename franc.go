package money

// Franc is franc struct
type Franc struct {
	Money
	currency string
}

func NewFranc(amount int64) *Franc {
	f := new(Franc)
	f.amount = amount
	f.currency = "CHF"

	return f
}

func (f Franc) Times(multiplier int64) *Money {
	return NewMoney(f.amount * multiplier)
}

func (f Franc) Currency() string {
	return f.currency
}
