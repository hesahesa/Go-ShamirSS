package shamirssgo_test

import (
	"math/big"
	"testing"

	"github.com/hesahesa/shamirssgo"
	"github.com/stretchr/testify/assert"
)

func TestShamir(t *testing.T) {
	secret := big.NewInt(1024)
	threshold := 3
	modulus := big.NewInt(1000003)

	shamirSecret := shamirssgo.New(secret, threshold, modulus)

	shares := make([]*big.Int, 6)
	for idx := 1; idx < 6; idx++ {
		val, err := shamirSecret.Shares(idx)
		assert.Nil(t, err)
		shares[idx] = val
	}

	selectedShares := make(map[int]*big.Int, 3)
	selectedShares[1] = shares[1]
	selectedShares[4] = shares[4]
	selectedShares[5] = shares[5]

	computedSecret, err := shamirssgo.ReconstructSecret(selectedShares, modulus)
	assert.Nil(t, err)
	assert.Equal(t, computedSecret, secret)
}
