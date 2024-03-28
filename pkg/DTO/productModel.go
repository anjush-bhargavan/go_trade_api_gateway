package dto


//Product represents the product data structure.
type Product struct {
	Name             string `json:"name"`
	Category         uint   `json:"category"`
	BasePrice        uint   `json:"base_price"`
	Details          string `json:"details"`
	Description      string `json:"description"`
	Image            string `json:"image"`
	BiddingType      bool   `json:"bidding_type"`
	BiddingStartTime string `json:"bidding_start_time"`
	BiddingEndTime   string `json:"bidding_end_time"`
}


//Category represents the category data structure.
type Category struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}