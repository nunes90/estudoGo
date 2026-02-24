package main

import (
	"fmt"
	"sort"
)

type Categoria int

const (
	Eletronico Categoria = iota
	Roupa
	Alimento
)

// Metodo String() - Go chama isso automaticamente no fmt.Printf com %v
func (c Categoria) String() string {
	switch c {
	case Eletronico:
		return "Eletronico"
	case Roupa:
		return "Roupa"
	case Alimento:
		return "Alimento"
	default:
		return "Desconhecido"
	}
}

type Item struct {
	Preco      float64
	Quantidade int
	Categoria  Categoria
}

type Inventario map[string]Item

func main() {
	// 5 produtos
	produtos := Inventario{
		"Notebook":  {Preco: 3500.00, Quantidade: 2, Categoria: Eletronico},
		"Camiseta":  {Preco: 89.90, Quantidade: 15, Categoria: Roupa},
		"Arroz":     {Preco: 25.00, Quantidade: 30, Categoria: Alimento},
		"Headphone": {Preco: 350.00, Quantidade: 1, Categoria: Eletronico},
		"Calça":     {Preco: 120.00, Quantidade: 8, Categoria: Roupa},
	}

	// Ordena as chaves para saida consistente
	chaves := make([]string, 0, len(produtos))
	for k := range produtos {
		chaves = append(chaves, k)
	}
	sort.Strings(chaves)

	fmt.Println("\n=== Inventário ===")
	for _, k := range chaves {
		item := produtos[k]
		fmt.Printf("%-12s | %-10v | R$ %8.2f | Qtd: %d\n", k, item.Categoria, item.Preco, item.Quantidade)
	}

	var total float64
	for _, item := range produtos {
		total += item.Preco * float64(item.Quantidade)
	}
	fmt.Printf("\nValor total em estoque: R$ %.2f\n", total)

	fmt.Println("\n=== Estoque Baixo ===")
	for _, k := range chaves {
		item := produtos[k]
		if item.Quantidade < 3 {
			fmt.Printf("%s (%d unidade(s))\n", k, item.Quantidade)
		}
	}

}
