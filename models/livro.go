package models

type Livro struct {
	Id           int
	Nome         string
	Autor        string
	Qnt_Capitulo int
	fk_id_biblia int
}
