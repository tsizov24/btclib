package btclib

import "math/big"

func getNum(s string) *big.Int {
	res := big.NewInt(0)
	res.SetString(s, 16)
	return res
}

func add(a, b *big.Int) *big.Int {
	res := big.NewInt(0)
	return mod(res.Add(a, b))
}

func sub(a, b *big.Int) *big.Int {
	res := big.NewInt(0)
	return mod(res.Sub(a, b))
}

func mul(a, b *big.Int) *big.Int {
	res := big.NewInt(0)
	return mod(res.Mul(a, b))
}

func div(a, b *big.Int) *big.Int {
	res := big.NewInt(0)
	return mod(res.ModInverse(b, p).Mul(res, a))
}

func pow(a *big.Int, n int) *big.Int {
	res := big.NewInt(1)
	a = mod(a)
	switch {
	case n < 0:
		return nil
	case n < 4:
		for i := 0; i < n; i++ {
			res = mod(res.Mul(res, a))
		}
	default:
		k := mod(a)
		for n > 0 {
			if n%2 == 1 {
				res = mod(res.Mul(res, k))
			}
			n /= 2
			if n > 0 {
				k = mod(k.Mul(k, k))
			}
		}
	}
	return res
}

func mod(a *big.Int) *big.Int {
	res := big.NewInt(0)
	return res.Mod(a, p)
}
