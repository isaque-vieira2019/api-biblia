package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isaque-vieira2019/api-biblia/database"
)

func ExibePaginaIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func ListarLivros(c *gin.Context) {
	type livrosApi struct {
		Nome         string
		Qnt_Capitulo int
	}

	var livros []livrosApi
	database.DB.Table("livros").Select("nome, qnt_capitulo").Order("id").Scan(&livros)

	c.JSON(http.StatusOK, livros)
}

func ListarCapituloInteiro(c *gin.Context) {
	capituloSlc := c.Params.ByName("capitulo")
	sigla := c.Params.ByName("livro")

	var id_capitulo int
	stmt := "SELECT cp.id FROM sigla_livros sl INNER JOIN livros l on sl.fk_id_livro = l.id"
	stmt += " INNER JOIN capitulos cp on l.id = cp.fk_id_livro"
	stmt += " WHERE sl.sigla = ? and cp.n_capitulo = ?"
	database.DB.Raw(stmt, sigla, capituloSlc).Scan(&id_capitulo)

	type VersiculosApi struct {
		N_versiculo int
		Conteudo    string
	}

	var versiculos []VersiculosApi
	database.DB.Table("versiculos").Where("fk_id_capitulo", id_capitulo).Select("n_versiculo, conteudo").Order("id").Scan(&versiculos)

	c.JSON(http.StatusOK, versiculos)
}

func ListarUmVersiculo(c *gin.Context) {
	capNumero := c.Params.ByName("capitulo")
	sigla := c.Params.ByName("livro")
	versNumero := c.Params.ByName("versiculo")

	var versiculo string
	stmt := "SELECT vs.conteudo FROM sigla_livros sl INNER JOIN livros l on sl.fk_id_livro = l.id"
	stmt += " INNER JOIN capitulos cp on l.id = cp.fk_id_livro"
	stmt += " INNER JOIN versiculos vs on cp.id = vs.fk_id_capitulo"
	stmt += " WHERE sl.sigla = ? and cp.n_capitulo = ? and vs.n_versiculo = ?"
	database.DB.Raw(stmt, sigla, capNumero, versNumero).Scan(&versiculo)

	c.JSON(http.StatusOK, versiculo)
}

func ListarIntervaloVersiculo(c *gin.Context) {
	capituloSlc := c.Params.ByName("capitulo")
	sigla := c.Params.ByName("livro")
	versInitNumero := c.Params.ByName("versiculo")
	versEndNumero := c.Params.ByName("versiculoEnd")

	var id_capitulo int
	stmt := "SELECT cp.id FROM sigla_livros sl INNER JOIN livros l on sl.fk_id_livro = l.id"
	stmt += " INNER JOIN capitulos cp on l.id = cp.fk_id_livro"
	stmt += " WHERE sl.sigla = ? and cp.n_capitulo = ?"
	database.DB.Raw(stmt, sigla, capituloSlc).Scan(&id_capitulo)

	type VersiculosApi struct {
		N_versiculo int
		Conteudo    string
	}

	var versiculos []VersiculosApi
	database.DB.Table("versiculos").Where("fk_id_capitulo = ? and n_versiculo BETWEEN ? and ?", id_capitulo, versInitNumero, versEndNumero).Select("n_versiculo, conteudo").Order("id").Scan(&versiculos)

	c.JSON(http.StatusOK, versiculos)
}
