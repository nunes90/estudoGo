// ex6 - Calculadora de Formas Geométricas(Desafio)
package main

import (
	"fmt"
	"math"
)

type Forma interface {
	Area() float64
	Perimetro() float64
	Descricao() string
}

//-----------------------------------------------------------------------------

type Circulo struct {
	Raio float64
}

func (c Circulo) Area() float64 {
	return math.Pi * c.Raio * c.Raio
}

func (c Circulo) Perimetro() float64 {
	return 2 * math.Pi * c.Raio
}

func (c Circulo) Descricao() string {
	return fmt.Sprintf("Círculo com raio %.2f", c.Raio)
}

//-----------------------------------------------------------------------------

type Retangulo struct {
	Largura float64
	Altura  float64
}

func (r Retangulo) Area() float64 {
	return r.Largura * r.Altura
}

func (r Retangulo) Perimetro() float64 {
	return 2 * (r.Largura + r.Altura)
}

func (r Retangulo) Descricao() string {
	return fmt.Sprintf("Retângulo com largura %.2f e altura %.2f", r.Largura, r.Altura)
}

//-----------------------------------------------------------------------------

type Triangulo struct {
	Base   float64
	Altura float64
	LadoA  float64
	LadoB  float64
	LadoC  float64
}

func (t Triangulo) Area() float64 {
	return 0.5 * t.Base * t.Altura
}

func (t Triangulo) Perimetro() float64 {
	return t.LadoA + t.LadoB + t.LadoC
}

func (t Triangulo) Descricao() string {
	return fmt.Sprintf("Triângulo com base %.2f, altura %.2f e lados %.2f, %.2f, %.2f", t.Base, t.Altura, t.LadoA, t.LadoB, t.LadoC)
}

//-----------------------------------------------------------------------------

type Quadrado struct {
	Lado float64
}

func (q Quadrado) Area() float64 {
	return q.Lado * q.Lado
}

func (q Quadrado) Perimetro() float64 {
	return 4 * q.Lado
}

func (q Quadrado) Descricao() string {
	return fmt.Sprintf("Quadrado com lado %.2f", q.Lado)
}

//-----------------------------------------------------------------------------

func ImprimirInfo(f Forma) {
	fmt.Println(f.Descricao())
	fmt.Printf("Área: %.2f\n", f.Area())
	fmt.Printf("Perímetro: %.2f\n", f.Perimetro())
	fmt.Println()
}

func MaiorArea(formas []Forma) Forma {
	var fMaiorArea Forma = formas[0]

	for _, forma := range formas {
		if forma.Area() > fMaiorArea.Area() {
			fMaiorArea = forma
		}
	}
	return fMaiorArea
}

func TotalArea(formas []Forma) float64 {
	var totalArea float64

	for _, forma := range formas {
		totalArea += forma.Area()
	}
	return totalArea
}

//-----------------------------------------------------------------------------

func main() {

	f1 := Circulo{Raio: 5}
	f2 := Retangulo{Largura: 2, Altura: 3}
	f3 := Triangulo{Base: 3, Altura: 4, LadoA: 3, LadoB: 4, LadoC: 5}
	f4 := Quadrado{Lado: 4}

	formas := []Forma{f1, f2, f3, f4}

	maior := MaiorArea(formas)
	fmt.Printf("\nForma com maior área: %s\n", maior.Descricao())

	total := TotalArea(formas)
	fmt.Printf("\nÁrea total das formas: %.2f\n", total)

	fmt.Println("\nInformações das formas:\n")
	for _, forma := range formas {
		ImprimirInfo(forma)
	}
}
