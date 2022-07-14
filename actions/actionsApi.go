package actions

import (
	"github.com/gin-gonic/gin"
	"github.com/isaque-vieira2019/api-biblia/controllers"
)

func ExibePaginaIndexAction(c *gin.Context) {
	controllers.ExibePaginaIndex(c)
}

func ListarLivrosAction(c *gin.Context) {
	controllers.ListarLivros(c)
}

func ListarCapituloInteiroAction(c *gin.Context) {
	controllers.ListarCapituloInteiro(c)
}

func ListarUmVersiculoAction(c *gin.Context) {
	controllers.ListarUmVersiculo(c)
}

func ListarIntervaloVersiculoAction(c *gin.Context) {
	controllers.ListarIntervaloVersiculo(c)
}
