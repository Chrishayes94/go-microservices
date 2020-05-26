package models

import (
	"github.com/shopspring/decimal"
	"time"
)

type BookingRequest struct {
	Hawb              string    `json:"Hawb"`
	Value             float64   `json:"Value"`
	ReadyByDate       time.Time `json:"ReadyByDate"`
	ShipperAddress    Address   `json:"ShipperAddress"`
	CollectionAddress Address   `json:"CollectionAddress"`
	ConsigneeAddress  Address   `json:"ConsigneeAddress"`
	Pieces            []Piece   `json:"Pieces"`
}

func (b *BookingRequest) ValueAsDecimal() *decimal.Decimal {
	value := decimal.NewFromFloat(b.Value)
	return &value
}