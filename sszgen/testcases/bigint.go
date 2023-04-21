package testcases

import (
	"math/big"
)

//go:generate go run ../main.go --path bigint.go

type BigIntType struct {
	BigInt big.Int `ssz-max:"129"`
}
