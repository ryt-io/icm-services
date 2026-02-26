package network

import (
	"crypto/ecdsa"

	"github.com/ryt-io/libevm/common"
)

type LocalNetwork interface {
	GetFundedAccountInfo() (common.Address, *ecdsa.PrivateKey)
	TearDownNetwork()
}
