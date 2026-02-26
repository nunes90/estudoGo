# Capítulo 8 — Generics

> **Revisão rápida:** Generics (introduzidos no Go 1.18) permitem escrever código que funciona com múltiplos tipos sem repetição, mantendo a segurança de tipos em tempo de compilação. É uma ferramenta poderosa, mas deve ser usada com moderação — Go valoriza a simplicidade.

---

## 1. Type Parameters (Parâmetros de Tipo)

```go
package main

import "fmt"

// Sem generics: precisa repetir para cada tipo
func somarInts(a, b int) int      { return a + b }
func somarFloat(a, b float64) float64 { return a + b }

// Com generics: um único código para todos os tipos numéricos
func somar[T int | float64 | int32 | int64](a, b T) T {
    return a + b
}

// Type parameter em structs
type Par[T, U any] struct {
    Primeiro T
    Segundo  U
}

func (p Par[T, U]) String() string {
    return fmt.Sprintf("(%v, %v)", p.Primeiro, p.Segundo)
}

// Stack genérico
type Pilha[T any] struct {
    itens []T
}

func (p *Pilha[T]) Push(item T) {
    p.itens = append(p.itens, item)
}

func (p *Pilha[T]) Pop() (T, bool) {
    var zero T
    if len(p.itens) == 0 {
        return zero, false
    }
    ultimo := p.itens[len(p.itens)-1]
    p.itens = p.itens[:len(p.itens)-1]
    return ultimo, true
}

func (p *Pilha[T]) Tamanho() int {
    return len(p.itens)
}

func main() {
    fmt.Println(somar(3, 5))           // int
    fmt.Println(somar(3.14, 2.71))     // float64

    coordenada := Par[float64, float64]{-23.55, -46.63}
    registro := Par[string, int]{"Lucas", 30}
    fmt.Println(coordenada)
    fmt.Println(registro)

    // Pilha de strings
    historico := Pilha[string]{}
    historico.Push("pagina-inicial")
    historico.Push("produtos")
    historico.Push("detalhes-produto-42")

    fmt.Println("\nHistórico de navegação (LIFO):")
    for historico.Tamanho() > 0 {
        if pagina, ok := historico.Pop(); ok {
            fmt.Println(" ←", pagina)
        }
    }

    // Pilha de inteiros
    operacoes := Pilha[int]{}
    operacoes.Push(10)
    operacoes.Push(20)
    operacoes.Push(30)
    fmt.Printf("\nPilha de operações tem %d itens\n", operacoes.Tamanho())
}
```

---

## 2. Type Constraints (Restrições de Tipo)

```go
package main

import (
    "fmt"
    "golang.org/x/exp/constraints" // ou defina suas próprias
)

// Constraint customizada
type Numerico interface {
    ~int | ~int8 | ~int16 | ~int32 | ~int64 |
        ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
        ~float32 | ~float64
}

// O ~ significa "tipos cujo tipo subjacente é T"
type Celsius float64
type Fahrenheit float64
// Graças ao ~, Celsius e Fahrenheit também satisfazem Numerico!

func maximo[T Numerico](a, b T) T {
    if a > b {
        return a
    }
    return b
}

func soma[T Numerico](nums []T) T {
    var total T
    for _, n := range nums {
        total += n
    }
    return total
}

// Constraint com método
type Formatavel interface {
    Formatar() string
}

type Produto struct {
    Nome  string
    Preco float64
}

func (p Produto) Formatar() string {
    return fmt.Sprintf("%s: R$%.2f", p.Nome, p.Preco)
}

type Evento struct {
    Titulo string
    Data   string
}

func (e Evento) Formatar() string {
    return fmt.Sprintf("[%s] %s", e.Data, e.Titulo)
}

func listar[T Formatavel](itens []T) {
    for i, item := range itens {
        fmt.Printf("%d. %s\n", i+1, item.Formatar())
    }
}

// Constraint combinando interface e tipos
type OrdenaveisBasicos interface {
    ~int | ~float64 | ~string
}

func ordenarTres[T OrdenaveisBasicos](a, b, c T) (T, T, T) {
    if a > b {
        a, b = b, a
    }
    if b > c {
        b, c = c, b
    }
    if a > b {
        a, b = b, a
    }
    return a, b, c
}

func main() {
    fmt.Println(maximo(10, 20))
    fmt.Println(maximo(3.14, 2.71))
    fmt.Println(maximo(Celsius(37.5), Celsius(36.0)))

    ints := []int{5, 3, 8, 1, 9, 2}
    floats := []float64{1.1, 2.2, 3.3}
    fmt.Println("Soma ints:", soma(ints))
    fmt.Println("Soma floats:", soma(floats))

    produtos := []Produto{
        {"Notebook", 4500.00},
        {"Mouse", 89.90},
        {"Teclado", 199.00},
    }
    fmt.Println("\nProdutos:")
    listar(produtos)

    eventos := []Evento{
        {"GopherCon", "2024-11-15"},
        {"TechDay SP", "2024-12-01"},
    }
    fmt.Println("\nEventos:")
    listar(eventos)

    a, b, c := ordenarTres(3, 1, 2)
    fmt.Println("\nOrdenados:", a, b, c)

    x, y, z := ordenarTres("banana", "abacate", "caju")
    fmt.Println("Ordenados:", x, y, z)
}
```

