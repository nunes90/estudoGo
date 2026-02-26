# Capítulo 4 — Tipos Compostos e Sistema de Tipos

> **Revisão rápida:** Este capítulo é o coração da modelagem de dados em Go. Arrays, slices e structs são os blocos de construção de quase toda aplicação real. O sistema de type assertion e type switch é o que permite trabalhar com tipos desconhecidos em tempo de execução.

---

## 1. Arrays

```go
package main

import "fmt"

func main() {
    // Array com tamanho fixo — faz parte do tipo!
    var diasSemana [7]string
    diasSemana[0] = "Segunda"
    diasSemana[6] = "Domingo"

    // Inicialização literal
    temperaturas := [5]float64{28.5, 31.0, 27.3, 29.8, 30.1}

    // Deixar o compilador contar
    vogais := [...]rune{'a', 'e', 'i', 'o', 'u'}

    fmt.Println("Dias:", diasSemana[0], "a", diasSemana[6])
    fmt.Println("Temperatura média:", media(temperaturas))
    fmt.Printf("Vogais: %c\n", vogais)

    // Arrays são comparáveis com ==
    a := [3]int{1, 2, 3}
    b := [3]int{1, 2, 3}
    c := [3]int{1, 2, 4}
    fmt.Println(a == b) // true
    fmt.Println(a == c) // false

    // Arrays são copiados ao serem passados para funções!
    original := [3]int{10, 20, 30}
    copia := original
    copia[0] = 999
    fmt.Println(original[0]) // ainda 10
}

func media(temps [5]float64) float64 {
    soma := 0.0
    for _, t := range temps {
        soma += t
    }
    return soma / float64(len(temps))
}
```

> **💡 Insight:** Arrays são raramente usados diretamente em Go. O tamanho faz parte do tipo (`[5]int` ≠ `[6]int`), o que os torna inflexíveis. Na prática, você usará **slices** quase sempre.

---

## 2. Slices

```go
package main

import (
    "fmt"
    "sort"
)

func main() {
    // Slice literal
    notas := []float64{8.5, 7.0, 9.2, 6.8, 10.0}

    // make: tamanho e capacidade explícitos
    buffer := make([]byte, 0, 64) // len=0, cap=64
    _ = buffer

    // append adiciona elementos
    tarefas := []string{}
    tarefas = append(tarefas, "Estudar Go")
    tarefas = append(tarefas, "Fazer exercícios")
    tarefas = append(tarefas, "Ler livro")
    tarefas = append(tarefas, "Dormir cedo")
    fmt.Println(tarefas)

    // Slicing (sub-slice)
    fmt.Println(tarefas[1:3])  // [Fazer exercícios Ler livro]
    fmt.Println(tarefas[:2])   // [Estudar Go Fazer exercícios]
    fmt.Println(tarefas[2:])   // [Ler livro Dormir cedo]

    // Sort
    sort.Float64s(notas)
    fmt.Println("Notas ordenadas:", notas)

    // Removendo elemento por índice (sem preservar ordem)
    i := 1
    notas[i] = notas[len(notas)-1]
    notas = notas[:len(notas)-1]
    fmt.Println("Após remoção:", notas)

    // Cópia segura (evita surpresas com slices compartilhados)
    original := []int{1, 2, 3}
    copia := make([]int, len(original))
    copy(copia, original)
    copia[0] = 999
    fmt.Println("Original:", original) // [1 2 3] — não afetado
    fmt.Println("Cópia:", copia)       // [999 2 3]

    // len vs cap
    s := make([]int, 3, 10)
    fmt.Printf("len=%d cap=%d\n", len(s), cap(s)) // len=3 cap=10
}
```

> **💡 Insight:** Slices são **referências** a arrays subjacentes. Duas slices podem compartilhar a mesma memória. Quando isso importa (ex: dados que vão ser modificados), use `copy()` para criar uma cópia independente.

---

## 3. Simple Custom Types

```go
package main

import "fmt"

// Tipo customizado sobre tipo primitivo
type Celsius float64
type Fahrenheit float64
type Kelvin float64

func (c Celsius) ToFahrenheit() Fahrenheit {
    return Fahrenheit(c*9/5 + 32)
}

func (c Celsius) ToKelvin() Kelvin {
    return Kelvin(c + 273.15)
}

type Metros float64
type Km float64

func (m Metros) ToKm() Km {
    return Km(m / 1000)
}

func main() {
    temp := Celsius(100)
    fmt.Printf("%.1f°C = %.1f°F = %.2fK\n",
        temp, temp.ToFahrenheit(), temp.ToKelvin())

    distancia := Metros(42195) // maratona
    fmt.Printf("Uma maratona tem %.3f km\n", distancia.ToKm())

    // O compilador impede confusão de tipos
    var t1 Celsius = 37
    // var t2 Fahrenheit = t1  // ERRO: tipos incompatíveis
    t2 := Fahrenheit(t1) // precisa de conversão explícita
    fmt.Println(t1, t2)
}
```

---

## 4. Structs

