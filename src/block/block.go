package block

import (
	. "BlockchainGO/src/utils"
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/parinpan/magicjson"
)

type Block struct {
	timestamp    int64          `json:"timestamp"`
	nonce        int            `json:"nonce"`
	previousHash [32]byte       `json:"previous_hash"`
	transactions []*Transaction `json:"transactions"`
}

func NewBlock(nonce int, previousHash [32]byte, transactions []*Transaction) *Block {
	b := new(Block)
	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previousHash
	b.transactions = transactions
	return b
}

func (b *Block) Hash() [32]byte {
	m, _ := magicjson.Marshal(b)
	//fmt.Println("JSON Hash: ", string(m))
	return sha256.Sum256([]byte(m))
}

func (b *Block) Print() {
	fmt.Printf("timestamp       %d\n", b.timestamp)
	fmt.Printf("nonce           %d\n", b.nonce)
	fmt.Printf("previous_hash   %x\n", b.previousHash)
	for _, t := range b.transactions {
		t.Print()
	}
}

func (b *Block) SetNonce(nonce int) *Block {
	b.nonce = nonce
	return b
}
