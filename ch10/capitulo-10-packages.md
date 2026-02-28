# Capítulo 10 – Packages Keep Projects Manageable

> **Livro:** Go Programming – From Beginner to Professional (2ª ed.) | Samantha Coyle
> **Parte 3:** Modules | **Páginas:** 313–339

---

## Visão Geral

O Capítulo 10 apresenta um dos pilares fundamentais do desenvolvimento em Go: os **packages** (pacotes). O objetivo central é mostrar como dividir o código em pacotes torna projetos mais **fáceis de manter**, **reutilizáveis** e **modulares**. O capítulo percorre desde o conceito básico de pacote até recursos avançados como código exportado/não exportado, alias de pacotes e a função especial `init()`.

---

## 1. Introdução – Por que precisamos de Packages?

Em programas pequenos, é comum concentrar todo o código em um único arquivo `main.go` dentro do pacote `main`. À medida que os projetos crescem, isso se torna inviável: arquivos gigantes são difíceis de ler, modificar e reutilizar.

Go resolve esse problema adotando o princípio **DRY** (*Don't Repeat Yourself*): ao invés de duplicar código entre projetos, você o empacota em unidades reutilizáveis chamadas **packages**. A própria biblioteca padrão do Go é um ótimo exemplo disso — pacotes como `fmt`, `strings`, `os` e `math` organizam funções relacionadas em unidades coesas.

Um exemplo concreto é o pacote `strings` da stdlib, que contém múltiplos arquivos (`builder.go`, `compare.go`, `reader.go`, `replace.go`, `search.go`, `strings.go`), cada um focado em uma responsabilidade dentro da manipulação de strings.

---

## 2. As Três Qualidades que Packages Promovem

### 2.1 Maintainable (Manutenível)

Código manutenível é aquele fácil de modificar, com baixo risco de efeitos colaterais. À medida que o software evolui ao longo do **SDLC** (*Software Development Life Cycle*), o custo de manutenção cresce — especialmente quando o código está mal organizado. Pacotes bem estruturados reduzem esse custo ao isolar responsabilidades.

### 2.2 Reusable (Reutilizável)

Código reutilizável pode ser aproveitado em novos projetos sem ser reescrito. Os benefícios incluem:
- Redução de custo em projetos futuros
- Entrega mais rápida (sem reinventar a roda)
- Maior qualidade (código mais testado, mais usado)
- Mais tempo para inovação
- Base sólida para projetos futuros

### 2.3 Modular (Modular)

Código modular significa que cada tarefa do sistema tem seu lugar definido. Sem modularidade, encontrar e entender uma funcionalidade específica em uma base de código grande é quase impossível. Packages são o mecanismo do Go para alcançar modularidade: cada conjunto de funções relacionadas vive em um pacote próprio.

```
         maintainable
               |
           package
          /         \
    reusable      modular
```

---

## 3. O que é um Package?

Um **package** em Go é essencialmente um **diretório** que contém um ou mais arquivos-fonte `.go` com código relacionado. Ele expõe apenas as partes necessárias para quem o usa — o restante fica encapsulado internamente.

A progressão natural de organização do código é:

```
funções → arquivos-fonte (.go) → packages
```

### 3.1 Estrutura de um Package

| Componente | Descrição |
|---|---|
| Diretório | Pasta que agrupa os arquivos |
| Um ou mais arquivos `.go` | Arquivos-fonte contendo o código |
| Código relacionado | Todo código dentro do pacote deve ter um propósito comum |

Regra fundamental: **todos os arquivos de um package devem estar no mesmo diretório**.

### 3.2 Nomenclatura de Packages

O nome do pacote funciona como **autodocumentação** — deve comunicar claramente seu propósito. As regras e boas práticas são:

**Regras obrigatórias:**
- Sempre em **letras minúsculas**
- Sem underscores (`_`)
- Sem camelCase

**Boas práticas:**
- Nomes **curtos e concisos** (substantivos simples)
- Abreviações são bem-vindas se conhecidas na comunidade
- Evitar nomes **genéricos** como `misc`, `util`, `common`, `data`

| ❌ Ruim | ✅ Bom |
|---|---|
| `stringconversion` | `strconv` |
| `synchronizationprimitives` | `sync` |
| `measuringtime` | `time` |
| `StringConversion` | `strings` |
| `synchronization_primitives` | `regexp` |

### 3.3 Declaração de Package

A **primeira linha** de todo arquivo `.go` deve ser a declaração do pacote:

```go
package <nomeDoPacote>
```

Todos os arquivos do mesmo pacote compartilham a mesma declaração. Por exemplo, os arquivos `builder.go`, `compare.go` e `replace.go` do pacote `strings` todos começam com:

```go
package strings
```

Internamente, todas as funções, tipos e variáveis declarados nos arquivos de um pacote são acessíveis entre si, mesmo que estejam em arquivos diferentes.

---

## 4. Código Exportado e Não Exportado

Esta é uma das regras mais importantes do Go: a visibilidade de código fora do pacote é controlada **exclusivamente pela capitalização do nome**.

| Começa com | Visibilidade | Nome técnico |
|---|---|---|
| Letra **maiúscula** | Visível fora do pacote | **Exported** (exportado) |
| Letra **minúscula** | Visível apenas dentro do pacote | **Unexported** (não exportado) |

Não existem modificadores de acesso como `public`, `private` ou `protected` em Go. A capitalização é tudo.

### Exemplo – Código Exportado

```go
// strings.go (pacote strings da stdlib)
func Contains(s, substr string) bool {  // ← maiúscula = exportado
    return Index(s, substr) >= 0
}
```

Para usar uma função exportada de outro pacote, o acesso é feito com a notação `pacote.Função`:

```go
package main

import (
    "strings"
    "fmt"
)

func main() {
    str := "found me"
    if strings.Contains(str, "found") {  // strings.Contains → exportado
        fmt.Println("value found in str")
    }
}
```

### Exemplo – Código Não Exportado

```go
// strings.go (pacote strings da stdlib)
func explode(s string, n int) []string {  // ← minúscula = não exportado
    // ...
}
```

Tentar chamar `strings.explode()` de fora do pacote resulta em erro de compilação:

```
prog.go:10:9: cannot refer to unexported name strings.explode
prog.go:10:9: undefined: strings.explode
Go build failed.
```

> **Boa prática:** Exponha apenas o que outros pacotes precisam. Mantenha oculto tudo o que é detalhe de implementação interna.

### 4.1 Package Alias (Apelido de Pacote)

Go permite criar **aliases** (apelidos) para pacotes importados. Isso é útil quando:
- O nome do pacote não é claro o suficiente
- O nome é muito longo
- Dois pacotes importados têm o mesmo nome

**Sintaxe:**

```go
import alias "caminho/do/pacote"
```

**Exemplo:**

```go
package main

import (
    f "fmt"  // "fmt" agora se chama "f"
)

func main() {
    f.Println("Hello, Gophers")  // usa o alias f
}
```

### 4.2 O Package `main` – Pacote Especial

Existem dois tipos de pacotes em Go:

| Tipo | Características |
|---|---|
| **Executable** (executável) | É o pacote `main`; requer a função `main()`; ao rodar `go build`, gera um binário |
| **Non-executable** (não executável) | Qualquer outro pacote; quando compilado, não gera binário ou código executável |

O pacote `main` é especial porque:
- É **obrigatório** ter uma função `main()`
- A função `main()` é o **ponto de entrada** do programa
- O binário gerado pelo `go build` recebe o nome do diretório onde o `main` está
- Só pode haver **uma** função `main()` no pacote

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello Gophers!")
}
```

---

## 5. Exercício 10.01 – Criando um Package para Calcular Áreas de Formas

Este exercício demonstra na prática como mover código para um pacote customizado. O código de cálculo de áreas (triângulos, retângulos e quadrados), que estava em `main.go`, é migrado para um pacote chamado `shape`.

**Estrutura de diretórios resultante:**

```
Exercise10.01/
├── cmd/
│   └── main.go       ← package main (executável)
├── pkg/
│   └── shape/
│       ├── shape.go      ← package shape (não executável)
│       └── shape_test.go
└── go.mod
```

**Pontos-chave da migração:**
- O arquivo `shape.go` começa com `package shape` (não-executável)
- Os tipos e funções que precisam ser usados pelo `main` devem ser **exportados** (nomes com maiúscula):
  - `Shape` (interface), `Triangle`, `Rectangle`, `Square` (structs)
  - `PrintShapeDetails()` (função)
- Métodos internos como `area()` e `name()` permanecem com minúscula (não precisam ser visíveis fora)
- O `main.go` importa o pacote usando o caminho completo do módulo: `import "exercise10.01/pkg/shape"`

```go
// main.go
package main

