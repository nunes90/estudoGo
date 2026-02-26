# Capítulo 7 — Interfaces

> **Revisão rápida:** Interfaces em Go são o mecanismo central de abstração. Diferente de Java ou C#, você não declara explicitamente que um tipo implementa uma interface — se ele tem os métodos certos, ele implementa. Isso é chamado de duck typing estrutural e é extremamente poderoso.

---

## 1. Interface

```go
package main

import (
    "fmt"
    "math"
)

// Interface define comportamento, não estrutura
type Forma interface {
    Area() float64
    Perimetro() float64
}

type Circulo struct {
    Raio float64
}

type Retangulo struct {
    Largura, Altura float64
}

type Triangulo struct {
    A, B, C float64 // lados
}

// Circulo implementa Forma (implicitamente)
func (c Circulo) Area() float64 {
    return math.Pi * c.Raio * c.Raio
}

func (c Circulo) Perimetro() float64 {
    return 2 * math.Pi * c.Raio
}

// Retangulo implementa Forma
func (r Retangulo) Area() float64 {
    return r.Largura * r.Altura
}

func (r Retangulo) Perimetro() float64 {
    return 2 * (r.Largura + r.Altura)
}

// Triangulo implementa Forma (Heron para área)
func (t Triangulo) Area() float64 {
    s := (t.A + t.B + t.C) / 2
    return math.Sqrt(s * (s - t.A) * (s - t.B) * (s - t.C))
}

func (t Triangulo) Perimetro() float64 {
    return t.A + t.B + t.C
}

// Função polimórfica — aceita qualquer Forma
func descreverForma(f Forma) {
    fmt.Printf("Tipo: %T | Área: %.2f | Perímetro: %.2f\n",
        f, f.Area(), f.Perimetro())
}

func main() {
    formas := []Forma{
        Circulo{Raio: 5},
        Retangulo{Largura: 8, Altura: 3},
        Triangulo{A: 3, B: 4, C: 5},
    }

    for _, f := range formas {
        descreverForma(f)
    }

    // Interface pode verificar se implementa
    var f Forma = Circulo{Raio: 10}
    fmt.Printf("\nÁrea do círculo: %.2f\n", f.Area())
}
```

---

## 2. Duck Typing

```go
package main

import "fmt"

// "Se anda como um pato e grasna como um pato, é um pato"
// Em Go: se tem os métodos certos, implementa a interface

type Escritor interface {
    Escrever(conteudo string) error
}

// Três tipos completamente diferentes, todos implementam Escritor
type ArquivoLog struct {
    caminho string
}

type SlackNotificador struct {
    canal string
}

type BancoDeDados struct {
    tabela string
}

func (a ArquivoLog) Escrever(conteudo string) error {
    fmt.Printf("[ARQUIVO %s] %s\n", a.caminho, conteudo)
    return nil
}

func (s SlackNotificador) Escrever(conteudo string) error {
    fmt.Printf("[SLACK #%s] %s\n", s.canal, conteudo)
    return nil
}

func (b BancoDeDados) Escrever(conteudo string) error {
    fmt.Printf("[DB %s] INSERT: %s\n", b.tabela, conteudo)
    return nil
}

// Esta função não sabe (e não precisa saber) qual é o tipo concreto
func publicarEvento(escritor Escritor, mensagem string) {
    if err := escritor.Escrever(mensagem); err != nil {
        fmt.Println("Erro ao publicar:", err)
    }
}

func main() {
    destinos := []Escritor{
        ArquivoLog{"/var/log/app.log"},
        SlackNotificador{"alertas"},
        BancoDeDados{"eventos"},
    }

    for _, d := range destinos {
        publicarEvento(d, "usuário fez login")
    }
}
```

> **💡 Insight:** Duck typing estrutural significa que tipos de pacotes externos podem implementar suas interfaces **sem modificação**. Isso é um superpoder — você pode criar abstrações sobre código que não controla.

