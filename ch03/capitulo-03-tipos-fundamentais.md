# Capítulo 3 — Tipos Fundamentais

> **Revisão rápida:** Go é estritamente tipado. Entender os tipos primitivos — especialmente a distinção entre `byte`, `rune` e `string` — é essencial para trabalhar com texto corretamente, especialmente em português e outros idiomas com caracteres especiais.

---

## 1. True e False (Booleanos)

```go
package main

import "fmt"

func main() {
    var ativo bool // zero value = false
    aprovado := true

    fmt.Println(ativo)    // false
    fmt.Println(aprovado) // true

    // Go não converte automaticamente outros tipos para bool
    // Isso NÃO compila: if 1 { }
    // É preciso ser explícito:
    contador := 0
    if contador == 0 {
        fmt.Println("Contador está zerado")
    }

    // Operações lógicas retornam bool
    temperatura := 35.0
    chovendo := false
    precisaGuardaChuva := chovendo || temperatura > 38
    fmt.Println("Precisa guarda-chuva?", precisaGuardaChuva) // false
}
```

> **💡 Insight:** A ausência de conversão implícita para `bool` em Go é intencional. Elimina bugs clássicos de linguagens como C onde `if (ptr)` pode ter comportamento inesperado.

---

## 2. Números

```go
package main

import (
    "fmt"
    "math"
)

func main() {
    // Inteiros com tamanho explícito
    var idade int8  = 25     // -128 a 127
    var porta int16 = 8080   // -32768 a 32767
    var usuarios int32 = 1_000_000
    var arquivoBytes int64 = 9_000_000_000

    // int = 64 bits em sistemas 64-bit (use por padrão)
    var itens int = 42

    // Sem sinal (unsigned)
    var r uint8  = 255  // byte é alias de uint8
    var g uint8  = 128
    var b uint8  = 0

    // Floats
    var preco float32 = 29.99
    var precisao float64 = math.Pi

    fmt.Println(idade, porta, usuarios, arquivoBytes, itens)
    fmt.Printf("Cor RGB: (%d, %d, %d)\n", r, g, b)
    fmt.Printf("Preço: %.2f\n", preco)
    fmt.Printf("Pi com 10 casas: %.10f\n", precisao)

    // Conversão explícita obrigatória
    var x int = 10
    var y float64 = float64(x) * 1.5
    fmt.Println(y) // 15.0

    // Constantes numéricas não tipadas são flexíveis
    const fator = 2.5
    resultadoInt := int(fator * float64(x))
    fmt.Println(resultadoInt) // 25
}
```

> **💡 Insight:** Use `int` e `float64` como padrão na maioria dos casos. Os tipos com tamanho específico (`int32`, `float32`) são úteis quando você está interfaceando com protocolos binários, arquivos ou APIs que especificam tamanhos exatos.

---

## 3. Byte

```go
package main

import "fmt"

func main() {
    // byte é um alias para uint8
    var b byte = 65
    fmt.Println(b)          // 65
    fmt.Printf("%c\n", b)  // A (valor ASCII)

    // Strings são slices de bytes internamente
    texto := "Gopher"
    fmt.Println("Bytes de 'Gopher':")
    for i, byt := range []byte(texto) {
        fmt.Printf("  [%d] %d = '%c'\n", i, byt, byt)
    }

    // Manipulando bytes
    dados := []byte{72, 101, 108, 108, 111}
    fmt.Println(string(dados)) // Hello

    // Útil para trabalhar com dados binários
    mascara := byte(0b11110000)
    valor := byte(0b10100101)
    resultado := mascara & valor
    fmt.Printf("Resultado da máscara: %08b\n", resultado)
}
```

---

## 4. Text (strings)

```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    // Strings são imutáveis em Go
    nome := "Lucas"
    sobrenome := "Nunes"

    // Concatenação
    nomeCompleto := nome + " " + sobrenome
    fmt.Println(nomeCompleto)

    // Formatação
    apresentacao := fmt.Sprintf("Olá, meu nome é %s e tenho %d anos", nome, 30)
    fmt.Println(apresentacao)

    // Strings multiline com raw literal
    json := `{
    "nome": "Lucas",
    "cidade": "São Paulo"
}`
    fmt.Println(json)

    // Operações comuns
    frase := "  Go é incrível!  "
    fmt.Println(strings.TrimSpace(frase))
    fmt.Println(strings.ToUpper(frase))
    fmt.Println(strings.Contains(frase, "incrível"))
    fmt.Println(strings.Replace(frase, "Go", "Golang", 1))
    partes := strings.Split("a,b,c,d", ",")
    fmt.Println(partes) // [a b c d]

    // Iterando por bytes (cuidado com UTF-8!)
    for i := 0; i < len(nome); i++ {
        fmt.Printf("%c ", nome[i])
    }
    fmt.Println()
}
```

