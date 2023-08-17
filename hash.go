package btclib

import (
	"crypto/sha256"
	"math/big"

	"golang.org/x/crypto/ripemd160"
)

// pref should be 2 or 3 for compressed public keys and 4 for uncompressed
func PubToHash160(x, y *big.Int, compress bool) []byte {
	b := make([]byte, 33)
	if compress {
		b[0] = byte(y.Bit(0)) + 2
	} else {
		b[0] = 4
	}
	c := 33 - len(x.Bytes())
	for i, v := range x.Bytes() {
		b[i+c] = v
	}
	return RIPEMD160(SHA256(b))
}

func RIPEMD160(b []byte) []byte {
	h := ripemd160.New()
	h.Write(b)
	return h.Sum(nil)
}

func SHA256(b []byte) []byte {
	h := sha256.New()
	h.Write(b)
	return h.Sum(nil)
}
