# Capítulo 5 — Funções

> **Revisão rápida:** Em Go, funções são cidadãs de primeira classe — podem ser armazenadas em variáveis, passadas como parâmetros e retornadas de outras funções. `defer` e closures são ferramentas extremamente poderosas que aparecem em código real o tempo todo.

---

## 1. Funções Básicas

```go
package main

import (
    "fmt"
    "math"
)

// Múltiplos retornos (idiomático em Go)
func dividir(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("divisão por zero")
    }
    return a / b, nil
}

// Parâmetros do mesmo tipo podem ser agrupados
func somarTres(a, b, c int) int {
    return a + b + c
}

// Retornando múltiplos valores do mesmo tipo
func minMax(nums []float64) (float64, float64) {
    if len(nums) == 0 {
        return 0, 0
    }
    min, max := nums[0], nums[0]
    for _, n := range nums[1:] {
        if n < min {
            min = n
        }
        if n > max {
            max = n
        }
    }
    return min, max
}

func main() {
    resultado, err := dividir(10, 3)
    if err != nil {
        fmt.Println("Erro:", err)
    } else {
        fmt.Printf("10 / 3 = %.4f\n", resultado)
    }

    fmt.Println(somarTres(1, 2, 3))

    temperaturas := []float64{22.5, 35.1, 18.0, 29.8, 31.5}
    min, max := minMax(temperaturas)
    fmt.Printf("Mín: %.1f  Máx: %.1f\n", min, max)

    // Descartando retornos com _
    distancia := math.Sqrt(math.Pow(3, 2) + math.Pow(4, 2))
    fmt.Printf("Hipotenusa: %.0f\n", distancia)
}
```

---

## 2. Parâmetros

```go
package main

import "fmt"

// Parâmetro por valor — cópia
func incrementar(n int) int {
    n++
    return n
}

// Parâmetro por ponteiro — referência
func incrementarPtr(n *int) {
    *n++
}

// Parâmetro slice — compartilha memória
func dobrarElementos(nums []int) {
    for i := range nums {
        nums[i] *= 2
    }
}

func main() {
    x := 10
    y := incrementar(x)
    fmt.Println(x, y) // 10, 11 — x não mudou

    incrementarPtr(&x)
    fmt.Println(x) // 11 — x mudou

    nums := []int{1, 2, 3, 4, 5}
    dobrarElementos(nums)
    fmt.Println(nums) // [2 4 6 8 10] — slice compartilha memória
}
```

---

## 3. Naked Returns (Retornos Nomeados)

```go
package main

import "fmt"

// Retornos nomeados — o return sem argumentos retorna os valores nomeados
func calcularCirculo(raio float64) (area, perimetro float64) {
    const pi = 3.14159265
    area = pi * raio * raio
    perimetro = 2 * pi * raio
    return // naked return — retorna area e perimetro
}

// Útil em funções curtas; evite em funções longas
func parseNome(nomeCompleto string) (primeiro, ultimo string) {
    for i, ch := range nomeCompleto {
        if ch == ' ' {
            primeiro = nomeCompleto[:i]
            ultimo = nomeCompleto[i+1:]
            return
        }
    }
    primeiro = nomeCompleto
    return
}

func main() {
    a, p := calcularCirculo(5)
    fmt.Printf("Área: %.2f  Perímetro: %.2f\n", a, p)

    nome, sobre := parseNome("Lucas Nunes")
    fmt.Println(nome, "|", sobre)

    nome2, sobre2 := parseNome("Gopher")
    fmt.Println(nome2, "|", sobre2) // Gopher | (vazio)
}
```

> **💡 Insight:** Naked returns reduzem repetição mas prejudicam a legibilidade em funções longas. Use-os apenas em funções curtas e simples. O código deve ser óbvio sobre o que está sendo retornado.

---

## 4. Variadic Functions

