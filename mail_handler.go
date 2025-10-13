package main

type Email struct {
	SeqNum uint32 `json:"seqNum"`
	From string `json:"from"`
	Subject string `json:"subject"`
	Date string `json:"date"`
	IsRead bool `json:"isRead"`
	IsStarred bool `json:"isStarred"`
}

// Represents a single attachment's metadata
type Attachment struct {
	Name string `json:"name"` // Matches the `file.name` in your Vue component
	Size int64  `json:"size"` // Matches the `file.size`
	FormattedSize string `json:"formattedSize"`
}

type FullEmail struct {
	SeqNum uint32 `json:"seqNum"`
	From string `json:"from"`
	To string `json:"to"`
	Cc string `json:"cc"`
	Subject string `json:"subject"`
	Body string `json:"body"`
	Attachments []Attachment `json:"attachments"`
	IsEncrypted bool `json:"isEncrypted"`
}
