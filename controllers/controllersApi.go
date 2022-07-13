package controllers

import (
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
	/*var capitulo models.Capitulo
	siglaLivro := c.Params.ByName("livro")*/

}

func ListarUmVersiculo(c *gin.Context) {

}

func ListarIntervaloVersiculo(c *gin.Context) {

}
