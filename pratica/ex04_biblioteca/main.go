// ex 4 - Biblioteca de Livros
package main

import (
	"fmt"
	"sort"
	"strings"
)

type Pesquisavel interface {
	Contem(termo string) bool
}

type Livro struct {
	Titulo        string
	Autor         string
	Genero        string
	AnoPublicacao int
}

func (l Livro) Contem(termo string) bool {
	return strings.Contains(strings.ToLower(l.Titulo), strings.ToLower(termo)) || strings.Contains(strings.ToLower(l.Autor), strings.ToLower(termo))
}

type Biblioteca []Livro

func (b *Biblioteca) Adicionar(livros ...Livro) {
	*b = append(*b, livros...)
}

func (b *Biblioteca) Pesquisar(termo string) []Livro {
	var resultados []Livro
	for _, livro := range *b {
		if livro.Contem(termo) {
			resultados = append(resultados, livro)
		}
	}
	return resultados
}

func (b *Biblioteca) PorGenero(genero string) []Livro {
	var resultados []Livro
	for _, livro := range *b {
		if strings.ToLower(livro.Genero) == strings.ToLower(genero) {
			resultados = append(resultados, livro)
		}
	}
	return resultados
}

// Recebe o slice de livros, e uma funcao que compara dois elementos pelos
// indices i e j.
func (b *Biblioteca) MaisRecentes(n int) []Livro {
	sort.Slice(*b, func(i, j int) bool {
		return (*b)[i].AnoPublicacao > (*b)[j].AnoPublicacao
	})

	if len(*b) < n {
		return *b
	}

	return (*b)[:n]
}

func main() {

	l1 := Livro{Titulo: "O Senhor dos Anéis - A Sociedade do Anel", Autor: "J.R.R. Tolkien", Genero: "Fantasia", AnoPublicacao: 1954}

	l2 := Livro{Titulo: "O Senhor dos Anéis - As Duas Torres ", Autor: "J.R.R. Tolkien", Genero: "Fantasia", AnoPublicacao: 1954}

	l3 := Livro{Titulo: "O Senhor dos Anéis - O Retorno do Rei", Autor: "J.R.R. Tolkien", Genero: "Fantasia", AnoPublicacao: 1955}

	l4 := Livro{Titulo: "A Study in Scarlet", Autor: "Arthur Conan Doyle", Genero: "Romance Policial", AnoPublicacao: 1888}

	l5 := Livro{Titulo: "The Hound of the Baskervilles", Autor: "Arthur Conan Doyle", Genero: "Romance Policial", AnoPublicacao: 1902}

	l6 := Livro{Titulo: "The Sign of Four", Autor: "Arthur Conan Doyle", Genero: "Romance Policial", AnoPublicacao: 1890}

	l7 := Livro{Titulo: "The Adventures of Sherlock Holmes", Autor: "Arthur Conan Doyle", Genero: "Romance Policial", AnoPublicacao: 1892}

	l8 := Livro{Titulo: "O Código Da Vinci", Autor: "Dan Brown", Genero: "Mistério", AnoPublicacao: 2003}

	var bib Biblioteca

	fmt.Println("\n(Adicionar) Adicionando 8 livros à biblioteca...\n")
	bib.Adicionar(l1, l2, l3, l4, l5, l6, l7, l8)

	fmt.Println("\nTodos os livros:\n")
	fmt.Printf("%v\n", bib)

	fmt.Println("\n(Pesquisar) Livros do autor J.R.R. Tolkien:\n")
	for _, l := range bib.Pesquisar("Tolkien") {
		fmt.Printf("  %s (%d)\n", l.Titulo, l.AnoPublicacao)
	}

	fmt.Println("\n(MaisRecentes(3)) Os livros mais recentes são:\n")
	for _, l := range bib.MaisRecentes(3) {
		fmt.Printf("  %s (%d)\n", l.Titulo, l.AnoPublicacao)
	}

	fmt.Println("\n(PorGenero) Livros de Mistério:\n")
	for _, l := range bib.PorGenero("Mistério") {
		fmt.Printf("  %s (%s)\n", l.Titulo, l.Genero)
	}

}
