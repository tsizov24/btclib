package btclib

import "math/big"

func getNum(s string) *big.Int {
	res := big.NewInt(0)
	res.SetString(s, 16)
	return res
}
