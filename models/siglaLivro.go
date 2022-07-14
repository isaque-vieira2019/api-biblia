package models

import "gorm.io/gorm"

type SiglaLivro struct {
	gorm.Model
	Sigla       string `JSON:"sigla"`
	Fk_id_livro int
	Livro       Livro `gorm:"foreignKey:Fk_id_livro"`
}
