package models

import (
	"github.com/shopspring/decimal"
)

type Piece struct {
	Depth           int32   `json:"Depth"`
	Height          int32   `json:"Height"`
	Width           int32   `json:"Width"`
	Weight          float64 `json:"Weight"`
	ItemDescription string  `json:"ItemDescription"`
	AgentBarcode	string	`json:"AgentBarcode"`
	NorskBarcode	string	`json:"NorskBarcode"`
}

func (p *Piece) WeightAsDecimal() *decimal.Decimal {
	value := decimal.NewFromFloat(p.Weight)
	return &value
}