---

## 3. Type Inference (Inferência de Tipos)

```go
package main

import "fmt"

func converter[De, Para any](valor De, fn func(De) Para) Para {
    return fn(valor)
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

func mapear[T, U any](slice []T, fn func(T) U) []U {
    resultado := make([]U, len(slice))
    for i, v := range slice {
        resultado[i] = fn(v)
    }
    return resultado
}

func main() {
    // Go infere os tipos automaticamente na maioria dos casos
    numeros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

    // Inferência: T=int
    pares := filtrar(numeros, func(n int) bool { return n%2 == 0 })
    fmt.Println("Pares:", pares)

    // Inferência: T=int, U=string
    textos := mapear(pares, func(n int) string {
        return fmt.Sprintf("num_%d", n)
    })
    fmt.Println("Textos:", textos)

    // Casos onde você PRECISA ser explícito
    // Quando o compilador não consegue inferir:
    resultado := converter[string, int]("12345", func(s string) int {
        total := 0
        for _, ch := range s {
            total += int(ch - '0')
        }
        return total
    })
    fmt.Println("Soma dos dígitos de '12345':", resultado) // 15

    // Inferência funciona em cadeia de chamadas simples
    palavras := []string{"Go", "é", "incrível", "mas", "simples"}
    longas := filtrar(palavras, func(s string) bool { return len([]rune(s)) > 2 })
    maiusculas := mapear(longas, func(s string) string {
        runes := []rune(s)
        if runes[0] >= 'a' && runes[0] <= 'z' {
            runes[0] -= 32
        }
        return string(runes)
    })
    fmt.Println("Palavras longas capitalizadas:", maiusculas)
}
```

---

## 4. Quando usar Generics vs Interfaces

```go
package main

import "fmt"

// ============================================================
// INTERFACES: quando o comportamento importa mais que o tipo
// ============================================================

type Serializable interface {
    Serializar() string
}

type JSON struct{ Dados string }
type XML struct{ Dados string }

func (j JSON) Serializar() string { return `{"dados":"` + j.Dados + `"}` }
func (x XML) Serializar() string  { return `<dados>` + x.Dados + `</dados>` }

// Use interface quando diferentes tipos têm comportamentos DIFERENTES
func publicar(s Serializable) {
    fmt.Println("Publicando:", s.Serializar())
}

// ============================================================
// GENERICS: quando a ESTRUTURA/ALGORITMO é o mesmo para vários tipos
// ============================================================

// Isso NÃO deveria ser interface — o comportamento é idêntico para todos os tipos
func primeiro[T any](slice []T) (T, bool) {
    var zero T
    if len(slice) == 0 {
        return zero, false
    }
    return slice[0], true
}

func ultimo[T any](slice []T) (T, bool) {
    var zero T
    if len(slice) == 0 {
        return zero, false
    }
    return slice[len(slice)-1], true
}

func contem[T comparable](slice []T, item T) bool {
    for _, v := range slice {
        if v == item {
            return true
        }
    }
    return false
}

// Mapa genérico de chave→valor com operações seguras
type Mapa[K comparable, V any] struct {
    dados map[K]V
}

func NovoMapa[K comparable, V any]() *Mapa[K, V] {
    return &Mapa[K, V]{dados: make(map[K]V)}
}

func (m *Mapa[K, V]) Set(chave K, valor V) {
    m.dados[chave] = valor
}

func (m *Mapa[K, V]) Get(chave K) (V, bool) {
    v, ok := m.dados[chave]
    return v, ok
}

func (m *Mapa[K, V]) Chaves() []K {
    chaves := make([]K, 0, len(m.dados))
    for k := range m.dados {
        chaves = append(chaves, k)
    }
    return chaves
}

func main() {
    // Interface — comportamentos diferentes por tipo
    publicar(JSON{Dados: "hello"})
    publicar(XML{Dados: "hello"})

    // Generics — mesma estrutura, tipos diferentes
    ints := []int{10, 20, 30, 40, 50}
    strs := []string{"Go", "Rust", "Zig"}

    if v, ok := primeiro(ints); ok {
        fmt.Println("Primeiro int:", v)
    }
    if v, ok := ultimo(strs); ok {
        fmt.Println("Última string:", v)
    }

    fmt.Println("Contém 30?", contem(ints, 30))
    fmt.Println("Contém 'Go'?", contem(strs, "Go"))
    fmt.Println("Contém 'Java'?", contem(strs, "Java"))

    // Mapa genérico
    cache := NovoMapa[string, int]()
    cache.Set("hits", 150)
    cache.Set("misses", 30)
    cache.Set("erros", 5)

    if v, ok := cache.Get("hits"); ok {
        fmt.Println("Cache hits:", v)
    }
    fmt.Println("Chaves:", cache.Chaves())
}
```