---

## 3. Polimorfismo

```go
package main

import "fmt"

// Interfaces compostas
type Leitor interface {
    Ler() string
}

type EscritorSimples interface {
    Escrever(s string)
}

// Interface composta de outras interfaces
type LeitorEscritor interface {
    Leitor
    EscritorSimples
}

type Buffer struct {
    dados []string
}

func (b *Buffer) Ler() string {
    if len(b.dados) == 0 {
        return ""
    }
    val := b.dados[0]
    b.dados = b.dados[1:]
    return val
}

func (b *Buffer) Escrever(s string) {
    b.dados = append(b.dados, s)
}

// Comportamento polimórfico: diferentes implementações, mesmo contrato
type Processador interface {
    Processar(entrada string) string
}

type MaiusculoProcessor struct{}
type CriptografadorROT13 struct{}
type ContadorPalavras struct{}

func (m MaiusculoProcessor) Processar(s string) string {
    resultado := []rune{}
    for _, r := range s {
        if r >= 'a' && r <= 'z' {
            r -= 32
        }
        resultado = append(resultado, r)
    }
    return string(resultado)
}

func (c CriptografadorROT13) Processar(s string) string {
    resultado := []rune{}
    for _, r := range s {
        switch {
        case r >= 'a' && r <= 'z':
            resultado = append(resultado, 'a'+(r-'a'+13)%26)
        case r >= 'A' && r <= 'Z':
            resultado = append(resultado, 'A'+(r-'A'+13)%26)
        default:
            resultado = append(resultado, r)
        }
    }
    return string(resultado)
}

func (c ContadorPalavras) Processar(s string) string {
    palavras := 0
    emPalavra := false
    for _, r := range s {
        if r == ' ' || r == '\t' || r == '\n' {
            emPalavra = false
        } else if !emPalavra {
            emPalavra = true
            palavras++
        }
    }
    return fmt.Sprintf("%d palavra(s) encontrada(s)", palavras)
}

func pipeline(texto string, processadores ...Processador) []string {
    resultados := []string{}
    for _, p := range processadores {
        resultados = append(resultados, p.Processar(texto))
    }
    return resultados
}

func main() {
    buf := &Buffer{}
    buf.Escrever("primeira mensagem")
    buf.Escrever("segunda mensagem")
    fmt.Println(buf.Ler())
    fmt.Println(buf.Ler())

    texto := "Hello World Go"
    resultados := pipeline(texto,
        MaiusculoProcessor{},
        CriptografadorROT13{},
        ContadorPalavras{},
    )
    for i, r := range resultados {
        fmt.Printf("Processador %d: %s\n", i+1, r)
    }
}
```

---

## 4. Empty Interface (interface{})

```go
package main

import "fmt"

// interface{} aceita qualquer valor — use com moderação
func inspecionar(valores ...interface{}) {
    for i, v := range valores {
        fmt.Printf("[%d] tipo: %T | valor: %v\n", i, v, v)
    }
}

// Mapa genérico de configurações
type Config map[string]interface{}

func (c Config) obter(chave string) interface{} {
    return c[chave]
}

func main() {
    inspecionar("texto", 42, 3.14, true, []int{1, 2, 3}, nil)

    config := Config{
        "porta":     8080,
        "debug":     true,
        "host":      "localhost",
        "timeout":   30.5,
        "tags":      []string{"web", "api"},
    }

    for chave, valor := range config {
        fmt.Printf("  %s = %v (%T)\n", chave, valor, valor)
    }

    // Precisa de type assertion para usar o valor concreto
    if porta, ok := config["porta"].(int); ok {
        fmt.Printf("\nServidor na porta: %d\n", porta)
    }
}
```

> **💡 Insight:** `interface{}` é uma válvula de escape do sistema de tipos. Use quando realmente precisar de heterogeneidade (configurações, JSON, dados desconhecidos). Para novos códigos Go (1.18+), prefira `any` (alias) ou genéricos quando o tipo puder ser determinado.

