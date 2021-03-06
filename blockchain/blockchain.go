package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

type Cls_Block struct {
	Data     string
	Hash     string
	PrevHash string
}

type Cls_BlockChain struct {
	Blocks []*Cls_Block
}

var (
	b    *Cls_BlockChain
	once sync.Once //  1번만 실행
)

// 블록 추가
func (b *Cls_BlockChain) Fn_addBlock(data string) {
	b.Blocks = append(b.Blocks, fn_createBlock(data))
}

// new hash 계산
func (b *Cls_Block) fn_calculateHash() {
	hash := sha256.Sum256([]byte(b.Data + b.PrevHash))
	b.Hash = fmt.Sprintf("%x", hash)
}

// 마지막 hash 찾기
func fn_getLastHash() string {
	totalBlocks := len(Fn_getBlockChain().Blocks)
	if totalBlocks == 0 {
		return ""
	}
	return Fn_getBlockChain().Blocks[totalBlocks-1].Hash
}

// 블록 생성
func fn_createBlock(data string) *Cls_Block {
	newBlock := Cls_Block{data, "", fn_getLastHash()}
	newBlock.fn_calculateHash()
	return &newBlock
}

// 블록체인 정보 조회
func Fn_getBlockChain() *Cls_BlockChain {
	if b == nil {
		once.Do(func() {
			b = &Cls_BlockChain{}

			// b.blocks = append(b.blocks, fn_createBlock("1st Block"))
			// ↓
			b.Fn_addBlock("1st Block")
		})
	}
	return b
}

// 전체 블록 조회
func (b *Cls_BlockChain) Fn_allBlocks() []*Cls_Block {
	return b.Blocks
}
