package models

import "gorm.io/gorm"

type Biblia struct {
	gorm.Model
	Nome   string `json:"nome"`
	Sigla  string `json:"sigla"`
	Idioma string `json:"idioma"`
}
