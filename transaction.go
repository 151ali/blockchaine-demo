package main

// Transaction :
type Transaction struct {
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	Amount   int    `json:"amount"`
}

// NewTransaction :
func NewTransaction(s, r string, a int) *Transaction {
	return &Transaction{
		Sender:   s,
		Receiver: r,
		Amount:   a,
	}
}
