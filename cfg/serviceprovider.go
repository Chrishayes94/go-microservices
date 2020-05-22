package cfg

type FieldType string

const (
	Options FieldType = "Options"
	Text FieldType = "Text"
	Password FieldType = "Password"
)

type IServiceField struct {
	Name		string	`json:"Name"`
	Description	string	`json:"Description"`
	IsRequired	bool	`json:"IsRequired"`
	FieldType	FieldType	`json:"FieldType"`
}

type IServiceOption struct {
	Name		string			`json:"Name"`
	Description	string			`json:"Description"`
	Id			string			`json:"Id"`
	SubFields	[]IServiceField	`json:"SubFields"`
}

type IPasswordField struct {
	Value		string	`json:"Value"`
	IServiceField
}

type ITextField struct {
	Value		string	`json:"Value"`
	IServiceField
}

type IServiceOptionField struct {
	Options		[]IServiceOption	`json:"Options"`
	IServiceField
}

type IServiceProvider struct {
	Name		string			`json:"Name"`
	Options		[]interface{}	`json:"Options"`
}