```go
package main

import "fmt"

// ... indica quantidade variável de argumentos
func somar(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}

func log(nivel string, args ...interface{}) {
    fmt.Printf("[%s] ", nivel)
    fmt.Println(args...)
}

func concatenar(sep string, partes ...string) string {
    resultado := ""
    for i, p := range partes {
        if i > 0 {
            resultado += sep
        }
        resultado += p
    }
    return resultado
}

func main() {
    fmt.Println(somar(1, 2, 3))           // 6
    fmt.Println(somar(10, 20, 30, 40))    // 100

    // Expandindo slice com ...
    numeros := []int{5, 10, 15, 20}
    fmt.Println(somar(numeros...)) // 50

    log("INFO", "Servidor iniciado na porta", 8080)
    log("ERRO", "Conexão recusada:", "timeout")

    fmt.Println(concatenar(", ", "Go", "Rust", "Zig")) // Go, Rust, Zig
    fmt.Println(concatenar("-", "2024", "01", "15"))   // 2024-01-15
}
```

---

## 5. Anonymous Functions

```go
package main

import (
    "fmt"
    "sort"
)

func main() {
    // Função anônima executada imediatamente (IIFE)
    resultado := func(a, b int) int {
        return a * b
    }(6, 7)
    fmt.Println(resultado) // 42

    // Atribuída a variável
    quadrado := func(n float64) float64 {
        return n * n
    }

    fmt.Printf("5² = %.0f\n", quadrado(5))
    fmt.Printf("12² = %.0f\n", quadrado(12))

    // Passada como argumento
    numeros := []int{5, 2, 8, 1, 9, 3}
    sort.Slice(numeros, func(i, j int) bool {
        return numeros[i] > numeros[j] // decrescente
    })
    fmt.Println(numeros) // [9 8 5 3 2 1]

    // Retornada de outra função
    pessoas := []struct{ Nome string; Idade int }{
        {"Alice", 30},
        {"Bruno", 25},
        {"Carlos", 35},
    }
    sort.Slice(pessoas, func(i, j int) bool {
        return pessoas[i].Idade < pessoas[j].Idade
    })
    for _, p := range pessoas {
        fmt.Printf("%s (%d)\n", p.Nome, p.Idade)
    }
}
```

---

## 6. Closures

```go
package main

import "fmt"

// Closure captura variáveis do escopo externo
func criarContador(inicio int) func() int {
    count := inicio
    return func() int {
        count++
        return count
    }
}

func criarAcumulador() func(float64) float64 {
    total := 0.0
    return func(valor float64) float64 {
        total += valor
        return total
    }
}

// Closure para memoização (cache de resultados)
func memoizarFibonacci() func(int) int {
    cache := map[int]int{}
    var fib func(n int) int
    fib = func(n int) int {
        if n <= 1 {
            return n
        }
        if v, ok := cache[n]; ok {
            return v
        }
        cache[n] = fib(n-1) + fib(n-2)
        return cache[n]
    }
    return fib
}

func main() {
    // Contadores independentes (cada um tem seu próprio estado)
    contador1 := criarContador(0)
    contador2 := criarContador(100)

    fmt.Println(contador1()) // 1
    fmt.Println(contador1()) // 2
    fmt.Println(contador2()) // 101
    fmt.Println(contador1()) // 3

    // Acumulador
    caixa := criarAcumulador()
    fmt.Printf("Total: R$ %.2f\n", caixa(50.00))
    fmt.Printf("Total: R$ %.2f\n", caixa(30.50))
    fmt.Printf("Total: R$ %.2f\n", caixa(19.90))

    // Fibonacci memoizado
    fib := memoizarFibonacci()
    for i := 0; i <= 10; i++ {
        fmt.Printf("fib(%d) = %d\n", i, fib(i))
    }
}
```

> **💡 Insight:** Closures são a base de muitos padrões em Go: middleware, callbacks, memoização, e o próprio `defer`. A variável capturada é **compartilhada** — não copiada — então mudanças dentro da closure afetam o valor original.

---

## 7. Function Types

