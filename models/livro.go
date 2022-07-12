package models

type Livro struct {
	Id           int
	Nome         string
	Qnt_Capitulo int
	fk_id_biblia int
}
