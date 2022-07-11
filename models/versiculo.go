package models

type Versiculo struct {
	Id             int
	N_Capitulo     int
	Conteudo       string
	fk_id_Capitulo int
}
