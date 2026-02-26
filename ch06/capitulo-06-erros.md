# Capítulo 6 — Tratamento de Erros

> **Revisão rápida:** O sistema de erros de Go é deliberadamente simples e explícito. Não há exceções — erros são valores como qualquer outro. Isso parece verboso no início, mas produz código mais previsível e fácil de depurar. `panic` e `recover` existem para situações verdadeiramente excepcionais.

---

## 1. O que são Erros?

```go
package main

import (
    "errors"
    "fmt"
)

// Em Go, um erro é qualquer valor que implementa a interface error:
// type error interface {
//     Error() string
// }

func abrirArquivo(caminho string) error {
    if caminho == "" {
        return errors.New("caminho do arquivo não pode ser vazio")
    }
    if caminho == "/etc/shadow" {
        return fmt.Errorf("acesso negado ao arquivo: %s", caminho)
    }
    return nil
}

func main() {
    // Erros são retornados, não lançados
    if err := abrirArquivo(""); err != nil {
        fmt.Println("Erro:", err)
    }

    if err := abrirArquivo("/etc/shadow"); err != nil {
        fmt.Println("Erro:", err)
    }

    if err := abrirArquivo("/home/lucas/config.json"); err != nil {
        fmt.Println("Erro:", err)
    } else {
        fmt.Println("Arquivo aberto com sucesso!")
    }
}
```

---

## 2. A Interface error

```go
package main

import (
    "fmt"
    "time"
)

// Criando tipos de erro customizados
type ErroValidacao struct {
    Campo   string
    Valor   interface{}
    Motivo  string
}

func (e *ErroValidacao) Error() string {
    return fmt.Sprintf("validação falhou: campo '%s' com valor '%v' — %s",
        e.Campo, e.Valor, e.Motivo)
}

type ErroTimeout struct {
    Operacao  string
    Duracao   time.Duration
}

func (e *ErroTimeout) Error() string {
    return fmt.Sprintf("timeout após %v na operação: %s", e.Duracao, e.Operacao)
}

func validarEmail(email string) error {
    if len(email) == 0 {
        return &ErroValidacao{Campo: "email", Valor: email, Motivo: "não pode ser vazio"}
    }
    contemArroba := false
    for _, ch := range email {
        if ch == '@' {
            contemArroba = true
            break
        }
    }
    if !contemArroba {
        return &ErroValidacao{Campo: "email", Valor: email, Motivo: "deve conter '@'"}
    }
    return nil
}

func consultarAPI(endpoint string) error {
    // Simulando timeout
    if endpoint == "/dados-lentos" {
        return &ErroTimeout{Operacao: "GET " + endpoint, Duracao: 30 * time.Second}
    }
    return nil
}

func main() {
    erros := []string{"", "emailsemarroba.com", "lucas@exemplo.com"}
    for _, e := range erros {
        if err := validarEmail(e); err != nil {
            fmt.Println("❌", err)
        } else {
            fmt.Printf("✅ Email '%s' válido\n", e)
        }
    }

    if err := consultarAPI("/dados-lentos"); err != nil {
        fmt.Println("\n❌", err)
    }
}
```

---

## 3. Panic

```go
package main

import "fmt"

func acessarIndice(slice []int, i int) int {
    // Panic ocorre automaticamente ao acessar índice inválido
    return slice[i]
}

func dividirInteiro(a, b int) int {
    // Panic por divisão por zero com inteiros
    return a / b
}

// Casos onde VOCÊ deve usar panic:
// - Condição que nunca deveria ocorrer
// - Erro de programação (bug), não erro de usuário
func configurarServidor(porta int) {
    if porta < 1 || porta > 65535 {
        // Configuração inválida = bug do programador, não erro do usuário
        panic(fmt.Sprintf("porta inválida: %d (deve ser 1-65535)", porta))
    }
    fmt.Printf("Servidor configurado na porta %d\n", porta)
}

func main() {
    configurarServidor(8080) // OK

    // As linhas abaixo causariam panic:
    // configurarServidor(99999)
    // acessarIndice([]int{1, 2, 3}, 10)
    // dividirInteiro(10, 0)

    fmt.Println("Programa funcionando normalmente")
}
```

> **💡 Insight:** A regra geral é clara — use `panic` para **bugs de programação** (situações que nunca deveriam acontecer se o código estiver correto). Para **erros esperáveis** em tempo de execução (arquivo não encontrado, input inválido, timeout), retorne um `error`.

---

## 4. Recover

