package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/isaque-vieira2019/api-biblia/actions"
)

func HandleRequest() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	//r.GET("/webScraping/", controllers.StartWebScraping)
	r.GET("/", actions.ExibePaginaIndexAction)
	r.GET("/:biblia/livros", actions.ListarLivrosAction)
	r.GET("/:biblia/livros/:livro/capitulos/:capitulo", actions.ListarCapituloInteiroAction)
	r.GET("/:biblia/livros/:livro/capitulos/:capitulo/versiculos/:versiculo", actions.ListarUmVersiculoAction)
	r.GET("/:biblia/livros/:livro/capitulos/:capitulo/versiculos/:versiculo/:versiculoEnd", actions.ListarIntervaloVersiculoAction)
	r.Run(":8000")
}