---

## 5. Type Assertion e Type Switch com Interfaces

```go
package main

import (
    "fmt"
    "math"
)

type Geometria interface {
    Area() float64
}

type Circulo struct{ Raio float64 }
type Quadrado struct{ Lado float64 }
type Hexagono struct{ Lado float64 }

func (c Circulo) Area() float64  { return math.Pi * c.Raio * c.Raio }
func (q Quadrado) Area() float64 { return q.Lado * q.Lado }
func (h Hexagono) Area() float64 { return 3 * math.Sqrt(3) / 2 * h.Lado * h.Lado }

func analisarForma(g Geometria) {
    // Type switch para tratamento específico por tipo
    switch forma := g.(type) {
    case Circulo:
        fmt.Printf("Círculo (r=%.1f): área=%.2f, diâm=%.2f\n",
            forma.Raio, forma.Area(), 2*forma.Raio)
    case Quadrado:
        fmt.Printf("Quadrado (l=%.1f): área=%.2f, diag=%.2f\n",
            forma.Lado, forma.Area(), forma.Lado*math.Sqrt(2))
    case Hexagono:
        fmt.Printf("Hexágono (l=%.1f): área=%.2f\n",
            forma.Lado, forma.Area())
    default:
        fmt.Printf("Forma desconhecida: %T com área %.2f\n", forma, forma.Area())
    }
}

func main() {
    formas := []Geometria{
        Circulo{5},
        Quadrado{4},
        Hexagono{3},
    }

    for _, f := range formas {
        analisarForma(f)
    }

    // Type assertion — forma específica
    var g Geometria = Circulo{Raio: 7}
    if c, ok := g.(Circulo); ok {
        fmt.Printf("\nRaio do círculo: %.1f\n", c.Raio)
    }
}
```

---

## 6. any

```go
package main

import "fmt"

// 'any' é um alias para interface{} — introduzido no Go 1.18
// É exatamente equivalente, mas mais legível

func imprimirTudo(itens ...any) {
    for _, item := range itens {
        fmt.Printf("%v ", item)
    }
    fmt.Println()
}

type Cache struct {
    dados map[string]any
}

func NovoCache() *Cache {
    return &Cache{dados: make(map[string]any)}
}

func (c *Cache) Set(chave string, valor any) {
    c.dados[chave] = valor
}

func (c *Cache) Get(chave string) (any, bool) {
    v, ok := c.dados[chave]
    return v, ok
}

func main() {
    imprimirTudo(1, "dois", 3.0, true, []byte("quatro"))

    cache := NovoCache()
    cache.Set("usuario", "Lucas")
    cache.Set("idade", 30)
    cache.Set("ativo", true)

    if nome, ok := cache.Get("usuario"); ok {
        // Ainda precisa de type assertion para usar o valor concreto
        fmt.Println("Usuário:", nome.(string))
    }

    // Preferir any sobre interface{} em código novo — é mais expressivo
    var x any = 42
    fmt.Printf("Tipo: %T, Valor: %v\n", x, x)
}
```

---

## 🔁 Revisão Rápida — O que lembrar

| Conceito | Ponto chave |
|---|---|
| Interface | Define comportamento — qualquer tipo que implementa os métodos a satisfaz |
| Duck typing | Implementação é implícita — sem `implements` |
| Composição de interfaces | `type RW interface { Reader; Writer }` |
| `interface{}`/`any` | Aceita qualquer tipo — use com moderação |
| Type assertion `x.(T)` | Extrai o tipo concreto — sempre use a forma segura `v, ok := x.(T)` |
| Type switch | Ramifica por tipo concreto — mais limpo que múltiplas assertions |

**Regra de design:** Interfaces pequenas são melhores. Prefira `io.Reader` (1 método) a interfaces monstro com 10 métodos. Componha interfaces pequenas quando precisar de mais comportamento.
