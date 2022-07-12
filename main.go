package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

func getLivrosFromWeb(url string) {

	var livro []string
	var endpoint []string

	doc := requestBody(url)

	doc.Find(".group-books-links div li ul li").Each(func(i int, s *goquery.Selection) {
		temp, _ := s.Find("a").Attr("href")
		endpoint = append(endpoint, temp)
		livro = append(livro, s.Find("a").Text())
	})
}

func getCapitulosFromWeb(livro string) {
	url := livro
	doc := requestBody(url)
	quantidade := len(doc.Find(".list-capitulos").Children().Nodes)
	fmt.Println(quantidade)
}

func getVersiculosFromWeb(livro string, capitulo int) {
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

func main() {
	getLivrosFromWeb("https://bibliaestudos.com/acf/")
	getCapitulosFromWeb("https://bibliaestudos.com/acf/apocalipse/")
	getVersiculosFromWeb("https://bibliaestudos.com/acf/genesis/", 1)
}

/*
import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	fmt.Println("Sejam Bem vindos a Api da Biblia")
	getLivrosFromWeb()
}

func getLivrosFromWeb() {
	res, err := http.Get("https://www.bkjfiel.com.br/")

	if err != nil {
		log.Fatal("Error " + err.Error())
		return
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		log.Fatal("Error " + err.Error())
		return
	}

	//var livros []string

	doc.Find(".sc-hKFxyN").Each(func(_ int, s *goquery.Selection) {
		name := s.Text()
		fmt.Println(name)
	})
}*/
