package models

type User struct {
	PassportNumber int    `json:"passportNumber,omitempty" gorm:"primaryKey"`
	Surname        string `json:"surname,omitempty"`
	Name           string `json:"name,omitempty"`
	Patronymic     string `json:"patronymic,omitempty"`
	Adress         string `json:"adress,omitempty"`
	TaskID         int    `json:"-"`
}