```go
package main

import "fmt"

// Struct com tags (útil para JSON, banco de dados, etc.)
type Produto struct {
    ID        int     `json:"id"`
    Nome      string  `json:"nome"`
    Preco     float64 `json:"preco"`
    Estoque   int     `json:"estoque"`
    Disponivel bool   `json:"disponivel"`
}

func NovoProduto(id int, nome string, preco float64, estoque int) Produto {
    return Produto{
        ID:         id,
        Nome:       nome,
        Preco:      preco,
        Estoque:    estoque,
        Disponivel: estoque > 0,
    }
}

func (p *Produto) Desconto(percentual float64) {
    p.Preco = p.Preco * (1 - percentual/100)
}

func (p Produto) String() string {
    return fmt.Sprintf("[%d] %s — R$ %.2f (estoque: %d)", p.ID, p.Nome, p.Preco, p.Estoque)
}

// Composição de structs (Go não tem herança)
type Endereco struct {
    Rua    string
    Cidade string
    CEP    string
}

type Cliente struct {
    Nome     string
    Email    string
    Endereco // embedding — campos e métodos promovidos
}

func main() {
    notebook := NovoProduto(1, "Notebook Pro", 4500.00, 10)
    fmt.Println(notebook)

    notebook.Desconto(10)
    fmt.Println("Com desconto:", notebook)

    // Struct literal vs ponteiro
    p1 := Produto{ID: 2, Nome: "Mouse", Preco: 89.90, Estoque: 50, Disponivel: true}
    p2 := &Produto{ID: 3, Nome: "Teclado", Preco: 199.00, Estoque: 0, Disponivel: false}

    fmt.Println(p1)
    fmt.Println(p2) // ponteiro ainda chama String()

    // Embedding
    cliente := Cliente{
        Nome:  "Lucas",
        Email: "lucas@exemplo.com",
        Endereco: Endereco{
            Rua:    "Av. Paulista",
            Cidade: "São Paulo",
            CEP:    "01310-100",
        },
    }
    // Acesso direto aos campos do Endereco (promovidos)
    fmt.Println(cliente.Cidade) // São Paulo
    fmt.Println(cliente.Endereco.CEP) // também funciona
}
```

---

## 5. Type Assertions

```go
package main

import "fmt"

func processarDado(dado interface{}) string {
    // Type assertion sem verificação (panic se errar)
    // s := dado.(string) — perigoso!

    // Type assertion segura com vírgula-ok
    if s, ok := dado.(string); ok {
        return fmt.Sprintf("String com %d caracteres: %q", len(s), s)
    }
    if n, ok := dado.(int); ok {
        return fmt.Sprintf("Inteiro: %d (par: %v)", n, n%2 == 0)
    }
    if f, ok := dado.(float64); ok {
        return fmt.Sprintf("Float: %.3f", f)
    }
    return "Tipo não reconhecido"
}

func main() {
    fmt.Println(processarDado("Golang"))
    fmt.Println(processarDado(42))
    fmt.Println(processarDado(3.14))
    fmt.Println(processarDado(true))
}
```

---

## 6. Type Switch

```go
package main

import "fmt"

type Cachorro struct{ Nome string }
type Gato struct{ Nome string }
type Passaro struct{ Nome string }

func fazerBarulho(animal interface{}) string {
    switch a := animal.(type) {
    case Cachorro:
        return fmt.Sprintf("%s diz: Au au!", a.Nome)
    case Gato:
        return fmt.Sprintf("%s diz: Miau!", a.Nome)
    case Passaro:
        return fmt.Sprintf("%s diz: Piu piu!", a.Nome)
    case string:
        return fmt.Sprintf("Texto recebido: %q", a)
    case nil:
        return "Nenhum animal fornecido"
    default:
        return fmt.Sprintf("Animal desconhecido: %T", a)
    }
}

func main() {
    animais := []interface{}{
        Cachorro{Nome: "Rex"},
        Gato{Nome: "Mimi"},
        Passaro{Nome: "Piu"},
        "uma string qualquer",
        nil,
        42,
    }

    for _, a := range animais {
        fmt.Println(fazerBarulho(a))
    }
}
```

---

## 7. Type Checker (verificações em tempo de compilação)

```go
package main

import "fmt"

// Go garante segurança de tipos em tempo de compilação
type IdUsuario int
type IdProduto int

func buscarUsuario(id IdUsuario) string {
    return fmt.Sprintf("Usuário #%d", id)
}

func buscarProduto(id IdProduto) string {
    return fmt.Sprintf("Produto #%d", id)
}

func main() {
    userID := IdUsuario(42)
    prodID := IdProduto(7)

    fmt.Println(buscarUsuario(userID))
    fmt.Println(buscarProduto(prodID))

    // buscarUsuario(prodID) — ERRO em tempo de compilação!
    // Isso é o type checker protegendo você de bugs
}
```

> **💡 Insight:** Criar tipos distintos para IDs (`IdUsuario` vs `IdProduto`) é um padrão poderoso. O compilador garante que você nunca passe um ID de produto onde um ID de usuário é esperado — bugs que seriam silenciosos em linguagens com tipagem fraca.

---

## 🔁 Revisão Rápida — O que lembrar

| Tipo | Característica Principal |
|---|---|
| Array `[N]T` | Tamanho fixo, copiado por valor |
| Slice `[]T` | Tamanho dinâmico, referência ao array |
| Struct | Composição de dados, sem herança — use embedding |
| Type assertion `x.(T)` | Converte interface para tipo concreto — use forma segura |
| Type switch | Ramifica por tipo — mais limpo que múltiplas assertions |
| Custom types | Segurança semântica em tempo de compilação |
