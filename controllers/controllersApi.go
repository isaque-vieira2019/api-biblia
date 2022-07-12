package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ExibePaginaIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
