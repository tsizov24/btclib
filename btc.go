package btclib

import "math/big"

func AddPoints(x1, y1, x2, y2 *big.Int) (x3, y3 *big.Int) {
	s1 := big.NewInt(0)
	s1.Sub(y1, y2).Mod(s1, p)
	s2 := big.NewInt(0)
	s2.Sub(x1, x2).Mod(s2, p)
	s := big.NewInt(0)
	s.ModInverse(s2, p).Mul(s, s1).Mod(s, p)
	x3 = big.NewInt(0)
	x3.Mul(s, s).Mod(x3, p).Sub(x3, x1).Sub(x3, x2).Mod(x3, p)
	y3 = big.NewInt(0)
	y3.Sub(x1, x3).Mul(y3, s).Mod(y3, p).Sub(y3, y1).Mod(y3, p)
	return
}

func DoublePoint(x, y *big.Int) (x2, y2 *big.Int) {
	s1 := big.NewInt(3)
	s1.Mul(s1, x).Mod(s1, p).Mul(s1, x).Mod(s1, p)
	s2 := big.NewInt(2)
	s2.Mul(s2, y).Mod(s2, p)
	s := big.NewInt(0)
	s.ModInverse(s2, p).Mul(s, s1).Mod(s, p)
	x21 := big.NewInt(0)
	x21.Mul(s, s).Mod(x21, p)
	x22 := big.NewInt(2)
	x22.Mul(x22, x).Mod(x22, p)
	x2 = big.NewInt(0)
	x2.Sub(x21, x22).Mod(x2, p)
	y2 = big.NewInt(0)
	y2.Sub(x, x2).Mul(y2, s).Mod(y2, p).Sub(y2, y).Mod(y2, p)
	return
}

// If k % n == 0 x and y will be nil
func PrivToPub(k *big.Int) (x, y *big.Int) {
	k.Mod(k, n)
	for i := 0; i < k.BitLen(); i++ {
		if k.Bit(i) == 1 {
			if x == nil {
				x = big.NewInt(0)
				x.Set(ax[i])
				y = big.NewInt(0)
				y.Set(ay[i])
			} else {
				x, y = AddPoints(x, y, ax[i], ay[i])
			}
		}
	}
	return
}
