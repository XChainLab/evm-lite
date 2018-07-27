package main

import (
	"evm/kernal"
	"math/big"
)

type MockStateDB struct {
	stateStore map[kernal.Address][]byte
}

func MakeNewMockStateDB() *MockStateDB {
	mockstatedb := new(MockStateDB)
	mockstatedb.stateStore = make(map[kernal.Address][]byte)
	return mockstatedb
}

func (MockStateDB) CreateAccount(kernal.Address)           {}
func (MockStateDB) SubBalance(kernal.Address, *big.Int)    {}
func (MockStateDB) AddBalance(kernal.Address, *big.Int)    {}
func (MockStateDB) GetBalance(kernal.Address) *big.Int     { return nil }
func (MockStateDB) GetNonce(kernal.Address) uint64         { return 0 }
func (MockStateDB) SetNonce(kernal.Address, uint64)        {}
func (MockStateDB) GetCodeHash(kernal.Address) kernal.Hash { return kernal.Hash{} }
func (mockstatedb MockStateDB) GetCode(address kernal.Address) []byte {
	_, ok := mockstatedb.stateStore[address]
	if ok {
		return mockstatedb.stateStore[address]
	} else {
		return nil
	}
}
func (mockstatedb MockStateDB) SetCode(address kernal.Address, data []byte) {
	mockstatedb.stateStore[address] = data
}
func (mockstatedb MockStateDB) GetCodeSize(address kernal.Address) int {
	_, ok := mockstatedb.stateStore[address]
	if ok {
		return len(mockstatedb.stateStore[address])
	} else {
		return 0
	}
}
func (MockStateDB) AddRefund(uint64)                                  {}
func (MockStateDB) GetRefund() uint64                                 { return 0 }
func (MockStateDB) GetState(kernal.Address, kernal.Hash) kernal.Hash  { return kernal.Hash{} }
func (MockStateDB) SetState(kernal.Address, kernal.Hash, kernal.Hash) {}
func (MockStateDB) Suicide(kernal.Address) bool                       { return false }
func (MockStateDB) HasSuicided(kernal.Address) bool                   { return false }
func (MockStateDB) Exist(kernal.Address) bool {
	return true
}
func (MockStateDB) Empty(kernal.Address) bool                                          { return false }
func (MockStateDB) RevertToSnapshot(int)                                               {}
func (MockStateDB) Snapshot() int                                                      { return 0 }
func (MockStateDB) AddLog(*kernal.Log)                                                 {}
func (MockStateDB) AddPreimage(kernal.Hash, []byte)                                    {}
func (MockStateDB) ForEachStorage(kernal.Address, func(kernal.Hash, kernal.Hash) bool) {}
func (MockStateDB) HaveSufficientBalance(kernal.Address, *big.Int) bool {
	return true
}
func (MockStateDB) TransferBalance(kernal.Address, kernal.Address, *big.Int) {

}
