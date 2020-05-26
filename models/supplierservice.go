package models

type SupplierService struct {
	Name	string	`json:"Name"`
	Options	[]ServiceField	`json:"Options"`
}
