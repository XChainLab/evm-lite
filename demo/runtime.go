package main

import (
	"evm/kernel"
	"math/big"
	"time"
)

func CreateLogTracer() *kernel.StructLogger {
	logConf := kernel.LogConfig{
		DisableMemory:  false,
		DisableStack:   false,
		DisableStorage: false,
		Debug:          false,
		Limit:          0,
	}
	return kernel.NewStructLogger(&logConf)

}
func CreateChainConfig() *kernel.ChainConfig {
	chainCfg := kernel.ChainConfig{
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
func CreateExecuteContext(caller kernel.Address) kernel.Context {
	context := kernel.Context{
		Origin:      caller,
		GasPrice:    new(big.Int),
		Coinbase:    kernel.BytesToAddress([]byte("coinbase")),
		GasLimit:    kernel.MaxUint64,
		BlockNumber: new(big.Int),
		Time:        big.NewInt(time.Now().Unix()),
		Difficulty:  new(big.Int),
	}
	return context
}
func CreateVMDefaultConfig() kernel.Config {
	return kernel.Config{
		Debug:                   true,
		Tracer:                  CreateLogTracer(),
		NoRecursion:             false,
		EnablePreimageRecording: false,
	}

}
func CreateExecuteRuntime(caller kernel.Address) *kernel.EVM {
	context := CreateExecuteContext(caller)
	stateDB := MakeNewMockStateDB()
	chainConfig := CreateChainConfig()
	vmConfig := CreateVMDefaultConfig()
	chainHandler := new(ETHChainHandler)

	evm := kernel.NewEVM(context, stateDB, chainHandler, chainConfig, vmConfig)
	return evm
}
