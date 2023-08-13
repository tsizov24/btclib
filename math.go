package btclib

import "math/big"

func getNum(s string) *big.Int {
	res := big.NewInt(0)
	res.SetString(s, 16)
	return res
}

func add(a, b *big.Int) *big.Int {
	res := big.NewInt(0)
	res.Add(a, b)
	return mod(res)
}

func sub(a, b *big.Int) *big.Int {
	res := big.NewInt(0)
	res.Sub(a, b)
	return mod(res)
}

func mul(a, b *big.Int) *big.Int {
	res := big.NewInt(0)
	res.Mul(a, b)
	return mod(res)
}

func div(a, b *big.Int) *big.Int {
	res := big.NewInt(0)
	res.ModInverse(b, p)
	res.Mul(res, a).Mod(res, p)
	return res
}

func pow(a *big.Int, n int) *big.Int {
	res := big.NewInt(1)
	for i := 0; i < n; i++ {
		res = mod(res.Mul(res, a))
	}
	return res
}

func mod(a *big.Int) *big.Int {
	res := big.NewInt(0)
	res.Mod(a, p)
	return res
}
