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
}
