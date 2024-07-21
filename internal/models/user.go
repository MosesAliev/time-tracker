package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	PassportNumber int    `json:"passportNumber,omitempty" gorm:"primaryKey"`
	Surname        string `json:"surname,omitempty"`
	Name           string `json:"name,omitempty"`
	Patronymic     string `json:"patronymic,omitempty"`
	Adress         string `json:"adress,omitempty"`
	TaskID         int
}