> **💡 Insight:** Strings em Go são **imutáveis** e **sequências de bytes** (não de caracteres). Para texto com acentos e caracteres especiais, use `range` com `rune` — não `len()` com índice de byte.

---

## 5. Rune

```go
package main

import (
    "fmt"
    "unicode/utf8"
)

func main() {
    // rune é alias para int32, representa um ponto Unicode
    var letra rune = 'ã'
    fmt.Println(letra)          // 227 (código Unicode)
    fmt.Printf("%c\n", letra)  // ã

    // O problema com bytes e acentos
    texto := "Olá, São Paulo"
    fmt.Println("Bytes:", len(texto))                       // mais que 14!
    fmt.Println("Caracteres:", utf8.RuneCountInString(texto)) // 14

    // Iteração correta com range (usa runes automaticamente)
    fmt.Println("\nCaractere por caractere:")
    for posicao, caractere := range texto {
        fmt.Printf("  pos %d: %c (U+%04X)\n", posicao, caractere, caractere)
    }

    // Convertendo string para []rune para manipulação segura
    runes := []rune(texto)
    fmt.Println("\nPrimeiro caractere:", string(runes[0]))   // O
    fmt.Println("Último caractere:", string(runes[len(runes)-1])) // o
}
```

> **💡 Insight:** Esta é a distinção mais importante para devs brasileiros: `len("São")` retorna **5**, não 3 — porque `ã` ocupa 2 bytes em UTF-8. Use sempre `[]rune` ou `utf8.RuneCountInString` quando precisar contar caracteres de verdade.

---

## 6. O Valor nil

```go
package main

import "fmt"

func buscarUsuario(id int) *string {
    usuarios := map[int]string{
        1: "Alice",
        2: "Bruno",
    }
    if nome, ok := usuarios[id]; ok {
        return &nome
    }
    return nil // ausência de valor
}

func main() {
    // nil é o zero value de: ponteiros, slices, maps, funções, interfaces, channels
    var ponteiro *int
    var slice []int
    var mapa map[string]int
    var funcao func()

    fmt.Println(ponteiro == nil) // true
    fmt.Println(slice == nil)    // true
    fmt.Println(mapa == nil)     // true
    fmt.Println(funcao == nil)   // true

    // Verificando nil antes de usar
    resultado := buscarUsuario(1)
    if resultado != nil {
        fmt.Println("Usuário encontrado:", *resultado)
    }

    resultado = buscarUsuario(99)
    if resultado == nil {
        fmt.Println("Usuário não encontrado")
    }

    // Slice nil vs slice vazio — diferença importante!
    var nilSlice []int
    emptySlice := []int{}

    fmt.Println(nilSlice == nil)   // true
    fmt.Println(emptySlice == nil) // false
    fmt.Println(len(nilSlice))     // 0
    fmt.Println(len(emptySlice))   // 0
}
```

> **💡 Insight:** Slice `nil` e slice vazio (`[]int{}`) se comportam de forma idêntica com `len`, `cap`, `append` e `range` — mas não são iguais na comparação com `nil`. Prefira retornar `nil` em vez de slice vazio quando a ausência de dados é relevante semanticamente.

---

## 🔁 Revisão Rápida — O que lembrar

| Tipo | Zero Value | Observação |
|---|---|---|
| `bool` | `false` | Sem conversão implícita de int |
| `int`, `float64` | `0`, `0.0` | Use estes por padrão |
| `byte` | `0` | Alias de `uint8` |
| `string` | `""` | Imutável, sequência de bytes |
| `rune` | `0` | Alias de `int32`, um ponto Unicode |
| ponteiros, slices, maps, funcs | `nil` | Ausência de valor |

**Regra de ouro para texto em português:** Sempre itere com `range` (usa runes), nunca com índice de byte quando houver acentos.
