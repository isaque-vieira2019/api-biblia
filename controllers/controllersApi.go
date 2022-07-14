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
	livros := database.ListarLivrosDB()

	c.JSON(http.StatusOK, livros)
}

func ListarCapituloInteiro(c *gin.Context) {
	sigla := c.Params.ByName("livro")
	capNumero := c.Params.ByName("capitulo")

	versiculos := database.ListarCapituloInteiroDB(sigla, capNumero)

	c.JSON(http.StatusOK, versiculos)
}

func ListarUmVersiculo(c *gin.Context) {
	sigla := c.Params.ByName("livro")
	capNumero := c.Params.ByName("capitulo")
	versNumero := c.Params.ByName("versiculo")

	versiculo := database.ListarUmVersiculoDB(sigla, capNumero, versNumero)

	c.JSON(http.StatusOK, versiculo)
}

func ListarIntervaloVersiculo(c *gin.Context) {
	sigla := c.Params.ByName("livro")
	capNumero := c.Params.ByName("capitulo")
	versInitNumero := c.Params.ByName("versiculo")
	versEndNumero := c.Params.ByName("versiculoEnd")

	versiculos := database.ListarIntervaloVersiculoDB(sigla, capNumero, versInitNumero, versEndNumero)

	c.JSON(http.StatusOK, versiculos)
}