import "exercise10.01/pkg/shape"

func main() {
    t := shape.Triangle{Base: 15.5, Height: 20.1}
    r := shape.Rectangle{Length: 20, Width: 10}
    s := shape.Square{Side: 10}
    shape.PrintShapeDetails(t, r, s)
}
```

**Saída esperada:**
```
The area of Triangle is: 155.78
The area of Rectangle is: 200.00
The area of Square is: 100.00
```

---

## 6. A Função `init()`

A função `init()` é uma **função especial** em Go que é executada **automaticamente antes** da função `main()`. Ela é usada para realizar inicializações que o pacote precisa antes de começar a executar a lógica principal.

### 6.1 Para que serve o `init()`?

Casos de uso comuns:
- Configurar conexões com banco de dados
- Inicializar variáveis de pacote
- Criar arquivos necessários
- Carregar dados de configuração
- Verificar ou reparar o estado do programa

### 6.2 Ordem de Execução

O Go segue uma ordem bem definida de inicialização:

```
1. Pacotes importados são inicializados (recursivamente)
        ↓
2. Variáveis de nível de pacote são inicializadas
        ↓
3. Função init() do pacote é chamada
        ↓
4. Função init() do pacote main é chamada
        ↓
5. Função main() é executada
```

### 6.3 Regras do `init()`

- **Não pode ter parâmetros** (nenhum argumento)
- **Não pode ter valores de retorno**
- É chamada **automaticamente** pelo runtime do Go (nunca chamada diretamente)
- **Pode haver múltiplos** `init()` no mesmo arquivo ou pacote

```go
// ❌ ERRADO – init() não aceita argumentos
func init(age int) { ... }  // Erro de compilação

