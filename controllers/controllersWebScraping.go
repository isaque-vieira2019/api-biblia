package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

func StartWebScraping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"mensage": "Iniciando Web Scraping",
	})
}

func setLivrosFromWeb(url string) {

	var livro []string
	var endpoint []string

	doc := requestBody(url)

	doc.Find(".group-books-links div li ul li").Each(func(i int, s *goquery.Selection) {
		temp, _ := s.Find("a").Attr("href")
		endpoint = append(endpoint, temp)
		livro = append(livro, s.Find("a").Text())
	})
}

func setCapitulosFromWeb(livro string) {
	url := livro
	doc := requestBody(url)
	quantidade := len(doc.Find(".list-capitulos").Children().Nodes)
	fmt.Println(quantidade)
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
