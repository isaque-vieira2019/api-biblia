package models

import "gorm.io/gorm"

type Versiculo struct {
	gorm.Model
	N_Versiculo    int
	Conteudo       string
	Fk_id_capitulo int
	Capitulo       Capitulo `gorm:"foreignKey:Fk_id_capitulo"`
}
