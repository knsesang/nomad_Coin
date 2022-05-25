package main

import (
	"fmt"

	blockchain "_nadochain_com/blockchain"
)

func main() {
	chain := blockchain.Fn_getBlockChain()
	chain.Fn_addBlock("2nd Block")
	chain.Fn_addBlock("3rd Block")
	chain.Fn_addBlock("4th Block")

	for _, blockdata := range chain.Fn_allBlocks() {
		fmt.Printf("data : %s\n", blockdata.Data)
		fmt.Printf("data : %s\n", blockdata.Hash)
		fmt.Printf("data : %s\n\n", blockdata.PrevHash)
	}
}
