package btclib

import "math/big"

var (
	n      *big.Int
	p      *big.Int
	ax, ay [256]*big.Int
)

func init() {
	n = getNum("fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141")
	p = getNum("fffffffffffffffffffffffffffffffffffffffffffffffffffffffefffffc2f")

	ax[0] = getNum("79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798")
	ay[0] = getNum("483ada7726a3c4655da4fbfc0e1108a8fd17b448a68554199c47d08ffb10d4b8")

	for i := 1; i < 256; i++ {
		ax[i], ay[i] = DoublePoint(ax[i-1], ay[i-1])
	}
}

func GetBasePoint() (gx, gy *big.Int) {
	return getIntCopy(ax[0]), getIntCopy(ay[0])
}

func GetN() *big.Int {
	return getIntCopy(n)
}

func GetP() *big.Int {
	return getIntCopy(p)
}

func getIntCopy(a *big.Int) *big.Int {
	res := big.NewInt(0)
	res.Set(a)
	return res
}
