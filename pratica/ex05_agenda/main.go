// ex 5 - Agenda de Contatos com Validação
package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

type ErrValidacao struct {
	Campo    string
	Mensagem string
}

func (e ErrValidacao) Error() string {
	return fmt.Sprintf("campo %s inválido: %s", e.Campo, e.Mensagem)
}

type Contato struct {
	Nome     string
	Email    string
	Telefone string
}

func ValidarContato(c Contato) []error {
	var err []error

	if c.Nome == "" {
		err = append(err, ErrValidacao{Campo: "Nome", Mensagem: "não pode ser vazio"})
	}

	if !strings.Contains(c.Email, "@") || !strings.Contains(c.Email, ".") {
		err = append(err, ErrValidacao{Campo: "Email", Mensagem: "deve conter @ e ."})
	}

	if len(c.Telefone) < 8 || len(c.Telefone) > 15 {
		err = append(err, ErrValidacao{Campo: "Telefone", Mensagem: "deve ter entre 8 e 15 caracteres (apenas digitos)"})
	}
	for _, r := range c.Telefone {
		if !unicode.IsDigit(r) {
			err = append(err, ErrValidacao{Campo: "Telefone", Mensagem: "deve ter entre 8 e 15 caracteres (apenas digitos)"})
			break
		}
	}

	return err
}

type Agenda map[string]Contato

func (a *Agenda) Adicionar(c Contato) error {
	if err := ValidarContato(c); err != nil {
		return fmt.Errorf("Erro ao adicionar: %v", err[0])
	}
	(*a)[c.Nome] = c
	fmt.Printf("Contato '%s' adicionado com sucesso!\n", c.Nome)
	return nil
}

func (a *Agenda) Buscar(nome string) (Contato, error) {
	for _, contato := range *a {
		if contato.Nome == nome {
			return contato, nil
		}
	}
	return Contato{}, fmt.Errorf("Contato não encontrado")
}

func (a *Agenda) Listar() []Contato {
	// 1. extrai os contatos do map para um slice
	lista := make([]Contato, 0, len(*a))
	for _, contato := range *a {
		lista = append(lista, contato)
	}

	// 2. ordena o slice por Nome (< = ordem alfabética crescente)
	sort.Slice(lista, func(i, j int) bool {
		return lista[i].Nome < lista[j].Nome
	})

	return lista
}

func main() {

	agenda := make(Agenda)

	// contato invalido 1 - email invalido
	c1 := Contato{
		Nome:     "João",
		Email:    "joao@",
		Telefone: "12345678",
	}
	if err := agenda.Adicionar(c1); err != nil {
		fmt.Println(err)
	}

	// contato invalido 2 - Telefone invalido
	c2 := Contato{
		Nome:     "Maria",
		Email:    "maria@email.com",
		Telefone: "12345678910234234",
	}
	if err := agenda.Adicionar(c2); err != nil {
		fmt.Println(err)
	}

	// contato valido
	c3 := Contato{
		Nome:     "Lucas Nunes",
		Email:    "pedro@email.com",
		Telefone: "12345678",
	}
	if err := agenda.Adicionar(c3); err != nil {
		fmt.Println(err)
	}

	fmt.Println("\n=== Lista de Contatos ===")
	for _, c := range agenda.Listar() {
		fmt.Printf("  %s | %s | %s\n", c.Nome, c.Email, c.Telefone)
	}

	fmt.Println("\n=== Buscar ===")
	if c, err := agenda.Buscar("Lucas Nunes"); err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Printf("Encontrado: %s (%s)\n", c.Nome, c.Email)
	}

	// testa busca por contato inexistente
	if _, err := agenda.Buscar("Zé Ninguém"); err != nil {
		fmt.Println("Erro:", err)
	}
}
