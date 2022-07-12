package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"github.com/isaque-vieira2019/api-biblia/database"
	"github.com/isaque-vieira2019/api-biblia/models"
)

func StartWebScraping(c *gin.Context) {
	var biblias []models.Biblia
	database.DB.Find(&biblias)
	for _, biblia := range biblias {
		setLivrosFromWeb(biblia.Url_Biblia)
	}
}

type livroWeb struct {
	livroNome     string
	endPoint      string
	qnt_capitulos int
}

type livrosWeb struct {
	livroWeb []livroWeb
}

func (lsw *livrosWeb) AddLivro(livro livroWeb) {
	lsw.livroWeb = append(lsw.livroWeb, livro)
}
func setLivrosFromWeb(url string) {
	lsw := livrosWeb{}

	doc := requestBody(url)
	doc.Find(".group-books-links div li ul li").Each(func(i int, s *goquery.Selection) {
		var livroWeb livroWeb
		livroWeb.livroNome = s.Find("a").Text()
		temp, _ := s.Find("a").Attr("href")
		livroWeb.endPoint = temp
		livroWeb.qnt_capitulos = getCapitulosFromWeb(temp)

		lsw.AddLivro(livroWeb)

	})

	for _, livro_web := range lsw.livroWeb {
		var livro models.Livro
		livro.Nome = livro_web.livroNome
		livro.Qnt_Capitulo = livro_web.qnt_capitulos
		livro.Fk_id_biblia = 1

		database.DB.Create(&livro)
	}
}

func getCapitulosFromWeb(livro string) int {
	url := livro
	doc := requestBody(url)
	quantidade := len(doc.Find(".list-capitulos").Children().Nodes)
	return quantidade
}

func setVersiculosFromWeb(livro string, capitulo int) {
	url := livro + strconv.Itoa(capitulo)
	doc := requestBody(url)

	var versiculo []string
	doc.Find("#mvp-content-main p").Each(func(i int, s *goquery.Selection) {
		versiculo = append(versiculo, s.Text())
	})

}

func requestBody(url string) *goquery.Document {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	return doc
}
