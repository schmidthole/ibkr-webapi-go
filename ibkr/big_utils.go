package ibkr

import "math/big"

func BigToSignedBytes(i *big.Int) []byte {
	absBytes := i.Bytes()

	if (i.BitLen() % 8) == 0 {
		return append([]byte{0}, absBytes...)
	} else {
		return absBytes
	}
}
