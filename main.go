package main

import (
	"crypto/sha256"
	"fmt"
)

type block struct {
	data     string
	hash     string
	prevHash string
}

type blockchain struct {
	blocks []block
}

//   new block add
func (b *blockchain) addBlock(data string) {
	// 블록체인의 길이가 0보다 크면 이전 블록체인의 hash가 존재해야 함
	newBlock := block{data, "", b.getLastHash()}
	hash := sha256.Sum256([]byte(newBlock.data + newBlock.prevHash))
	newBlock.hash = fmt.Sprintf("%x", hash) //  16진수 변환

	// 받아온 b.blocks에 새 블록의 hash를 추가하낟
	b.blocks = append(b.blocks, newBlock)
}

//  all block print
func (b *blockchain) listBlocks() {
	for _, block := range b.blocks {
		fmt.Printf("Data : %s\n", block.data)
		fmt.Printf("Data : %s\n", block.hash)
		fmt.Printf("Data : %s\n", block.prevHash)
	}
}

//  주어진 블록의 이전 블록 hash 가져오기
func (b *blockchain) getLastHash() string {
	if len(b.blocks) > 0 {
		return b.blocks[len(b.blocks)-1].hash
	}
	return ""
}

func main() {
	// genesisBlock : 최초 블럭
	genesisBlock := block{
		"Genesis Block", "", "",
	}

	// Sum256은 data 를 []byte 로 받고 [32]byte 로 리턴한다
	hash := sha256.Sum256([]byte(genesisBlock.data + genesisBlock.prevHash))
	// var hash [32]byte = sha256.Sum256([]byte(genesisBlock.data + genesisBlock.prevHash))
	// fmt.Println(hash)       //  [137 235 10 192 49 166 61 36 33 205 5 162 251 228 31 62 163 95 92 55 18 202 131 156 191 107 133 196 238 7 183 163]
	// fmt.Printf("%x", hash)  //  89eb0ac031a63d2421cd05a2fbe41f3ea35f5c3712ca839cbf6b85c4ee07b7a3
	hexHash := fmt.Sprintf("%x", hash) //  89eb0ac031a63d2421cd05a2fbe41f3ea35f5c3712ca839cbf6b85c4ee07b7a3    콘솔에 출력하지 않고 return

	genesisBlock.hash = hexHash

	fmt.Println(genesisBlock) //  {Genesis Block 89eb0ac031a63d2421cd05a2fbe41f3ea35f5c3712ca839cbf6b85c4ee07b7a3 }

	chain := blockchain{}
	chain.addBlock("Genesis Block")
	// Data : Genesis Block
	// Data : 89eb0ac031a63d2421cd05a2fbe41f3ea35f5c3712ca839cbf6b85c4ee07b7a3
	// Data :

	chain.addBlock("Second Block")
	// Data : Second Block
	// Data : ec6f43cf27c78760cc38a4855b36e83a7f054f5205a751628591d3f505599c08
	// Data : 89eb0ac031a63d2421cd05a2fbe41f3ea35f5c3712ca839cbf6b85c4ee07b7a3

	chain.addBlock("Third Block")
	// Data : Third Block
	// Data : d90e93ce55d3ef3b137b01112c8aad9f304efde1f970c41035c9c7658ad5e8dd
	// Data : ec6f43cf27c78760cc38a4855b36e83a7f054f5205a751628591d3f505599c08

	chain.listBlocks()
}
