package btclib

import "math/big"

func AddPoints(x1, y1, x2, y2 *big.Int) (x3, y3 *big.Int) {
	s := div(sub(y1, y2), sub(x1, x2))
	x3 = sub(pow(s, 2), add(x1, x2))
	y3 = sub(mul(sub(x1, x3), s), y1)
	return
}

func DoublePoint(x, y *big.Int) (x2, y2 *big.Int) {
	s := div(mul(getNum("3"), pow(x, 2)), mul(getNum("2"), y))
	x2 = sub(pow(s, 2), mul(getNum("2"), x))
	y2 = sub(mul(s, sub(x, x2)), y)
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