```go
package main

import (
    "fmt"
    "strings"
)

// Definindo um tipo de função
type Transformador func(string) string
type Predicado func(string) bool

func aplicarTransformacoes(texto string, transforms ...Transformador) string {
    for _, t := range transforms {
        texto = t(texto)
    }
    return texto
}

func filtrar(palavras []string, pred Predicado) []string {
    resultado := []string{}
    for _, p := range palavras {
        if pred(p) {
            resultado = append(resultado, p)
        }
    }
    return resultado
}

func main() {
    // Funções como valores
    maiusculo := Transformador(strings.ToUpper)
    semEspacos := Transformador(strings.TrimSpace)
    exclamar  := Transformador(func(s string) string { return s + "!" })

    resultado := aplicarTransformacoes("  olá, mundo  ", semEspacos, maiusculo, exclamar)
    fmt.Println(resultado) // OLÁ, MUNDO!

    linguagens := []string{"Go", "Rust", "Python", "Zig", "Java", "C"}
    curtas := filtrar(linguagens, func(s string) bool {
        return len(s) <= 3
    })
    fmt.Println(curtas) // [Go Zig C]
}
```

---

## 8. Defer

```go
package main

import "fmt"

func operacaoBancaria(valor float64) error {
    fmt.Println("Iniciando transação...")
    defer fmt.Println("Transação finalizada (sempre executa)")

    if valor <= 0 {
        return fmt.Errorf("valor inválido: %.2f", valor)
    }

    fmt.Printf("Transferindo R$ %.2f...\n", valor)
    return nil
}

func ordemDeExecucao() {
    // Defers executam em ordem LIFO (último a entrar, primeiro a sair)
    defer fmt.Println("Terceiro defer — executa primeiro")
    defer fmt.Println("Segundo defer — executa segundo")
    defer fmt.Println("Primeiro defer — executa terceiro")
    fmt.Println("Corpo da função")
}

func main() {
    err := operacaoBancaria(500.00)
    if err != nil {
        fmt.Println("Erro:", err)
    }

    fmt.Println()
    err = operacaoBancaria(-100)
    if err != nil {
        fmt.Println("Erro:", err)
    }

    fmt.Println()
    ordemDeExecucao()
}
```

---

## 9. Separando Código Similar (Higher-Order Functions)

```go
package main

import "fmt"

type Numero interface {
    ~int | ~float64
}

// Ao invés de repetir lógica de iteração, extrai em funções reutilizáveis
func mapear[T, U any](slice []T, fn func(T) U) []U {
    resultado := make([]U, len(slice))
    for i, v := range slice {
        resultado[i] = fn(v)
    }
    return resultado
}

func filtrar[T any](slice []T, pred func(T) bool) []T {
    var resultado []T
    for _, v := range slice {
        if pred(v) {
            resultado = append(resultado, v)
        }
    }
    return resultado
}

func reduzir[T, U any](slice []T, inicial U, fn func(U, T) U) U {
    acumulador := inicial
    for _, v := range slice {
        acumulador = fn(acumulador, v)
    }
    return acumulador
}

func main() {
    numeros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

    pares := filtrar(numeros, func(n int) bool { return n%2 == 0 })
    fmt.Println("Pares:", pares) // [2 4 6 8 10]

    quadrados := mapear(pares, func(n int) int { return n * n })
    fmt.Println("Quadrados:", quadrados) // [4 16 36 64 100]

    soma := reduzir(quadrados, 0, func(acc, n int) int { return acc + n })
    fmt.Println("Soma dos quadrados dos pares:", soma) // 220

    // Compondo
    palavras := []string{"go", "rust", "zig", "python", "c", "java"}
    curtas := filtrar(palavras, func(s string) bool { return len(s) <= 3 })
    emMaiusculo := mapear(curtas, func(s string) string {
        return "[" + s + "]"
    })
    fmt.Println(emMaiusculo) // [[go] [zig] [c]]
}
```

---

## 🔁 Revisão Rápida — O que lembrar

| Conceito | Ponto chave |
|---|---|
| Múltiplos retornos | `func f() (int, error)` — idiomático em Go |
| Naked returns | Só em funções curtas e óbvias |
| Variadic `...T` | Expanda slices com `slice...` na chamada |
| Closures | Capturam variáveis por referência, não por cópia |
| `defer` | Executa ao sair da função — LIFO — ideal para limpeza |
| Function types | Funções são valores — passe-as como parâmetros |
