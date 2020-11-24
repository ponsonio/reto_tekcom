package services

import (
	"math/big"
)

type Multiplicator interface {
	Multiply(a int64, b int64) (res *big.Int, err error)
}

type MultiplicatorSimple struct {
}

func (m *MultiplicatorSimple) Multiply(a int64, b int64) (res *big.Int, err error) {
	res = big.NewInt(a)
	res.Mul(res, big.NewInt(b))

	//TODO: the idea is to detect possible overflows and return errors on those cases
	return res, nil
}

func NewMultiplicator() (Multiplicator, error) {
	return &MultiplicatorSimple{}, nil
}
