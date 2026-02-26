# Capítulo 2 — Controle de Fluxo

> **Revisão rápida:** Go tem uma abordagem minimalista para controle de fluxo — sem `while`, sem `do-while`, sem `ternário`. Tudo é feito com `if`, `switch` e `for`. Menos keywords, mais clareza.

---

## 1. If Statements

```go
package main

import "fmt"

func classificarIdade(idade int) string {
    // if com inicialização de variável (escopo limitado ao bloco)
    if anos := idade; anos < 0 {
        return "Idade inválida"
    } else if anos < 13 {
        return "Criança"
    } else if anos < 18 {
        return "Adolescente"
    } else if anos < 60 {
        return "Adulto"
    } else {
        return "Idoso"
    }
}

func dividir(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("divisão por zero não permitida")
    }
    return a / b, nil
}

func main() {
    fmt.Println(classificarIdade(25))  // Adulto
    fmt.Println(classificarIdade(8))   // Criança
    fmt.Println(classificarIdade(-1))  // Idade inválida

    // Padrão idiomático: if err != nil
    resultado, err := dividir(10, 3)
    if err != nil {
        fmt.Println("Erro:", err)
        return
    }
    fmt.Printf("Resultado: %.2f\n", resultado)
}
```

> **💡 Insight:** O padrão `if err != nil` é onipresente em Go. Aprenda a amar ele — é explícito e elimina exceções ocultas que tornam outros códigos imprevisíveis.

---

## 2. Expression Switch

```go
package main

import (
    "fmt"
    "time"
)

func diaDaSemana(d time.Weekday) string {
    switch d {
    case time.Monday:
        return "Segunda-feira — semana começou!"
    case time.Friday:
        return "Sexta-feira — quase lá!"
    case time.Saturday, time.Sunday:
        return "Final de semana!"
    default:
        return "Dia de trabalho comum"
    }
}

func categorizarNota(nota int) string {
    // Switch sem expressão = substituto elegante para if-else
    switch {
    case nota >= 90:
        return "A"
    case nota >= 80:
        return "B"
    case nota >= 70:
        return "C"
    case nota >= 60:
        return "D"
    default:
        return "F"
    }
}

func main() {
    fmt.Println(diaDaSemana(time.Monday))   // Segunda-feira — semana começou!
    fmt.Println(diaDaSemana(time.Saturday)) // Final de semana!

    fmt.Println(categorizarNota(95)) // A
    fmt.Println(categorizarNota(72)) // C
    fmt.Println(categorizarNota(50)) // F
}
```

> **💡 Insight:** Diferente de C ou Java, o `switch` em Go **não precisa de `break`** — cada caso termina automaticamente. Se quiser que a execução "caia" para o próximo caso, use `fallthrough` (raro e geralmente desaconselhado).

---

## 3. Loops (apenas `for`)

Go tem apenas uma estrutura de loop: o `for`. Mas ele é versátil o suficiente para substituir todos os outros.

```go
package main

import "fmt"

func main() {
    // Loop clássico
    for i := 0; i < 5; i++ {
        fmt.Printf("Iteração %d\n", i)
    }

    // Equivalente ao while
    contador := 10
    for contador > 0 {
        fmt.Println(contador)
        contador -= 3
    }

    // Loop infinito com saída controlada
    tentativas := 0
    for {
        tentativas++
        if tentativas >= 3 {
            fmt.Println("Número máximo de tentativas atingido")
            break
        }
    }

    // Range sobre slice
    frutas := []string{"manga", "abacate", "goiaba"}
    for indice, fruta := range frutas {
        fmt.Printf("[%d] %s\n", indice, fruta)
    }

    // Range sobre map
    capitais := map[string]string{
        "BR": "Brasília",
        "JP": "Tóquio",
        "DE": "Berlim",
    }
    for pais, capital := range capitais {
        fmt.Printf("%s → %s\n", pais, capital)
    }

    // Range apenas com índice (valor descartado)
    numeros := []int{10, 20, 30}
    for i := range numeros {
        numeros[i] *= 2
    }
    fmt.Println(numeros) // [20 40 60]
}
```

---

## 4. Break e Continue

```go
package main

import "fmt"

func main() {
    // continue: pula para a próxima iteração
    fmt.Println("Números ímpares de 1 a 10:")
    for i := 1; i <= 10; i++ {
        if i%2 == 0 {
            continue // pula os pares
        }
        fmt.Println(i)
    }

    // break: interrompe o loop
    fmt.Println("\nBusca no slice:")
    nomes := []string{"Alice", "Bruno", "Carlos", "Diana"}
    busca := "Carlos"
    encontrado := false
    for _, nome := range nomes {
        if nome == busca {
            encontrado = true
            break
        }
    }
    fmt.Printf("'%s' encontrado: %v\n", busca, encontrado)

    // break com label: sai de loop aninhado
    fmt.Println("\nMatrix com break em label:")
externo:
    for linha := 0; linha < 3; linha++ {
        for coluna := 0; coluna < 3; coluna++ {
            if linha == 1 && coluna == 1 {
                fmt.Println("Centro encontrado, saindo!")
                break externo
            }
            fmt.Printf("(%d,%d) ", linha, coluna)
        }
    }
}
```

> **💡 Insight:** Labels com `break` são raros, mas legítimos em Go quando você precisa sair de loops aninhados. É preferível a usar variáveis booleanas de controle.

---

## 5. Goto

```go
package main

import "fmt"

func main() {
    i := 0

volta:
    if i < 5 {
        fmt.Println("i =", i)
        i++
        goto volta
    }

    fmt.Println("Loop terminou com goto")
}
```

> **💡 Insight:** `goto` existe em Go, mas é **raramente justificado**. O único caso de uso legítimo comum é pular para limpeza de recursos em funções longas (mas `defer` geralmente resolve melhor). Evite `goto` — ele torna o fluxo difícil de seguir.

---

## 🔁 Revisão Rápida — O que lembrar

| Conceito | Ponto chave |
|---|---|
| `if` com inicialização | `if v := calc(); v > 0 { }` — `v` só existe no bloco |
| `switch` sem `break` | Cada case termina sozinho. Use `fallthrough` com cautela |
| `for` universal | Substitui `while`, `do-while` e `for` clássico |
| `range` | Itera sobre slices, maps, strings e channels |
| `break` com label | Sai de loops aninhados sem variável de controle |
| `goto` | Existe, mas evite — `defer` e funções resolvem melhor |
