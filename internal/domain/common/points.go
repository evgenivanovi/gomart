package common

import "math/big"

/* __________________________________________________ */

type Points struct {
	amount *big.Float
}

func NewPoints(amount float64) *Points {
	return NewBigPoints(big.NewFloat(amount))
}

func NewBigPoints(amount *big.Float) *Points {
	return &Points{
		amount: amount,
	}
}

func NewEmptyPoints() *Points {
	return NewPoints(0)
}

func (p Points) Amount() float64 {
	am, _ := p.amount.Float64()
	return am
}

func (p Points) MoreThan(amount float64) bool {
	return p.amount.Cmp(big.NewFloat(amount)) > 0
}

func (p Points) MoreThanEQ(amount float64) bool {
	return p.amount.Cmp(big.NewFloat(amount)) >= 0
}

func (p Points) LessThan(amount float64) bool {
	return p.amount.Cmp(big.NewFloat(amount)) < 0
}

func (p Points) LessThanEQ(amount float64) bool {
	return p.amount.Cmp(big.NewFloat(amount)) <= 0
}

func (p Points) Add(amount float64) Points {
	am := new(big.Float)
	am.Add(p.amount, big.NewFloat(amount))
	return *NewBigPoints(am)
}

func (p Points) Minus(amount float64) Points {
	am := new(big.Float)
	am.Sub(p.amount, big.NewFloat(amount))
	return *NewBigPoints(am)
}

/* __________________________________________________ */
