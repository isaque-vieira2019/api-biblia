package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/isaque-vieira2019/api-biblia/controllers"
)

func HandleRequest() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	//r.GET("/webScraping/", controllers.StartWebScraping)
	r.GET("/", controllers.ExibePaginaIndex)
	r.GET("/:biblia/livros", controllers.ListarLivros)
	r.GET("/:biblia/livros/:livro/capitulos/:capitulo", controllers.ListarCapituloInteiro)
	r.GET("/:biblia/livros/:livro/capitulos/:capitulo/versiculos/:versiculo", controllers.ListarUmVersiculo)
	r.GET("/:biblia/livros/:livro/capitulos/:capitulo/versiculos/:versiculo/:versiculoEnd", controllers.ListarIntervaloVersiculo)
	r.Run(":8000")
}
