## 概览
---
该项目旨在提供一个以太坊智能合约运行的最小的环境。项目的构建致力于减少区块链开发过程中对以太坊智能合约兼容的成本。新的区块链项目可以通过在程序中构建该项目提供的EVM实例，并实现相关接口快速准确的达到兼容以太坊智能合约的目的。项目主要工作在于将以太坊项目中的虚拟机部分代码抽离出来，精简并合并虚拟机部分代码的依赖，最后实现一个精简的，可复用的以太坊智能合约运行环境。

### 项目依赖
---
该项目无其他任何外部依赖，可以直接下载使用和进行二次开发,这也是该项目的主要目的。
### 具体实现
---
工程下主要有三个目录：
* crypto为加密函数库，函数库来源于go-ethereum,这部分单独出目录
* kernel为以太坊虚拟机核心代码，实现了智能合约的运行环境，代码来自go-ethereum
* demo为一个具体的使用示例

通过demo我们来演示如何让你的区块链支持以太坊智能合约</br>
#### 第一步实现数据访问接口
由于不同区块链底层依赖的数据存储不同，而以太坊智能合约中有对存储的操作，因此这里我们需要实现数据访问接口，接口的描述见文件kernel/statedb.go。
demo中我们实现了其中的部分接口,具体见mockstatedb.go，这里需要说明一下，demo中实现的是以太坊智能合约运行必须实现的接口，其他接口可以考虑不实现，必要的接口函数为如下：
```
GetCode(address kernel.Address) []byte
GetCodeHash(kernel.Address) kernel.Hash
SetCode(address kernel.Address, data []byte)
GetCodeSize(address kernel.Address) int
Exist(kernel.Address) bool
Empty(kernel.Address) bool
//关于snapshot的接口需要根据具体情况进行实现
RevertToSnapshot(int)                                             
Snapshot() int
HaveSufficientBalance(kernel.Address, *big.Int) bool
TransferBalance(kernel.Address, kernel.Address, *big.Int)
```
除此之外还要实现一个链访问的接口，具体见kernel/chain.go,这里只需要实现一个接口函数即可
```
GetBlockHeaderHash(uint64) kernel.Hash
```
#### 第二步创建EVM执行实例
具体见demo/runtime.go，这里主要工作是初始化相关的配置，该项目的原则上保留了以太坊的相关配置，使用者可以根据自己的情况设置其中的具体数值，demo中采用的均是默认值，使用者可以进行参考，创建EVM部分的代码如下：
```
func CreateExecuteRuntime(caller kernel.Address) *kernel.EVM {
    context := CreateExecuteContext(caller)
    stateDB := MakeNewMockStateDB()
    chainConfig := CreateChainConfig()
    vmConfig := CreateVMDefaultConfig()
    chainHandler := new(ETHChainHandler)

    evm := kernel.NewEVM(context, stateDB, chainHandler, chainConfig, vmConfig)
    return evm
}
```
#### 第三部调用智能合约
在第二步中我们创建了EVM的运行实例，这里我们通过调用EVM的Call函数直接运行代码的方式来运行智能合约
```
HexTestCode := "6060604052600a8060106000396000f360606040526008565b00"
TestInput := []byte("Contract")
TestCallerAddress := []byte("TestAddress")
TestContractAddress := []byte("TestContract")
calleraddress := kernel.BytesToAddress(TestCallerAddress)
contractaddress := kernel.BytesToAddress(TestContractAddress)
evm := CreateExecuteRuntime(calleraddress)
evm.StateDBHandler.CreateAccount(contractaddress)
evm.StateDBHandler.SetCode(contractaddress, kernel.Hex2Bytes(HexTestCode))
caller := kernel.AccountRef(evm.Origin)
ret, _, err := evm.Call(
    caller,
    contractaddress,
    TestInput,
    evm.GasLimit,
    new(big.Int))
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(ret)
}
```
这里我们直接将代码传递给了EVM，目前EVM对外的接口保留源代码中的各个接口，可以通过调用Create函数来实现创建一个智能合约。
#### 编译运行
执行上面的demo十分的简单主要执行以下的几步操作即可：
* 确认你的机器上有golang的编译环境
* git clone 代码到$GOPATH/src/evm下
* 进入demo文件夹，执行go build命令
* 运行demo即可

## 其他重要说明
* 由于是EVM的精简版，代码上已经尽量的做了删减，但考虑到最大的兼容，因此保留了几乎所有的针对以太坊的配置设置,具体的配置可以根据实际集成的链进行调整，可以参考demo/runtime.go进行调整
* 数据接口需要根据你的项目进行实现，同时注意实现唯一的链访问接口

## 后续计划
* 完善项目单元测试和示例说明
* 抽离interpreter部分实现，更加的通用的解释器，确定指令集合
* gas模型剥离，提供resource的消费接口
* 通用的StateDB接口设计
.....

## License
项目采用License为[License](https://www.gnu.org/licenses/lgpl-3.0.en.html)
