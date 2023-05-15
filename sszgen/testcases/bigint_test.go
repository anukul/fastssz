package testcases

import (
	"bytes"
	"fmt"
	ssz "github.com/anukul/fastssz"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestBigIntEncodeDecode(t *testing.T) {
	tests := []struct {
		name   string
		bigint big.Int
	}{
		{
			name:   "128",
			bigint: *big.NewInt(128),
		},
		{
			name:   "2^1024",
			bigint: *big.NewInt(0).Exp(big.NewInt(2), big.NewInt(1024), nil),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			fmt.Println(test.bigint.Bytes())
			buf := make([]byte, test.bigint.BitLen()/4)
			buf = ssz.MarshalBigInt(buf, test.bigint)
			n := ssz.UnmarshalBigInt(buf)
			assert.Equal(t, test.bigint, n)
		})
	}
}

func TestBigIntRoot(t *testing.T) {
	tests := []struct {
		name   string
		bigint big.Int
		root   []byte
	}{
		{
			name:   "128",
			bigint: *big.NewInt(128),
			root:   []byte{36, 165, 93, 23, 246, 17, 99, 63, 21, 144, 241, 36, 50, 251, 108, 136, 142, 154, 119, 255, 250, 179, 211, 93, 7, 185, 3, 150, 243, 67, 171, 148},
		},
		{
			name:   "2^1024",
			bigint: *big.NewInt(0).Exp(big.NewInt(2), big.NewInt(1024), nil),
			root:   []byte{187, 64, 122, 238, 209, 243, 210, 39, 4, 92, 131, 114, 103, 155, 189, 105, 63, 209, 0, 41, 148, 183, 170, 40, 251, 194, 181, 205, 94, 128, 162, 97},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			bigintStruct := BigIntType{BigInt: test.bigint}
			bigintRoot, err := bigintStruct.HashTreeRoot()
			if err != nil {
				t.Fatalf("unexpected error: %v\n", err)
			}
			// Expect the root to match the expectation.
			if !bytes.Equal(test.root, bigintRoot[:]) {
				fmt.Println(bigintRoot)
				t.Fatalf("root mismatch with bigint type")
			}
		})
	}
}
