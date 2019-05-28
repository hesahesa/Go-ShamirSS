package shamirssgo

import (
	"crypto/rand"
	"math/big"
)

const PrimeCertainty = 128

type ShamirSecret struct {
	secret            *big.Int
	threshold         int
	modulus           *big.Int
	reconstructVector []*big.Int
}

func New(secret *big.Int, threshold int, modulus *big.Int) *ShamirSecret {
	if !modulus.ProbablyPrime(PrimeCertainty) {
		panic("modulus fails prime test by ProbablePrime(), try to use prime (or prime enough)")
	}

	modulusBitLen := modulus.BitLen()
	max := new(big.Int)
	max.Exp(
		big.NewInt(2),
		big.NewInt(int64(modulusBitLen)),
		nil,
	).Sub(max, big.NewInt(1))

	reconstructVector := make([]*big.Int, threshold)
	reconstructVector[0] = big.NewInt(1)

	for i := 1; i < threshold-1; i++ {
		randNum, _ := rand.Int(rand.Reader, max)
		reconstructVector[i] = randNum.Mod(randNum, modulus)
	}

	return &ShamirSecret{
		secret:            secret.Mod(secret, modulus),
		threshold:         threshold,
		modulus:           modulus,
		reconstructVector: reconstructVector,
	}
}
