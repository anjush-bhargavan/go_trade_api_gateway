package dto

type Bid struct {
	ProductID uint    `json:"product_id"`
	Amount    float64 `json:"amount"`
}