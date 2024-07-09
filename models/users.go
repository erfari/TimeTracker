package models

type Users struct {
	ID             string `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	Surname        string `json:"surname,omitempty"`
	Patronymic     string `json:"patronymic,omitempty"`
	Address        string `json:"address,omitempty"`
	PassportSerie  string `json:"passportSerie,omitempty"`
	PassportNumber string `json:"passportNumber,omitempty"`
}

type PassportDocument struct {
	PassportNumber string `json:"passportNumber"`
}
