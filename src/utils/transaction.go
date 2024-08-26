package utils

import (
	"fmt"
	"strings"
)

type Transaction struct {
	senderBlockchainAddress    string  `json:"sender_blockchain_address"`
	recipientBlockchainAddress string  `json:"recipient_blockchain_address"`
	value                      float32 `json:"value"`
}

func NewTransaction(sender string, recipient string, value float32) *Transaction {
	return &Transaction{sender, recipient, value}
}

func (t *Transaction) GetSenderAddress() string {
	return t.senderBlockchainAddress
}

func (t *Transaction) GetRecipientAddress() string {
	return t.recipientBlockchainAddress
}

func (t *Transaction) GetFunds() float32 {
	return t.value
}

func (t *Transaction) Print() {
	fmt.Printf("%s\n", strings.Repeat("-", 40))
	fmt.Printf(" sender_blockchain_address      %s\n", t.senderBlockchainAddress)
	fmt.Printf(" recipient_blockchain_address   %s\n", t.recipientBlockchainAddress)
	fmt.Printf(" value                          %.1f\n", t.value)
}
