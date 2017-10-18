package money

type Money struct {
	currency string
	amount   int
}

type Expression interface {
}

func NewDollar(amount int) Money {
	return Money{currency: "USD", amount: amount}
}

func NewFranc(amount int) Money {
	return Money{currency: "CHF", amount: amount}
}

func (d Money) Times(multiplier int) Money {
	return Money{currency: d.currency, amount: d.amount * multiplier}
}

func (d Money) Plus(addend Money) Money {
	return Money{currency: d.currency, amount: d.amount + addend.amount}
}
