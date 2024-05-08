package dto

// Bid Model is for holding the bid details from user
type Bid struct {
	ProductID uint    `json:"product_id"`
	Amount    float64 `json:"amount"`
}
