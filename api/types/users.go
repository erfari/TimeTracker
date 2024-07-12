package types

type Users struct {
	ID             string `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	Surname        string `json:"surname,omitempty"`
	Patronymic     string `json:"patronymic,omitempty"`
	Address        string `json:"address,omitempty"`
	PassportSerial string `json:"passportSerial,omitempty"`
	PassportNumber string `json:"passportNumber,omitempty"`
}

type PassportDocument struct {
	PassportNumber string `json:"passportNumber"`
}
