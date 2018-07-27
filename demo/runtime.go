package main

import (
	"evm/kernal"
	"math/big"
	"time"
)

func CreateLogTracer() *kernal.StructLogger {
	logConf := kernal.LogConfig{
		DisableMemory:  false,
		DisableStack:   false,
		DisableStorage: false,
		Debug:          false,
		Limit:          0,
	}
	return kernal.NewStructLogger(&logConf)

}
func CreateChainConfig() *kernal.ChainConfig {
	chainCfg := kernal.ChainConfig{
		ChainID:        big.NewInt(1),
		HomesteadBlock: new(big.Int),
		DAOForkBlock:   new(big.Int),
		DAOForkSupport: false,
		EIP150Block:    new(big.Int),
		EIP155Block:    new(big.Int),
		EIP158Block:    new(big.Int),
	}
	return &chainCfg
}
func CreateExecuteContext(caller kernal.Address) kernal.Context {
	context := kernal.Context{
		Origin:      caller,
		GasPrice:    new(big.Int),
		Coinbase:    kernal.BytesToAddress([]byte("coinbase")),
		GasLimit:    kernal.MaxUint64,
		BlockNumber: new(big.Int),
		Time:        big.NewInt(time.Now().Unix()),
		Difficulty:  new(big.Int),
	}
	return context
}
func CreateVMDefaultConfig() kernal.Config {
	return kernal.Config{
		Debug:                   true,
		Tracer:                  CreateLogTracer(),
		NoRecursion:             false,
		EnablePreimageRecording: false,
	}

}
func CreateExecuteRuntime(caller kernal.Address) *kernal.EVM {
	context := CreateExecuteContext(caller)
	stateDB := MakeNewMockStateDB()
	chainConfig := CreateChainConfig()
	vmConfig := CreateVMDefaultConfig()
	chainHandler := new(ETHChainHandler)

	evm := kernal.NewEVM(context, stateDB, chainHandler, chainConfig, vmConfig)
	return evm
}