---

## 5. Boas Práticas com Generics

```go
package main

import "fmt"

// ✅ BOA PRÁTICA 1: Funções em vez de métodos em tipos genéricos
// Mais flexível e reutilizável

type Resultado[T any] struct {
    Valor T
    Erro  error
}

// Função genérica — mais fácil de reutilizar
func Ok[T any](valor T) Resultado[T] {
    return Resultado[T]{Valor: valor}
}

func Err[T any](err error) Resultado[T] {
    return Resultado[T]{Erro: err}
}

func (r Resultado[T]) OK() bool {
    return r.Erro == nil
}

// ✅ BOA PRÁTICA 2: Constraints gerais em vez de constraints que exigem métodos

// ❌ Menos flexível — exige método específico
type Printable interface {
    Print()
}

// ✅ Mais flexível — passa a função como argumento
func imprimirTodos[T any](itens []T, formatar func(T) string) {
    for i, item := range itens {
        fmt.Printf("%d: %s\n", i+1, formatar(item))
    }
}

// ✅ BOA PRÁTICA 3: Comece com funções, converta para métodos se necessário
func zip[T, U any](ts []T, us []U) []struct{ T; U } {
    tamanho := len(ts)
    if len(us) < tamanho {
        tamanho = len(us)
    }
    resultado := make([]struct{ T; U }, tamanho)
    for i := range resultado {
        resultado[i].T = ts[i]
        resultado[i].U = us[i]
    }
    return resultado
}

func main() {
    // Usando Resultado genérico (estilo Result type de Rust)
    r1 := Ok[int](42)
    r2 := Err[string](fmt.Errorf("algo deu errado"))

    if r1.OK() {
        fmt.Println("Valor:", r1.Valor)
    }
    if !r2.OK() {
        fmt.Println("Erro:", r2.Erro)
    }

    // Função flexível com formatador como parâmetro
    numeros := []int{10, 20, 30}
    imprimirTodos(numeros, func(n int) string {
        return fmt.Sprintf("número %d", n)
    })

    nomes := []string{"Alice", "Bruno", "Carlos"}
    imprimirTodos(nomes, func(s string) string {
        return "👤 " + s
    })

    // Zip: combina dois slices em pares
    ids := []int{1, 2, 3}
    usuarios := []string{"Alice", "Bruno", "Carlos"}
    pares := zip(ids, usuarios)
    for _, par := range pares {
        fmt.Printf("  ID %d → %s\n", par.T, par.U)
    }
}
```

---

## 🔁 Revisão Rápida — O que lembrar

| Conceito | Ponto chave |
|---|---|
| `[T any]` | Parâmetro de tipo — funciona com qualquer tipo |
| `[T comparable]` | Permite uso de `==` e `!=` |
| `~int` | O tipo subjacente é int (inclui tipos customizados) |
| Inferência de tipos | Go geralmente infere — especifique só quando necessário |
| Generic vs Interface | Interface = **comportamento diferente**. Generic = **estrutura igual, tipos diferentes** |

**Quando usar generics?** Quando você se pega copiando a mesma função para `int`, `string`, `float64`, etc. — e a lógica é **exatamente a mesma**. Se cada tipo precisa de lógica diferente, use interfaces.

**Quando NÃO usar?** Quando interfaces já resolvem bem. Complexidade desnecessária prejudica a legibilidade — Go é uma linguagem que valoriza código simples e claro.

---

## 🗺️ Visão Geral dos 8 Capítulos

```
Cap 1: Fundamentos    → Variáveis, tipos, ponteiros, constantes
Cap 2: Fluxo         → if, switch, for (único loop), break/continue
Cap 3: Tipos         → bool, int, float, byte, rune, string, nil
Cap 4: Composição    → arrays, slices, structs, type assertions
Cap 5: Funções       → múltiplos retornos, closures, defer, variadic
Cap 6: Erros         → error como valor, wrapping, panic/recover
Cap 7: Interfaces    → duck typing, polimorfismo, any
Cap 8: Generics      → type params, constraints, inference, boas práticas
```
