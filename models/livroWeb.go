package models

type LivroWeb struct {
	LivroNome     string
	EndPoint      string
	Qnt_capitulos int
}

type LivrosWeb struct {
	LivroWeb []LivroWeb
}

func (lsw *LivrosWeb) AddLivro(livro LivroWeb) {
	lsw.LivroWeb = append(lsw.LivroWeb, livro)
}
