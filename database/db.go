package database

import (
	"fmt"
	"log"

	"github.com/isaque-vieira2019/api-biblia/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	stringDeConexao := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}
	err := DB.AutoMigrate(&models.Biblia{}, &models.Livro{}, &models.Versiculo{}, &models.Capitulo{}, &models.SiglaLivro{})

	if err != nil {
		panic("Erro ao Criar as Tabelas")
	}

	insertFirstBiblia()
}

func insertFirstBiblia() {
	var biblia []models.Biblia
	DB.Find(&biblia)

	fmt.Println(len(biblia))
	if len(biblia) == 0 {
		var firstBiblia models.Biblia
		firstBiblia.Nome = "Almeida Corrigida Fiel"
		firstBiblia.Sigla = "ACF"
		firstBiblia.Idioma = "Portugues"
		firstBiblia.Url_Biblia = "https://bibliaestudos.com/acf/"
		DB.Create(&firstBiblia)
	}
}
