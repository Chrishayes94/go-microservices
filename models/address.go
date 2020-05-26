package models

type Address struct {
	Company			string	`json:"Company"`
	ContactName		string 	`json:"ContactName"`
	Address1		string	`json:"Address1"`
	Address2		string	`json:"Address2"`
	Address3		string	`json:"Address3"`
	City			string	`json:"City"`
	Email			string	`json:"Email"`
	PhoneNumber		string	`json:"PhoneNumber"`
	Zipcode			string	`json:"Zipcode"`
	CountryCode		string	`json:"CountryCode"`
	Country			string	`json:"Country"`
	Fax				string	`json:"Fax"`
	County			string 	`json:"County"`
	CountyCode		string	`json:"CountyCode"`
	MobileNumber	string	`json:"MobileNumber"`
}
