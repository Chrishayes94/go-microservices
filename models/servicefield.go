package models

import "github.com/chrishayes94/go-microservices/cfg"

type ServiceField struct {
	Name  string        `json:"Name"`
	Type  cfg.FieldType `json:"Type"`
	Value string        `json:"Value"`
}
