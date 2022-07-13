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

func setLivrosFromWeb(url string) {
	lsw := models.LivrosWeb{}

	doc := requestBody(url)
	doc.Find(".group-books-links div li ul li").Each(func(i int, s *goquery.Selection) {
		var livroWeb models.LivroWeb
		livroWeb.LivroNome = s.Find("a").Text()
		temp, _ := s.Find("a").Attr("href")
		livroWeb.EndPoint = temp
		livroWeb.Qnt_capitulos = getCapitulosFromWeb(temp)

		lsw.AddLivro(livroWeb)

	})

	for _, livro_web := range lsw.LivroWeb {
		var livro models.Livro
		livro.Nome = livro_web.LivroNome
		livro.Qnt_Capitulo = livro_web.Qnt_capitulos
		livro.Fk_id_biblia = 1

		database.DB.Create(&livro)

		setCapitulosFromWeb(livro_web.EndPoint, livro.Qnt_Capitulo, int(livro.ID))
	}
}

func getCapitulosFromWeb(livro string) int {
	url := livro
	doc := requestBody(url)
	quantidade := len(doc.Find(".list-capitulos").Children().Nodes)
	return quantidade
}

func setCapitulosFromWeb(url string, qnt int, id_livro int) {
	for i := 1; i <= qnt; i++ {
		var capitulo models.Capitulo
		capitulo.N_Capitulo = i
		capitulo.Fk_id_livro = id_livro

		database.DB.Create(&capitulo)

		setVersiculosFromWeb(url, capitulo)
	}
}

func setVersiculosFromWeb(livro string, cap models.Capitulo) {
	url := livro + strconv.Itoa(cap.N_Capitulo)
	doc := requestBody(url)

	doc.Find("#mvp-content-main p").Each(func(i int, s *goquery.Selection) {
		var versiculo models.Versiculo
		versiculo.N_Versiculo = i + 1
		versiculo.Conteudo = s.Text()
		versiculo.Fk_id_capitulo = cap.N_Capitulo

		database.DB.Create(&versiculo)
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
