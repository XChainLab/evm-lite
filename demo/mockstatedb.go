package main

import (
	"evm/kernel"
	"math/big"
)

type MockStateDB struct {
	stateStore map[kernel.Address][]byte
}

func MakeNewMockStateDB() *MockStateDB {
	mockstatedb := new(MockStateDB)
	mockstatedb.stateStore = make(map[kernel.Address][]byte)
	return mockstatedb
}

func (MockStateDB) CreateAccount(kernel.Address)           {}
func (MockStateDB) SubBalance(kernel.Address, *big.Int)    {}
func (MockStateDB) AddBalance(kernel.Address, *big.Int)    {}
func (MockStateDB) GetBalance(kernel.Address) *big.Int     { return nil }
func (MockStateDB) GetNonce(kernel.Address) uint64         { return 0 }
func (MockStateDB) SetNonce(kernel.Address, uint64)        {}
func (MockStateDB) GetCodeHash(kernel.Address) kernel.Hash { return kernel.Hash{} }
func (mockstatedb MockStateDB) GetCode(address kernel.Address) []byte {
	_, ok := mockstatedb.stateStore[address]
	if ok {
		return mockstatedb.stateStore[address]
	} else {
		return nil
	}
}
func (mockstatedb MockStateDB) SetCode(address kernel.Address, data []byte) {
	mockstatedb.stateStore[address] = data
}
func (mockstatedb MockStateDB) GetCodeSize(address kernel.Address) int {
	_, ok := mockstatedb.stateStore[address]
	if ok {
		return len(mockstatedb.stateStore[address])
	} else {
		return 0
	}
}
func (MockStateDB) AddRefund(uint64)                                  {}
func (MockStateDB) GetRefund() uint64                                 { return 0 }
func (MockStateDB) GetState(kernel.Address, kernel.Hash) kernel.Hash  { return kernel.Hash{} }
func (MockStateDB) SetState(kernel.Address, kernel.Hash, kernel.Hash) {}
func (MockStateDB) Suicide(kernel.Address) bool                       { return false }
func (MockStateDB) HasSuicided(kernel.Address) bool                   { return false }
func (MockStateDB) Exist(kernel.Address) bool {
	return true
}
func (MockStateDB) Empty(kernel.Address) bool                                          { return false }
func (MockStateDB) RevertToSnapshot(int)                                               {}
func (MockStateDB) Snapshot() int                                                      { return 0 }
func (MockStateDB) AddLog(*kernel.Log)                                                 {}
func (MockStateDB) AddPreimage(kernel.Hash, []byte)                                    {}
func (MockStateDB) ForEachStorage(kernel.Address, func(kernel.Hash, kernel.Hash) bool) {}
func (MockStateDB) HaveSufficientBalance(kernel.Address, *big.Int) bool {
	return true
}
func (MockStateDB) TransferBalance(kernel.Address, kernel.Address, *big.Int) {

}
