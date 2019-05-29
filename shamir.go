package shamirssgo

import (
	"crypto/rand"
	"math/big"

	"github.com/pkg/errors"
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

func (ss *ShamirSecret) Shares(index int) (share *big.Int, err error) {
	if index <= 0 {
		err = errors.New("shares index need to be greater or equal than 0")
		return
	} else {
		// compute polynomial P(x) from secret and reconstruction vector
		share = big.NewInt(0)
		share.Add(share, ss.secret)
		for i := 1; i <= ss.threshold-1; i++ {
			xPowi := big.NewInt(int64(index))
			xPowi.Exp(xPowi, big.NewInt(int64(i)), ss.modulus)

			recVerMul := new(big.Int)
			recVerMul.Mul(ss.reconstructVector[i], xPowi)
			recVerMul.Mod(recVerMul, ss.modulus)

			share.Add(share, recVerMul)
			share.Mod(share, ss.modulus)
		}

		return
	}
}

func ReconstructSecret(sharesMap map[int]*big.Int, modulus *big.Int) (secret *big.Int, err error) {
	secret = big.NewInt(0)
	keySet := make(map[int]bool, len(sharesMap))
	for k := range sharesMap {
		keySet[k] = true
	}
	for k, v := range sharesMap {
		langrangeCoeff, err := langrangeCoeff(k, keySet, modulus)
		if err != nil {
			err = errors.Wrap(err, "fail getting lagrange coefficient in index "+string(k))
		}
		share := v

		lagrangeCoeffMult := new(big.Int)
		lagrangeCoeffMult.Mul(langrangeCoeff, share).Mod(lagrangeCoeffMult, modulus)
		secret.Add(secret, lagrangeCoeffMult).Mod(secret, modulus)
	}
	return
}

func langrangeCoeff(i int, setIndices map[int]bool, modulus *big.Int) (coeff *big.Int, err error) {
	coeff = big.NewInt(1)
	for t := range setIndices {
		if i == t {
			continue
		}
		numerator := big.NewInt(0)
		numerator.Sub(numerator, big.NewInt(int64(t))).Mod(numerator, modulus)

		denumerator := big.NewInt(int64(i - t))
		denumerator.Mod(denumerator, modulus)

		currVal := new(big.Int)
		denumeratorInverse := new(big.Int)
		denumeratorInverse.ModInverse(denumerator, modulus)
		currVal.Mul(numerator, denumeratorInverse)

		coeff.Mul(coeff, currVal).Mod(coeff, modulus)
	}
	return
}
