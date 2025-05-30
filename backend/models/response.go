package models

type TransactionResponse struct {
	Type     string  `json:"type"`
	Amount   float32 `json:"amount"`
	Ticker   string  `json:"ticker"`
	Quantity int     `json:"quantity"`
	Date     string  `json:"date"`
}

type ProfileRespone struct {
	Id           int                   `json:"id"`
	Name         string                `json:"name"`
	Email        string                `json:"email"`
	Phone        string                `json:"phone"`
	Balance      float32               `json:"balance"`
	Stocks       []PortfolioStock      `json:"stocks"`
	Transactions []TransactionResponse `json:"transactions"`
}