// ✅ CORRETO
func init() { ... }
```

### 6.4 Exemplo Básico de `init()`

```go
package main

import "fmt"

var name = "Gopher"  // 1º: variável de pacote é inicializada

func init() {
    fmt.Println("Hello,", name)  // 2º: init() é executado
}

func main() {
    fmt.Println("Hello, main function")  // 3º: main() é executado
}
```

**Saída:**
```
Hello, Gopher
Hello, main function
```

### 6.5 Múltiplos `init()` – Exercício 10.02 e 10.03

Um pacote pode ter **mais de uma** função `init()`. Elas são executadas na **ordem em que aparecem no código**.

```go
package main

import "fmt"

var name = "Gopher"

func init() {
    fmt.Println("Hello,", name)  // 1º init → executado primeiro
}

func init() {
    fmt.Println("Second")  // 2º init → executado segundo
}

func init() {
    fmt.Println("Third")   // 3º init → executado terceiro
}

func main() {
    fmt.Println("Hello, main function")  // executado por último
}
```

**Saída:**
```
Hello, Gopher
Second
Third
Hello, main function
```

**Exercício 10.02 – Carregando categorias de orçamento:**
Demonstra o uso de `init()` para popular um mapa global com categorias de orçamento antes da execução do `main()`. O `init()` preenche o mapa e o `main()` apenas o itera e imprime.

```go
var budgetCategories = make(map[int]string)

func init() {
    fmt.Println("Initializing our budgetCategories")
    budgetCategories[1] = "Car Insurance"
    budgetCategories[2] = "Mortgage"
    budgetCategories[3] = "Electricity"
    // ...
}

func main() {
    for k, v := range budgetCategories {
        fmt.Printf("key: %d, value: %s\n", k, v)
    }
}
```

**Exercício 10.03 – Associando pagadores a categorias:**
Expande o exemplo anterior com **dois** `init()` functions: o primeiro inicializa o mapa de categorias, o segundo cria um mapa de pagadores e os associa às categorias. O `main()` imprime as associações finais.

---

## 7. Atividade 10.01 – Refatoração com Packages

A atividade final do capítulo une todos os conceitos aprendidos. O objetivo é pegar o código de cálculo de salário e avaliação de desempenho (desenvolvido no Capítulo 7 com interfaces) e **refatorá-lo usando packages**.

**Tarefas:**
1. Mover os tipos `Developer`, `Employee` e `Manager` para um pacote próprio em `pkg/payroll`
2. Nomear o pacote `payroll`
3. Separar os tipos e seus métodos em arquivos diferentes dentro do pacote (boa prática de organização)
4. Exportar corretamente tipos e métodos que precisam ser acessíveis de fora
5. Criar o `main()` usando o pacote `payroll`
6. Usar dois `init()` no `main`: um para exibir mensagem de boas-vindas, outro para inicializar variáveis

**Saída esperada:**
```
Welcome to the Employee Pay and Performance Review
++++++++++++++++++++++++++++++++++++++++++++++++++
Initializing variables
Eric Davis got a review rating of 2.80
Eric Davis got paid 84000.00 for the year
Mr. Boss got paid 160500.00 for the year
```

---

## 8. Resumo do Capítulo

| Conceito | O que é | Como funciona |
|---|---|---|
| **Package** | Diretório com arquivos `.go` relacionados | Agrupa código por responsabilidade |
| **Declaração** | `package <nome>` | Primeira linha de todo arquivo `.go` |
| **Exportado** | Nome começa com maiúscula | Visível fora do pacote |
| **Não exportado** | Nome começa com minúscula | Visível apenas dentro do pacote |
| **Alias** | `import f "fmt"` | Renomeia pacote localmente |
| **Package main** | Pacote executável | Requer `main()`, gera binário |
| **`init()`** | Função de inicialização | Executa antes de `main()`, sem args, sem retorno |
| **Múltiplos `init()`** | Vários `init()` no mesmo pacote | Executados em ordem de aparição |

---

## 9. Boas Práticas Consolidadas

- **Nomeie pacotes com substantivos simples, em minúsculas, sem underscores**
- **Exponha apenas o que é necessário** — deixe o restante não exportado (encapsulamento)
- **Organize arquivos por responsabilidade** dentro do mesmo pacote
- **Use `init()` para configurações que devem acontecer antes do `main()`**
- **Cuidado com a ordem dos múltiplos `init()`** — resultados inesperados podem ocorrer se a ordem importar
- **Evite nomes genéricos** como `util`, `common`, `misc` — eles não comunicam propósito
- **O pacote `main` é executável** — não deve ser importado por outros pacotes

---

*Próximo capítulo: **Capítulo 11 – Bug-Busting Debugging Skills** – técnicas de depuração para encontrar e corrigir erros de forma eficiente.*
