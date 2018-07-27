package kernal

import (
	"math/big"
)

// Context provides the EVM with auxiliary information. Once provided
// it shouldn't be modified.
type Context struct {
	// Message information
	Origin   Address  // Provides information for ORIGIN
	GasPrice *big.Int // Provides information for GASPRICE

	// Block information
	Coinbase    Address  // Provides information for COINBASE
	GasLimit    uint64   // Provides information for GASLIMIT
	BlockNumber *big.Int // Provides information for NUMBER
	Time        *big.Int // Provides information for TIME
	Difficulty  *big.Int // Provides information for DIFFICULTY
}
