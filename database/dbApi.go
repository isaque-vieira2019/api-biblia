package database

import "github.com/isaque-vieira2019/api-biblia/models"

func ListarLivrosDB() []models.LivrosApi {
	var livros []models.LivrosApi
	DB.Table("livros").Select("nome, qnt_capitulo").Order("id").Scan(&livros)

	return livros
}

func ListarCapituloInteiroDB(sigla, capNumero string) []models.VersiculosApi {
	var id_capitulo int
	stmt := "SELECT cp.id FROM sigla_livros sl INNER JOIN livros l on sl.fk_id_livro = l.id"
	stmt += " INNER JOIN capitulos cp on l.id = cp.fk_id_livro"
	stmt += " WHERE sl.sigla = ? and cp.n_capitulo = ?"
	DB.Raw(stmt, sigla, capNumero).Scan(&id_capitulo)

	var versiculos []models.VersiculosApi
	DB.Table("versiculos").Where("fk_id_capitulo", id_capitulo).Select("n_versiculo, conteudo").Order("id").Scan(&versiculos)

	return versiculos
}

func ListarUmVersiculoDB(sigla, capNumero, versNumero string) string {
	var versiculo string
	stmt := "SELECT vs.conteudo FROM sigla_livros sl INNER JOIN livros l on sl.fk_id_livro = l.id"
	stmt += " INNER JOIN capitulos cp on l.id = cp.fk_id_livro"
	stmt += " INNER JOIN versiculos vs on cp.id = vs.fk_id_capitulo"
	stmt += " WHERE sl.sigla = ? and cp.n_capitulo = ? and vs.n_versiculo = ?"
	DB.Raw(stmt, sigla, capNumero, versNumero).Scan(&versiculo)

	return versiculo
}

func ListarIntervaloVersiculoDB(sigla, capNumero, versInitNumero, versEndNumero string) []models.VersiculosApi {
	var id_capitulo int
	stmt := "SELECT cp.id FROM sigla_livros sl INNER JOIN livros l on sl.fk_id_livro = l.id"
	stmt += " INNER JOIN capitulos cp on l.id = cp.fk_id_livro"
	stmt += " WHERE sl.sigla = ? and cp.n_capitulo = ?"
	DB.Raw(stmt, sigla, capNumero).Scan(&id_capitulo)

	var versiculos []models.VersiculosApi

	query := DB.Table("versiculos").Where("fk_id_capitulo = ? and n_versiculo BETWEEN ? and ?", id_capitulo, versInitNumero, versEndNumero)
	query.Select("n_versiculo, conteudo").Order("id").Scan(&versiculos)

	return versiculos
}
