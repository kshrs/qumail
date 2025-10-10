package main

type Email struct {
	SeqNum uint32 `json:"seqNum"`
	From string `json:"from"`
	Subject string `json:"subject"`
	Date string `json:"date"`
	IsRead bool `json:"isRead"`
}
