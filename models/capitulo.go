package models

type Capitulo struct {
	Id              int
	N_Capitulo      string
	fk_id_livro     int
	fk_id_versiculo int
}
