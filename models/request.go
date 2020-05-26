package models

import (
	"github.com/shopspring/decimal"
	"time"
)

type BookingRequest struct {
	Hawb              string          `json:"Hawb"`
	Value             float64         `json:"Value"`
	ReadyByDate       time.Time       `json:"ReadyByDate"`
	ShipperAddress    AgentAddress         `json:"ShipperAddress"`
	CollectionAddress Address         `json:"CollectionAddress"`
	ConsigneeAddress  Address         `json:"ConsigneeAddress"`
	Pieces            []Piece         `json:"Pieces"`
	Plan              SupplierService `json:"Plan"`
	BpRelationshipId  int64           `json:"BpRelationshipId"`
	OwnerId           int64           `json:"OwnerId"`
}

func (b *BookingRequest) ValueAsDecimal() *decimal.Decimal {
	value := decimal.NewFromFloat(b.Value)
	return &value
}

func (b *BookingRequest) IsEmaCustomer() bool {
	return b.OwnerId == 992
}

func (b *BookingRequest) IsLhrCustomer() bool {
	return b.OwnerId == 1
}

func (b *BookingRequest) IsSupplerChainCustomer() bool {
	return b.OwnerId == 2714
}

func (b *BookingRequest) IsWhiteLabelAppCustomer() bool {
	return b.OwnerId == 3612
}