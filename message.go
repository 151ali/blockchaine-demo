package main

// Message : utili
type Message struct {
	Message string `json:"message"`
}

// NewMessage :
func NewMessage() Message {
	return Message{}
}

// SetMessage :
func (d *Message) SetMessage(m string) {
	d.Message = m
}

// GetMessage :
func (d *Message) GetMessage() string {
	return d.Message
}
