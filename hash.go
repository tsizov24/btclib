package btclib

import (
	"crypto/sha256"
	"math/big"

	"golang.org/x/crypto/ripemd160"
)

// pref should be 2 or 3 for compressed public keys and 4 for uncompressed
func PubToHash160(x, y *big.Int, compress bool) []byte {
	var pref byte
	if compress {
		pref = byte(y.Bit(0)) + 2
	} else {
		pref = 4
	}
	return RIPEMD160(SHA256(append([]byte{pref}, x.Bytes()...)))
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
