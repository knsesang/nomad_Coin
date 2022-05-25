package main

import (
	"crypto/sha256"
	"fmt"
)

// 블록 선언
type clsBlock struct {
	data     string
	hash     string
	prevHash string
}

// 블록체인 선언
type clsBlockChain struct {
	blocks []clsBlock
}

// 이전 블록의 hash 가져오기
func (b *clsBlockChain) fn_getLastHash() string {
	if len(b.blocks) > 0 {
		return b.blocks[len(b.blocks)-1].hash // 블록이 있다면 바로 직전의 블록 hash 를 가져온다
	}
	return ""
}

// 블록 추가
func (b *clsBlockChain) fn_addBlock(data string) {
	newBlock := clsBlock{data, "", b.fn_getLastHash()}               //	새 블록 생성
	hash := sha256.Sum256([]byte(newBlock.data + newBlock.prevHash)) //	// sha256.Sum256 ( data []byte, [Size]byte)
	newBlock.hash = fmt.Sprintf("%x", hash)                          //	hash 를 16진수로 바꾸어 저장
	b.blocks = append(b.blocks, newBlock)
}

// * 포인터에 있는 값
func (b *clsBlockChain) fn_listBlocks() {
	for _, block := range b.blocks {
		fmt.Printf("data : %s\n", block.data)
		fmt.Printf("hash : %s\n", block.hash)
		fmt.Printf("prevHash : %s\n\n", block.prevHash)

	}
}

func main() {
	chain := clsBlockChain{}
	chain.fn_addBlock("1st Block")
	chain.fn_addBlock("2nd Block")
	chain.fn_addBlock("4rdt Block")

	chain.fn_listBlocks()
}
