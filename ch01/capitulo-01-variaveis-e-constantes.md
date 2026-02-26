# Capítulo 1 — Variáveis, Operadores, Ponteiros e Constantes

> **Revisão rápida:** Este capítulo estabelece os fundamentos da linguagem. Entender ponteiros e a diferença entre valor e referência é crítico para tudo que vem depois.

---

## 1. Declarando Variáveis

Go oferece três formas principais de declarar variáveis:

```go
package main

import "fmt"

func main() {
    // Forma 1: var com tipo explícito
    var cidade string = "São Paulo"

    // Forma 2: var com inferência de tipo
    var populacao = 12_300_000

    // Forma 3: short declaration (só dentro de funções)
    pais := "Brasil"

    // Declaração múltipla
    var (
        latitude  float64 = -23.55
        longitude float64 = -46.63
    )

    fmt.Println(cidade, populacao, pais, latitude, longitude)
}
```

> **💡 Insight:** O `:=` é açúcar sintático — mas só funciona dentro de funções. No escopo de pacote, você é obrigado a usar `var`. Prefira `:=` no dia a dia pelo código mais limpo.

---

## 2. Alterando o Valor de uma Variável

```go
package main

import "fmt"

func main() {
    pontuacao := 0

    // Atribuição simples
    pontuacao = 10

    // Atribuição composta
    pontuacao += 5  // 15
    pontuacao *= 2  // 30
    pontuacao -= 3  // 27

    fmt.Println("Pontuação final:", pontuacao)

    // Troca de valores (idiomático em Go)
    a, b := "X", "Y"
    a, b = b, a
    fmt.Println(a, b) // Y X
}
```

---

## 3. Operadores

```go
package main

import "fmt"

func main() {
    // Aritméticos
    fmt.Println(10 + 3)   // 13
    fmt.Println(10 - 3)   // 7
    fmt.Println(10 * 3)   // 30
    fmt.Println(10 / 3)   // 3  (divisão inteira!)
    fmt.Println(10 % 3)   // 1  (resto)

    // Comparação
    fmt.Println(10 > 3)   // true
    fmt.Println(10 == 10) // true
    fmt.Println(10 != 5)  // true

    // Lógicos
    ativo := true
    admin := false
    fmt.Println(ativo && admin)  // false
    fmt.Println(ativo || admin)  // true
    fmt.Println(!ativo)          // false

    // Bitwise (muito usado em flags e sistemas)
    permissao := 0b0100
    leitura   := 0b0100
    fmt.Println(permissao & leitura) // 4 (tem permissão de leitura)
}
```

> **💡 Insight:** A divisão inteira em Go é uma fonte comum de bugs para iniciantes. `10 / 3` retorna `3`, não `3.33`. Se precisar do resultado float, converta: `float64(10) / float64(3)`.

---

## 4. Valor versus Ponteiro

Esta é uma das distinções mais importantes em Go.

```go
package main

import "fmt"

// Recebe uma cópia — não modifica o original
func dobrarValor(n int) {
    n = n * 2
    fmt.Println("Dentro da função:", n)
}

// Recebe o endereço — modifica o original
func dobrarPonteiro(n *int) {
    *n = *n * 2
    fmt.Println("Dentro da função:", *n)
}

func main() {
    numero := 10

    dobrarValor(numero)
    fmt.Println("Após dobrarValor:", numero) // ainda 10

    dobrarPonteiro(&numero)
    fmt.Println("Após dobrarPonteiro:", numero) // agora 20

    // Criando ponteiro com new
    p := new(int)
    *p = 42
    fmt.Println("Valor via ponteiro:", *p)
    fmt.Println("Endereço de memória:", p)
}
```

> **💡 Insight:** Use ponteiros quando:
> - Precisar que a função modifique o valor original
> - Estiver trabalhando com structs grandes (evita cópias caras)
> - O valor puder ser `nil` (ausência de valor)
>
> Evite ponteiros desnecessários — Go passa por valor por padrão por uma boa razão: é mais seguro e previsível.

---

## 5. Constantes

```go
package main

import "fmt"

// Constantes simples
const Pi = 3.14159
const AppVersion = "1.0.0"
const MaxTentativas = 3

// Grupo de constantes
const (
    KB = 1024
    MB = 1024 * KB
    GB = 1024 * MB
)

func main() {
    arquivo := 2.5 * GB
    fmt.Printf("Arquivo ocupa %.2f bytes\n", float64(arquivo))
    fmt.Println("Versão:", AppVersion)
}
```

> **💡 Insight:** Constantes em Go são avaliadas em tempo de compilação, o que as torna mais eficientes que variáveis. O compilador também garante que seu valor nunca muda — use-as para valores que representam conceitos fixos do seu domínio.

---

## 6. Enums (com iota)

Go não tem um tipo `enum` nativo, mas simula com `iota`:

```go
package main

import "fmt"

type StatusPedido int

const (
    Aguardando  StatusPedido = iota // 0
    Confirmado                       // 1
    Enviado                          // 2
    Entregue                         // 3
    Cancelado                        // 4
)

func descricaoStatus(s StatusPedido) string {
    switch s {
    case Aguardando:
        return "Aguardando pagamento"
    case Confirmado:
        return "Pedido confirmado"
    case Enviado:
        return "Em rota de entrega"
    case Entregue:
        return "Entregue com sucesso"
    case Cancelado:
        return "Pedido cancelado"
    default:
        return "Status desconhecido"
    }
}

func main() {
    meuPedido := Enviado
    fmt.Println(descricaoStatus(meuPedido)) // Em rota de entrega

    // iota com operações
    type ByteSize float64
    const (
        _           = iota // ignora o primeiro valor
        KiloBytes ByteSize = 1 << (10 * iota) // 1024
        MegaBytes                               // 1048576
        GigaBytes                               // 1073741824
    )
    fmt.Printf("1 KB = %.0f bytes\n", float64(KiloBytes))
    fmt.Printf("1 MB = %.0f bytes\n", float64(MegaBytes))
}
```

> **💡 Insight:** O padrão de criar um tipo personalizado para o enum (como `StatusPedido int`) é importante — ele impede que você passe um `int` qualquer onde um status é esperado, adicionando segurança de tipos ao seu código.

---

## 🔁 Revisão Rápida — O que lembrar

| Conceito | Ponto chave |
|---|---|
| `:=` | Só dentro de funções, infere o tipo |
| `var` | Escopo de pacote ou quando quer ser explícito |
| Ponteiro `*` e `&` | `&` pega o endereço, `*` acessa o valor |
| `const` + `iota` | Simula enums de forma idiomática |
| Divisão inteira | `10/3 == 3`, não `3.33` |
