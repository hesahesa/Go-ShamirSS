package shamirss

import "math/big"

type ShamirSecret struct {
	secret            big.Int
	threshold         int
	modulus           big.Int
	reconstructVector []big.Int
}

func New(secret big.Int, threshold int, modulus big.Int) *ShamirSecret {
	return &ShamirSecret{
		secret:    secret,
		threshold: threshold,
		modulus:   modulus,
	}
}
