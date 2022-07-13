package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/isaque-vieira2019/api-biblia/controllers"
)

func HandleRequest() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	//r.GET("/webScraping/", controllers.StartWebScraping) rota desativada por enquanto
	r.GET("/", controllers.ExibePaginaIndex)
	r.Run(":8000")
}
