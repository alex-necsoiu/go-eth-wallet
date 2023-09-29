package wallet_test

import (
	"os"
	"testing"

	e2types "github.com/alex-necsoiu/go-eth-wallet/types"
)

func TestMain(m *testing.M) {
	if err := e2types.InitBLS(); err != nil {
		panic(err)
	}
	os.Exit(m.Run())
}
