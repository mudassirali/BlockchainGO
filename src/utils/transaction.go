package utils

type Transaction struct {
	senderBlockchainAddress    string  `json:"sender_blockchain_address"`
	recipientBlockchainAddress string  `json:"recipient_blockchain_address"`
	value                      float32 `json:"value"`
}

func NewTransaction(sender string, recipient string, value float32) *Transaction {
	return &Transaction{sender, recipient, value}
}
