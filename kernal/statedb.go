package kernal

import (
	"math/big"
)

// StateDB is an EVM database for full state querying.
type StateDB interface {
	CreateAccount(Address)

	SubBalance(Address, *big.Int)
	AddBalance(Address, *big.Int)
	GetBalance(Address) *big.Int

	GetNonce(Address) uint64
	SetNonce(Address, uint64)

	GetCodeHash(Address) Hash
	GetCode(Address) []byte
	SetCode(Address, []byte)
	GetCodeSize(Address) int

	AddRefund(uint64)
	GetRefund() uint64

	GetState(Address, Hash) Hash
	SetState(Address, Hash, Hash)

	Suicide(Address) bool
	HasSuicided(Address) bool

	// Exist reports whether the given account exists in state.
	// Notably this should also return true for suicided accounts.
	Exist(Address) bool
	// Empty returns whether the given account is empty. Empty
	// is defined according to EIP161 (balance = nonce = code = 0).
	Empty(Address) bool

	RevertToSnapshot(int)
	Snapshot() int
	//Define function aimed at replacement of CanTransfer and Transfer
	HaveSufficientBalance(Address, *big.Int) bool
	TransferBalance(Address, Address, *big.Int)

	AddLog(*Log)
	AddPreimage(Hash, []byte)
	ForEachStorage(Address, func(Hash, Hash) bool)
}

type AddressHandler interface {
	CreateAddress(b Address, nonce uint64) Address
}
