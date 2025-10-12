package main

type Email struct {
	SeqNum uint32 `json:"seqNum"`
	From string `json:"from"`
	Subject string `json:"subject"`
	Date string `json:"date"`
	IsRead bool `json:"isRead"`
	IsStarred bool `json:"isStarred"`
}

type FullEmail struct {
	From string `json:"from"`
	To string `json:"to"`
	Cc string `json:"cc"`
	Subject string `json:"subject"`
	Body string `json:"body"`
	// Attachements - Future Work
}
