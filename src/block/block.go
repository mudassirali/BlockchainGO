package block

import (
	. "BlockchainGO/src/utils"
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/goccy/go-json"
)

type Block struct {
	Timestamp    int64          `json:"timestamp"`
	Nonce        int            `json:"nonce"`
	PreviousHash [32]byte       `json:"previous_hash"`
	Transactions []*Transaction `json:"transactions"`
}

func NewBlock(nonce int, previousHash [32]byte, transactions []*Transaction) *Block {
	b := new(Block)
	b.Timestamp = time.Now().UnixNano()
	b.Nonce = nonce
	b.PreviousHash = previousHash
	b.Transactions = transactions
	return b
}

func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	fmt.Println("JSON Hash: ", string(m))
	return sha256.Sum256([]byte(m))
}

func (b *Block) Print() {
	fmt.Printf("timestamp       %d\n", b.Timestamp)
	fmt.Printf("nonce           %d\n", b.Nonce)
	fmt.Printf("previous_hash   %x\n", b.PreviousHash)
	for _, t := range b.Transactions {
		t.Print()
	}
}
