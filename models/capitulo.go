package models

import "gorm.io/gorm"

type Capitulo struct {
	gorm.Model
	N_Capitulo  int `json:"n_capitulo"`
	Fk_id_livro int
	Livro       Livro `gorm:"foreignKey:Fk_id_livro"`
}
