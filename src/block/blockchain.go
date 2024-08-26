package block

import (
	. "BlockchainGO/src/utils"
	"crypto/ecdsa"
	"crypto/sha256"
	"fmt"
	"log"
	"strings"

	"github.com/parinpan/magicjson"
)

const (
	MINING_DIFFICULTY = 3
	MINING_SENDER     = "THE BLOCKCHAIN"
	MINING_REWARD     = 1.0
)

type Blockchain struct {
	transactionPool   []*Transaction
	chain             []*Block
	blockchainAddress string
}

func (bc *Blockchain) CreateBlock(nonce int, previousHash [32]byte) *Block {
	fmt.Println("bc.Transaction Pool: ", bc.transactionPool)
	b := NewBlock(nonce, previousHash, bc.transactionPool)
	bc.chain = append(bc.chain, b)
	bc.transactionPool = []*Transaction{}
	return b
}

func (bc *Blockchain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("%s Chain %d %s\n", strings.Repeat("=", 25), i,
			strings.Repeat("=", 25))
		block.Print()
	}
	fmt.Printf("%s\n", strings.Repeat("*", 25))
}

func (bc *Blockchain) AddTransaction(t *Transaction, s *Signature, senderPK *ecdsa.PublicKey) bool {

	if t.GetSenderAddress() == MINING_SENDER {
		bc.transactionPool = append(bc.transactionPool, t)
		return true
	}

	if bc.VerifyTransactionSignature(t, s, senderPK) {
		/*
			if bc.CalculateTotalAmount(sender) < value {
				log.Println("ERROR: Not enough balance in a wallet")
				return false
			}
		*/
		bc.transactionPool = append(bc.transactionPool, t)
		return true
	} else {
		log.Println("ERROR: Verify Transaction")
	}
	return false

}

func (bc *Blockchain) VerifyTransactionSignature(t *Transaction, s *Signature, senderPK *ecdsa.PublicKey) bool {
	m, _ := magicjson.Marshal(t)
	h := sha256.Sum256([]byte(m))
	return ecdsa.Verify(senderPK, h[:], s.R, s.S)
}

func (bc *Blockchain) Mining() bool {
	t := NewTransaction(MINING_SENDER, bc.blockchainAddress, MINING_REWARD)
	bc.AddTransaction(t, nil, nil)
	previousHash := bc.LastBlock().Hash()
	nonce := bc.ProofOfWork(previousHash)
	bc.CreateBlock(nonce, previousHash)
	log.Println("action=mining, status=success")
	return true
}

func (bc *Blockchain) LastBlock() *Block {
	fmt.Println(len(bc.chain))
	if len(bc.chain)-1 <= 0 {
		return &Block{nonce: 0}
	} else {
		return bc.chain[len(bc.chain)-1]
	}

}

func (bc *Blockchain) ProofOfWork(previousHash [32]byte) int {
	transactions := bc.transactionPool
	zeros := strings.Repeat("0", MINING_DIFFICULTY)
	nonce := 0

	for {
		guessHashStr := fmt.Sprintf("%x", NewBlock(nonce, previousHash, transactions).Hash())
		if guessHashStr[:MINING_DIFFICULTY] == zeros {
			return nonce
		}
		nonce += 1
	}
}

func (bc *Blockchain) CalculateTotalAmount(blockchainAddress string) float32 {
	var totalAmount float32 = 0.0
	for _, b := range bc.chain {
		for _, t := range b.transactions {
			value := t.GetFunds()
			if blockchainAddress == t.GetRecipientAddress() {
				totalAmount += value
			}

			if blockchainAddress == t.GetSenderAddress() {
				totalAmount -= value
			}
		}
	}
	return totalAmount
}
