// ex03 - Sistema Bancário Simples
package main

import (
	"errors"
	"fmt"
)

type ContaBancaria interface {
	Depositar(valor float64) error
	Sacar(valor float64) error
	Saldo() float64
	Extrato() []string
}

// ----------------------------------------------------------------------------
// permite saldo negativo ate um limite de -500.00
type ContaCorrente struct {
	nome    string
	saldo   float64
	extrato []string // campo próprio
}

func (c *ContaCorrente) Depositar(valor float64) error {
	if valor <= 0 {
		return errors.New("valor invalido")
	}
	c.saldo += valor
	c.extrato = append(c.extrato, fmt.Sprintf("Deposito: +R$ %.2f", valor))

	return nil
}

func (c *ContaCorrente) Sacar(valor float64) error {
	if valor < 0 {
		return errors.New("valor invalido")
	}
	if (c.saldo - valor) < -500 {
		return errors.New("saldo insuficiente")
	}
	c.saldo -= valor
	c.extrato = append(c.extrato, fmt.Sprintf("Saque: -R$ %.2f", valor))
	return nil
}

func (c *ContaCorrente) Saldo() float64 {
	return c.saldo
}

func (c *ContaCorrente) Extrato() []string {
	fmt.Printf("\n=== Extrato -- Conta Corrente (%s) ===\n", c.nome)
	return c.extrato
}

// ----------------------------------------------------------------------------
// nao permite saldo negativo (erro se tentar sacar mais do que tem)
type ContaPoupanca struct {
	nome    string
	saldo   float64
	extrato []string
}

func (c *ContaPoupanca) Depositar(valor float64) error {
	if valor < 0 {
		return errors.New("valor invalido")
	}
	c.saldo += valor
	c.extrato = append(c.extrato, fmt.Sprintf("Deposito: +R$ %.2f", valor))
	return nil
}

func (c *ContaPoupanca) Sacar(valor float64) error {
	if valor < 0 {
		return errors.New("valor invalido")
	}
	if valor > c.saldo {
		return errors.New("saldo insuficiente")
	}
	c.saldo -= valor
	c.extrato = append(c.extrato, fmt.Sprintf("Saque   : -R$ %.2f", valor))
	return nil
}

func (c *ContaPoupanca) Saldo() float64 {
	return c.saldo
}

func (c *ContaPoupanca) Extrato() []string {
	fmt.Printf("\n=== Extrato -- Conta Poupança (%s) ===\n", c.nome)

	return c.extrato

}

// ----------------------------------------------------------------------------
func RealizarTransferencia(origem, destino ContaBancaria, valor float64) error {
	err := origem.Sacar(valor)
	if err != nil {
		return err
	}
	err = destino.Depositar(valor)
	if err != nil {
		origem.Depositar(valor) // desfaz o saque
		return err
	}
	return nil
}

func main() {

	cc := &ContaCorrente{
		nome:    "Lucas",
		saldo:   0,
		extrato: []string{},
	}
	cp := &ContaPoupanca{
		nome:    "Maria",
		saldo:   0,
		extrato: []string{},
	}

	cc.Depositar(1000)
	cc.Sacar(200)

	for _, op := range cc.Extrato() {
		fmt.Println(op)
	}
	fmt.Printf("Saldo atual R$ %.2f\n", cc.Saldo())

	fmt.Println("\n=== Transferência ===")
	err := RealizarTransferencia(cc, cp, 300)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Printf("Transferência de R$ 300.00 realizada com sucesso")
	}

	println("")
	for _, op := range cp.Extrato() {
		fmt.Println(op)
	}

	fmt.Printf("Saldo atual: R$ %.2f\n", cp.Saldo())
}
