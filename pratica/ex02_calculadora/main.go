// ex2 - Calculadora com Tratamento de Erros
package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func Somar(a, b float64) (float64, error) {
	return a + b, nil
}

func Subtrair(a, b float64) (float64, error) {
	return a - b, nil
}

func Multiplicar(a, b float64) (float64, error) {
	return a * b, nil
}

func Dividir(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("divisão por zero não é permitida")
	}
	return a / b, nil
}

func Calcular(a, b float64, op string) (float64, error) {
	var resultado float64
	var err error

	switch op {
	case "+":
		resultado, err = Somar(a, b)
	case "-":
		resultado, err = Subtrair(a, b)
	case "*":
		resultado, err = Multiplicar(a, b)
	case "/":
		resultado, err = Dividir(a, b)
	default:
		err = fmt.Errorf("operação '%s' não suportada", op)
	}

	return resultado, err
}

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Uso: calculadora <numero1> <operacao> <numero2>")
		os.Exit(1)
	}

	// string em float
	num1, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("Número inválido:", os.Args[1])
		os.Exit(1)
	}

	num2, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		fmt.Println("Número inválido:", os.Args[3])
		os.Exit(1)
	}

	resultado, err := Calcular(num1, num2, os.Args[2])
	if err != nil {
		fmt.Println("Erro:", err)
		os.Exit(1)
	}

	fmt.Printf("Resultado: %.2f\n", resultado)
}
