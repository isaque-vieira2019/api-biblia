package models

import "gorm.io/gorm"

type Livro struct {
	gorm.Model
	Nome         string
	Qnt_Capitulo int
	Fk_id_biblia int
	Biblia       Biblia `gorm:"foreignKey:Fk_id_biblia"`
}
