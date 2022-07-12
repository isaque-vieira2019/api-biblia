package models

import "gorm.io/gorm"

type Livro struct {
	gorm.Model
	Nome         string `json:"nome"`
	Qnt_Capitulo int    `json:"qnt_capitulos"`
	Fk_id_biblia int
	Biblia       Biblia `gorm:"foreignKey:Fk_id_biblia"`
}
