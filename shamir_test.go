package shamirssgo_test

import (
	"math/big"
	"testing"

	"github.com/hesahesa/shamirssgo"
)

func TestShamir(t *testing.T) {
	secret := big.NewInt(100)
	shamirssgo.New(secret)
}
