package models

type AgentAddress struct {
	Company      string `json:"Company"`
	ContactName  string `json:"ContactName"`
	Address1     string `json:"Address1"`
	Address2     string `json:"Address2"`
	Address3     string `json:"Address3"`
	City         string `json:"City"`
	Email        string `json:"Email"`
	PhoneNumber  string `json:"PhoneNumber"`
	Zipcode      string `json:"Zipcode"`
	CountryCode  string `json:"CountryCode"`
	Country      string `json:"Country"`
	Fax          string `json:"Fax"`
	County       string `json:"County"`
	CountyCode   string `json:"CountyCode"`
	MobileNumber string `json:"MobileNumber"`
	CountryId	int16	`json:"CountryId"`
	VatNumber	string	`json:"VatNumber"`
}

type Address struct {
	Company      string `json:"Company"`
	ContactName  string `json:"ContactName"`
	Line1        string `json:"Address1"`
	Line2        string `json:"Address2"`
	Line3        string `json:"Address3"`
	City         string `json:"City"`
	Email        string `json:"Email"`
	PhoneNumber  string `json:"PhoneNumber"`
	Zipcode      string `json:"Zipcode"`
	CountryCode  string `json:"CountryCode"`
	Country      Country `json:"Country"`
	Fax          string `json:"Fax"`
	Division     string `json:"County"`
	MobileNumber string `json:"MobileNumber"`
}

type Country struct {
	Id           int32  `json:"Id"`
	Isocode      string `json:"Isocode"`
	Name         string `json:"Name"`
	CurrencyCode string `json:"CurrencyCode"`
	IsEuMember   bool   `json:"IsEuMember"`
	IsoNumber    int    `json:"IsoNumber"`
}