```go
package main

import "fmt"

// recover() só funciona dentro de um defer
func recuperarDePanic() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Printf("Panic recuperado: %v\n", r)
        }
    }()

    panic("algo deu muito errado!")
}

// Padrão útil: wrapper que converte panic em error
func executarSeguro(fn func()) (err error) {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("panic capturado: %v", r)
        }
    }()
    fn()
    return nil
}

func operacaoArriscada(divisor int) {
    resultado := 100 / divisor // panic se divisor == 0
    fmt.Println("Resultado:", resultado)
}

func main() {
    recuperarDePanic()
    fmt.Println("Continuando após recover...")

    // Executando código arriscado de forma segura
    err := executarSeguro(func() {
        operacaoArriscada(5)
    })
    if err != nil {
        fmt.Println("Erro:", err)
    }

    err = executarSeguro(func() {
        operacaoArriscada(0) // vai causar panic
    })
    if err != nil {
        fmt.Println("Erro capturado:", err)
    }

    fmt.Println("Programa sobreviveu!")
}
```

---

## 5. Error Wrapping

```go
package main

import (
    "errors"
    "fmt"
)

// Erros sentinela — valores de erro conhecidos
var (
    ErrNaoEncontrado = errors.New("registro não encontrado")
    ErrAcessoNegado  = errors.New("acesso negado")
    ErrConexao       = errors.New("falha na conexão")
)

type BancoDeDados struct {
    conectado bool
}

func (db *BancoDeDados) buscarUsuario(id int) (string, error) {
    if !db.conectado {
        return "", fmt.Errorf("buscarUsuario(%d): %w", id, ErrConexao)
    }
    if id == 99 {
        return "", fmt.Errorf("buscarUsuario(%d): %w", id, ErrNaoEncontrado)
    }
    if id < 0 {
        return "", fmt.Errorf("buscarUsuario(%d): id negativo: %w", id, ErrAcessoNegado)
    }
    return fmt.Sprintf("Usuário#%d", id), nil
}

func processarRequisicao(db *BancoDeDados, userID int) error {
    nome, err := db.buscarUsuario(userID)
    if err != nil {
        // Adiciona contexto sem perder o erro original
        return fmt.Errorf("processarRequisicao: %w", err)
    }
    fmt.Println("Processando para:", nome)
    return nil
}

func main() {
    db := &BancoDeDados{conectado: true}

    casos := []int{1, 99, -5}
    for _, id := range casos {
        err := processarRequisicao(db, id)
        if err != nil {
            // errors.Is — verifica se o erro (ou algum wrappado) é do tipo esperado
            switch {
            case errors.Is(err, ErrNaoEncontrado):
                fmt.Printf("ID %d: não existe no banco\n", id)
            case errors.Is(err, ErrAcessoNegado):
                fmt.Printf("ID %d: sem permissão\n", id)
            case errors.Is(err, ErrConexao):
                fmt.Printf("ID %d: problema de conexão\n", id)
            default:
                fmt.Printf("ID %d: erro inesperado: %v\n", id, err)
            }
            // Mostra a cadeia completa de erros
            fmt.Printf("  Cadeia completa: %v\n", err)
        }
    }

    // errors.As — extrai o tipo concreto do erro
    type ErroDetalhe struct {
        Codigo  int
        Mensagem string
    }
    func (e *ErroDetalhe) Error() string {
        return fmt.Sprintf("erro %d: %s", e.Codigo, e.Mensagem)
    }

    errOriginal := fmt.Errorf("operação falhou: %w", &ErroDetalhe{404, "não encontrado"})
    var detalhe *ErroDetalhe
    if errors.As(errOriginal, &detalhe) {
        fmt.Printf("\nCódigo de erro: %d\n", detalhe.Codigo)
    }
}
```

> **💡 Insight:** O `%w` em `fmt.Errorf` é o operador de wrapping. Ele preserva o erro original dentro do novo erro. Use `errors.Is` para verificar **qual** erro é e `errors.As` para **extrair** o tipo concreto. Sempre adicione contexto ao fazer wrapping — onde no código o erro ocorreu e com quais parâmetros.

---

## 🔁 Revisão Rápida — O que lembrar

| Conceito | Quando usar |
|---|---|
| `errors.New` | Erro simples e estático |
| `fmt.Errorf` | Erro com informação dinâmica |
| Tipo de erro customizado | Quando precisa de campos extras ou `errors.As` |
| `%w` (wrapping) | Para adicionar contexto sem perder o erro original |
| `errors.Is` | Verifica o tipo de erro (inclusive em cadeias) |
| `errors.As` | Extrai o tipo concreto de erro |
| `panic` | Só para bugs — condições que nunca deveriam acontecer |
| `recover` | Só dentro de `defer` — converte panic em comportamento controlado |

**Fluxo ideal de erro em Go:**
```
função baixo nível → retorna erro básico
    ↓ fmt.Errorf("contexto: %w", err)
função intermediária → adiciona contexto
    ↓ fmt.Errorf("contexto: %w", err)
handler/main → usa errors.Is/As para decidir o que fazer
```
