package models

import "gorm.io/gorm"

type Versiculo struct {
	gorm.Model
	N_Versiculo    int    `json:"n_versiculo"`
	Conteudo       string `json:"conteudo"`
	Fk_id_capitulo int
	Capitulo       Capitulo `gorm:"foreignKey:Fk_id_capitulo"`
}
