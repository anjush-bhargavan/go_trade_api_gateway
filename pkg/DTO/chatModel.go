package dto

// Message represents the chat details model
type Message struct {
	UserID     uint   `json:"userId"`
	ReceiverID uint   `json:"receiverId"`
	Message    string `json:"content"`
}
