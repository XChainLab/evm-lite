package kernal

type ChainHandler interface {
	GetBlockHeaderHash(uint64) Hash
}
