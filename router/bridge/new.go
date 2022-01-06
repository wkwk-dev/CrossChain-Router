package bridge

import (
	"math/big"

	"github.com/anyswap/CrossChain-Router/v3/tokens"
	"github.com/anyswap/CrossChain-Router/v3/tokens/eth"
	"github.com/anyswap/CrossChain-Router/v3/tokens/solana"
)

// NewCrossChainBridge new bridge
func NewCrossChainBridge(chainID *big.Int) tokens.IBridge {
	switch {
	case solana.SupportChainID(chainID):
		return solana.NewCrossChainBridge()
	default:
		return eth.NewCrossChainBridge()
	}
}