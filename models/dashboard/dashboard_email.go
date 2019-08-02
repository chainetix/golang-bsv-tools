package models

type Email struct {
	From string `json:"from"`
	Recipients []string `json:"recipients"`
	Subject string `json:"subject"`
	Body string `json:"body"`
}
