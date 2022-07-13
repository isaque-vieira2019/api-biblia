package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isaque-vieira2019/api-biblia/database"
	"github.com/isaque-vieira2019/api-biblia/models"
)

func ExibePaginaIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func ListarLivros(c *gin.Context) {
	var livros []models.Livro
	database.DB.Find(&livros)

	c.JSON(http.StatusOK, livros)
}

func ListarCapituloInteiro(c *gin.Context) {
	capituloSlc := c.Params.ByName("capitulo")
	sigla := c.Params.ByName("livro")

	type Result struct {
		n_versiculo int
		conteudo    string
	}
	var result []Result

	stms := "select	vs.n_versiculo,	vs.conteudo from sigla_livros sl inner join"
	stms += " livros l on sl.fk_id_livro = l.id inner join capitulos cp"
	stms += " on l.id = cp.fk_id_livro inner join versiculos vs"
	stms += " on cp.id = fk_id_capitulo where sl.sigla LIKE ? and cp.n_capitulo = ? "
	database.DB.Raw(stms, sigla, capituloSlc).Scan(&result)

	fmt.Println(result)
}

func ListarUmVersiculo(c *gin.Context) {

}

func ListarIntervaloVersiculo(c *gin.Context) {

}
