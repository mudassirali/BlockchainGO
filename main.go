package main

import (
	. "BlockchainGO/src/block"
	wallet "BlockchainGO/src/wallet"
	"fmt"
	"log"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	/*
		b := Block{}
		b.Print()
		bh := b.Hash()
		fmt.Println("Hash 1: ", hex.EncodeToString(bh[:]))

		b.SetNonce(1)
		b.Print()
		bh = b.Hash()
		fmt.Println("Hash 2: ", hex.EncodeToString(bh[:]))
	*/

	walletA := wallet.NewWallet()
	walletB := wallet.NewWallet()

	// Wallet
	t1, s1, pk1 := walletA.SendFunds(walletB.GetBlockchainAddress(), 1.0)
	t2, s2, pk2 := walletB.SendFunds(walletB.GetBlockchainAddress(), 1.0)

	bc := new(Blockchain)
	bc.AddTransaction(t1, s1, pk1)
	bc.AddTransaction(t2, s2, pk2)

	bc.Mining()
	bc.Print()

	fmt.Printf("A %.1f\n", bc.CalculateTotalAmount(walletA.GetBlockchainAddress()))
	fmt.Printf("B %.1f\n", bc.CalculateTotalAmount(walletB.GetBlockchainAddress()))
	/*
		// Blockchain
		blockchain := block.NewBlockchain(walletM.BlockchainAddress())
		isAdded := blockchain.AddTransaction(walletA.BlockchainAddress(), walletB.BlockchainAddress(), 1.0,
			walletA.PublicKey(), t.GenerateSignature())
		fmt.Println("Added? ", isAdded)

		blockchain.Mining()
		blockchain.Print()

		fmt.Printf("A %.1f\n", blockchain.CalculateTotalAmount(walletA.BlockchainAddress()))
		fmt.Printf("B %.1f\n", blockchain.CalculateTotalAmount(walletB.BlockchainAddress()))
		fmt.Printf("M %.1f\n", blockchain.CalculateTotalAmount(walletM.BlockchainAddress()))
	*/
}
