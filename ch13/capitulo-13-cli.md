# Capítulo 13 — Programando a partir da Linha de Comando (CLI)

> **Livro:** Go Programming: From Beginner to Professional — Samantha Coyle
> **Requisitos técnicos:** Go 1.21 ou superior
> **Código-fonte:** [GitHub — Chapter13](https://github.com/PacktPublishing/Go-Programming-From-Beginner-to-Professional-Second-Edition-/tree/main/Chapter13)

---

## Visão Geral

Este capítulo mostra por que Go é uma escolha excelente para criar utilitários e aplicações de linha de comando (CLI). Você aprende desde o básico — capturar argumentos e flags — até técnicas avançadas como streaming de grandes volumes de dados, tratamento de sinais do SO, execução de subprocessos e construção de interfaces gráficas no terminal (TUI).

Os grandes blocos do capítulo são:

1. Leitura de argumentos via `os.Args`
2. Controle de comportamento com o pacote `flag`
3. Streaming de dados com `stdin`, `stdout` e pipes
4. Códigos de saída e boas práticas de CLI
5. Tratamento gracioso de interrupções (`os/signal`)
6. Execução de comandos externos (`os/exec`)
7. Terminal User Interfaces (TUI) com Bubble Tea
8. Instalação de CLIs com `go install` e o pacote Cobra

---

## 1. Introdução: CLI como Interface de Aplicação

Interfaces de usuário não precisam ser páginas web. O terminal é uma interface legítima e muito poderosa — especialmente em automação, scripts, pipelines e ferramentas de DevOps. Go oferece um ecossistema rico de pacotes para construir CLIs robustas e performáticas.

---

## 2. Lendo Argumentos com `os.Args`

### Conceito

O pacote `os` expõe a slice `os.Args`, que contém todos os argumentos passados ao programa na invocação. O índice `0` sempre é o nome do executável; os argumentos do usuário começam no índice `1`.

### Por que isso importa

Argumentos de linha de comando permitem que o usuário personalize o comportamento do programa **sem alterar o código-fonte**. Isso é fundamental para automação e scripting.

### Exercício 13.01 — Saudar o usuário pelo nome

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    args := os.Args

    if len(args) < 2 {
        fmt.Println("Usage: go run main.go <name>")
        return
    }

    name := args[1]
    greeting := fmt.Sprintf("Hello, %s! Welcome to the command line.", name)
    fmt.Println(greeting)
}
```

**Execução:**
```bash
go run main.go Sam
# Hello, Sam! Welcome to the command line.
```

**Pontos-chave:**
- `os.Args[0]` = nome do executável (descartado na validação `len(args) < 2`)
- `os.Args[1]` = primeiro argumento real fornecido pelo usuário
- Validação de quantidade de argumentos evita panics por index out of range

---

## 3. Usando Flags para Controlar o Comportamento

### Conceito

O pacote `flag` oferece uma abordagem mais estruturada do que `os.Args` diretamente. Ele permite definir flags com tipos, valores padrão e descrições, além de gerar automaticamente uma mensagem de ajuda (`--help`).

### Fluxo de uso do pacote `flag`

1. **Definir flags** — com tipo e valor padrão
2. **Parsear as flags** — com `flag.Parse()`
3. **Acessar os valores** — via ponteiro ou variável associada

### Exercício 13.02 — Saudar condicionalmente com flags

```go
package main

import (
    "flag"
    "fmt"
)

var (
    nameFlag  = flag.String("name", "Sam", "Name of the person to say hello to")
    quietFlag = flag.Bool("quiet", false, "Toggle to be quiet when saying hello")
)

func main() {
    flag.Parse()

    if !*quietFlag {
        greeting := fmt.Sprintf("Hello, %s! Welcome to the command line.", *nameFlag)
        fmt.Println(greeting)
    }
}
```

**Formas de execução:**

```bash
go run main.go
# Hello, Sam! Welcome to the command line.

go run main.go --name=Cassie --quiet=false
# Hello, Cassie! Welcome to the command line.

go run main.go --quiet=true
# (sem saída — silencioso)

go run main.go --help
# Usage of ...:
#   -name string
#       Name of the person to say hello to (default "Sam")
#   -quiet
#       Toggle to be quiet when saying hello
```

**Vantagens das flags:**
- Documentação automática via `--help`
- Valores padrão explícitos no código
- Código mais legível e autodocumentado
- Suporte a múltiplos tipos: `String`, `Bool`, `Int`, `Float64`, etc.

---

## 4. Streaming de Grandes Volumes de Dados

### Conceito

Aplicações CLI frequentemente fazem parte de um pipeline maior — recebendo dados de outro programa via `stdin` e enviando resultados via `stdout`. Processar dados em *streaming* (linha a linha) é muito mais eficiente do que carregar tudo na memória de uma vez.

### Benefícios do streaming em Go

- **Eficiência de memória:** dados processados linha a linha, sem carregar o arquivo inteiro
- **Análise em tempo real:** o usuário vê resultados conforme são produzidos
- **Interface interativa:** o programa pode exibir detalhes dinâmicos durante o processamento

### Rot13: o exemplo prático

**Rot13** é uma cifra de substituição simples: cada letra é substituída pela letra 13 posições à frente no alfabeto (A → N, B → O, etc.). É simétrico — aplicar duas vezes retorna o original. Usado aqui como exemplo lúdico de encoding.

### Exercício 13.03 — Rot13 via pipes, stdin e stdout

```go
package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
)

// Função de encoding Rot13
func rot13(s string) string {
    result := make([]byte, len(s))
    for i := 0; i < len(s); i++ {
        char := s[i]
        switch {
        case char >= 'a' && char <= 'z':
            result[i] = 'a' + (char-'a'+13)%26
        case char >= 'A' && char <= 'Z':
            result[i] = 'A' + (char-'A'+13)%26
        default:
            result[i] = char
        }
    }
    return string(result)
}

// Lê do stdin linha a linha e aplica Rot13
func processStdin() {
    reader := bufio.NewReader(os.Stdin)
    for {
        input, err := reader.ReadString('\n')
        if err == io.EOF {
            break
        } else if err != nil {
            fmt.Println("Error reading stdin:", err)
            return
        }
        fmt.Print(rot13(input))
    }
}

// Processa arquivo ou entrada do usuário
func processFileOrInput() {
    var inputReader io.Reader

    if len(os.Args) > 1 {
        file, err := os.Open(os.Args[1])
        if err != nil {
            fmt.Println("Error opening file:", err)
            return
        }
        defer file.Close()
        inputReader = file
    } else {
        fmt.Print("Enter text: ")
        inputReader = os.Stdin
    }

    scanner := bufio.NewScanner(inputReader)
    for scanner.Scan() {
        fmt.Println(rot13(scanner.Text()))
    }
    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading input:", err)
    }
}

func main() {
    stat, _ := os.Stdin.Stat()
    if (stat.Mode() & os.ModeCharDevice) == 0 {
        processStdin() // dados disponíveis via pipe
    } else {
        processFileOrInput() // leitura interativa ou de arquivo
    }
}
```

**Formas de uso:**

```bash
# Entrada interativa
go run main.go
# Enter text: enjoy the book
# rawbl gur obbx

# Via pipe (cat no Linux, type no Windows)
cat data.txt | go run main.go

# Passando arquivo como argumento
go run main.go data.txt
```

**Conceitos demonstrados:**
- `bufio.NewReader` para leitura eficiente linha a linha
- `os.Stdin.Stat()` para detectar se há dados sendo piped
- Interface `io.Reader` para abstrair a fonte de dados (arquivo ou stdin)
- `bufio.NewScanner` como alternativa ao `NewReader`

---

## 5. Códigos de Saída e Boas Práticas de CLI

### Conceito

Códigos de saída permitem que o programa comunique seu status ao ambiente que o chamou (shell, scripts, CI/CD). O padrão Unix é: `0` = sucesso, qualquer valor diferente de zero = erro.

```go
const (
    ExitCodeSuccess      = 0
    ExitCodeInvalidInput = 1
    ExitCodeFileNotFound = 2
)

// Uso:
os.Exit(ExitCodeSuccess)
os.Exit(ExitCodeFileNotFound)
```

Para verificar o código de saída no shell:
```bash
echo $?   # Imprime o código de saída do último comando
```

### Boas práticas de CLI listadas no capítulo

- **Logging consistente:** use mensagens significativas para facilitar o debug
- **Informações de uso claras:** documente flags, argumentos e exemplos de uso
- **Suporte a `--help` e `--version`:** torna a ferramenta mais amigável
- **Terminação graciosa:** execute limpeza (fechar arquivos, liberar recursos) antes de sair

---

## 6. Tratamento de Interrupções com `os/signal`

### Conceito

Programas robustos precisam responder adequadamente a sinais do sistema operacional — como quando o usuário pressiona `Ctrl+C` (SIGINT) ou `Ctrl+Z` (SIGTSTP). Em vez de terminar abruptamente, a aplicação deve:

- Finalizar transações em andamento
- Fechar conexões e arquivos
- Salvar estado se necessário
- Liberar recursos alocados

Isso é chamado de **graceful shutdown** (terminação graciosa).

### Funcionamento

O pacote `os/signal` fornece a função `signal.Notify`:

```go
func Notify(c chan<- os.Signal, sig ...os.Signal)
```

Ela registra um canal para receber os sinais especificados.

### Exemplo de captura de SIGINT

```go
package main

import (
    "fmt"
    "os"
    "os/signal"
    "syscall"
)

func main() {
    sigs := make(chan os.Signal, 1)
    done := make(chan struct{})

    signal.Notify(sigs, syscall.SIGINT)

    go func() {
        for {
            s := <-sigs
            switch s {
            case syscall.SIGINT:
                fmt.Println("\nInterrupção detectada (Ctrl+C)")
                fmt.Println("Realizando limpeza...")
                done <- struct{}{}
            }
        }
    }()

    fmt.Println("Aguardando sinal...")
    <-done
    fmt.Println("Encerrando com segurança.")
}
```

**Sinais comuns:**

| Sinal | Tecla | Descrição |
|-------|-------|-----------|
| `SIGINT` | Ctrl+C | Interrupção do usuário |
| `SIGTSTP` | Ctrl+Z | Suspensão do processo |
| `SIGTERM` | — | Pedido de terminação (ex: `kill`) |

**Padrão de implementação:**
1. Criar canal de sinais com buffer
2. Criar canal `done` para bloquear o programa
3. Registrar sinais com `signal.Notify`
4. Goroutine fica em loop aguardando sinais
5. Ao receber sinal, executa limpeza e sinaliza `done`
6. Main bloqueia em `<-done` até receber a confirmação

---

## 7. Executando Comandos Externos com `os/exec`

### Conceito

O pacote `os/exec` permite que sua aplicação Go inicie e interaja com processos externos. Isso abre possibilidades como:

- Executar ferramentas do sistema
- Capturar saída de outros programas
- Comunicação bidirecional com subprocessos
- Execução em paralelo/background

**Consideração cross-platform:** comportamentos de shell diferem entre SO. O `os/exec` oferece uma solução portável.

### Exercício 13.04 — Stopwatch com execução de comando externo

```go
package main

import (
    "fmt"
    "os"
    "os/exec"
    "time"
)

func main() {
    timeLimit := 5 * time.Second

    fmt.Println("Pressione Enter para iniciar o cronômetro...")
    _, err := fmt.Scanln()
    if err != nil {
        fmt.Println("Erro ao ler stdin:", err)
        return
    }

    fmt.Println("Cronômetro iniciado. Aguardando", timeLimit)
    time.Sleep(timeLimit)

    fmt.Println("Tempo esgotado! Executando comando externo.")
    cmd := exec.Command("echo", "Hello")
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    if err = cmd.Run(); err != nil {
        fmt.Println("Erro ao executar comando:", err)
    }
}
```

**Saída:**
```
Pressione Enter para iniciar o cronômetro...
Cronômetro iniciado. Aguardando 5s
Tempo esgotado! Executando comando externo.
Hello
```

**Destaques da API `os/exec`:**
- `exec.Command(name, args...)` — define o comando
- `cmd.Stdout` / `cmd.Stderr` — conecta saída ao processo pai
- `cmd.Run()` — executa e aguarda a conclusão
- `cmd.Start()` + `cmd.Wait()` — para execução assíncrona

---

## 8. Terminal User Interfaces (TUI)

### Conceito

TUIs permitem criar interfaces gráficas interativas dentro do terminal. Go tem uma biblioteca popular para isso: **Bubble Tea** (`github.com/charmbracelet/bubbletea`).

Uma TUI com Bubble Tea segue a **arquitetura Model-View-Update (MVU)**, parecida com Elm:

- **Model:** estrutura de dados que representa o estado da interface
- **Update:** função que processa mensagens (eventos) e retorna o novo estado
- **View:** função que renderiza o estado atual como string

### Interface `tea.Model`

```go
type model struct {
    cursor int
    choice string
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) { ... }

func (m model) View() string { ... }
```

### Exercício 13.05 — TUI para o pipeline Rot13

Este exercício envolve criar uma interface de menu para o programa Rot13 do Exercício 13.03:

```go
package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
    tea "github.com/charmbracelet/bubbletea"
)

var choices = []string{"File input", "Type in input"}

type model struct {
    cursor int
    choice string
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "ctrl+c", "q", "esc":
            return m, tea.Quit
        case "enter":
            m.choice = choices[m.cursor]
            return m, tea.Quit
        case "down", "j":
            m.cursor++
            if m.cursor >= len(choices) {
                m.cursor = 0
            }
        case "up", "k":
            m.cursor--
            if m.cursor < 0 {
                m.cursor = len(choices) - 1
            }
        }
    }
    return m, nil
}

func (m model) View() string {
    s := strings.Builder{}
    s.WriteString("Selecione o modo de entrada:\n\n")
    for i, choice := range choices {
        if m.cursor == i {
            s.WriteString("(•) ")
        } else {
            s.WriteString("( ) ")
        }
        s.WriteString(choice)
        s.WriteString("\n")
    }
    s.WriteString("\n(q para sair)\n")
    return s.String()
}

func main() {
    p := tea.NewProgram(model{})
    m, err := p.Run()
    if err != nil {
        fmt.Println("Erro:", err)
        os.Exit(1)
    }

    if m, ok := m.(model); ok && m.choice != "" {
        fmt.Printf("\n---\nVocê escolheu: %s!\n", m.choice)
        if m.choice == "File input" {
            processFile("data.txt")
        } else {
            processStdin()
        }
    }
}
```

**Saída ao selecionar "File input":**
```
Selecione o modo de entrada:

(•) File input
( ) Type in input

(q para sair)
---
Você escolheu: File input!
rawbl
gur
obbx
```

**Componentes de uma TUI:**
- **Componentes:** botões, campos de input, listas
- **Layouts:** organização visual dos componentes
- **Tratamento de input:** resposta a teclas, mouse, etc.

**Teclas suportadas no exemplo:**
- `↑` / `k` — mover cursor para cima
- `↓` / `j` — mover cursor para baixo
- `Enter` — confirmar seleção
- `q` / `Esc` / `Ctrl+C` — sair

---

## 9. Instalando CLIs com `go install`

### Conceito

O comando `go install` compila e instala aplicações Go no diretório `bin` do workspace (`$GOPATH/bin`), tornando-as disponíveis globalmente no terminal.

```bash
go install ./...                            # Instala o projeto atual
go install github.com/autor/pacote@latest   # Instala da internet
```

### Flags de compilação cross-platform

```bash
GOOS=linux   GOARCH=amd64 go build -o app-linux
GOOS=windows GOARCH=amd64 go build -o app.exe
GOOS=darwin  GOARCH=arm64 go build -o app-mac
```

- `GOOS` — sistema operacional alvo (`linux`, `windows`, `darwin`)
- `GOARCH` — arquitetura alvo (`amd64`, `arm64`, `386`)

### Exemplo: instalando o Cobra CLI

**Cobra** é um dos pacotes mais populares do ecossistema Go para construir CLIs profissionais. Ele é usado por projetos como Docker, Kubernetes e GitHub CLI.

```bash
go install github.com/spf13/cobra-cli@latest
```

Após a instalação:
```bash
cobra-cli --help
# Cobra is a CLI library for Go that empowers applications.
#
# Usage:
#   cobra-cli [command]
#
# Available Commands:
#   add         Add a command to a Cobra Application
#   completion  Generate the autocompletion script for the specified shell
#   init        Initialize a Cobra Application
#
# Flags:
#   -a, --author string   author name for copyright attribution
#   -h, --help            help for cobra-cli
#   -l, --license string  name of license for the project
#       --viper           use Viper for configuration
```

O Cobra também oferece scaffolding automático de projetos CLI, acelerando muito o desenvolvimento.

---

## 10. Resumo Geral do Capítulo

| Tópico | Pacote/Ferramenta | O que foi aprendido |
|--------|-------------------|---------------------|
| Argumentos | `os` (`os.Args`) | Capturar parâmetros brutos da linha de comando |
| Flags | `flag` | Definir, parsear e acessar flags tipadas com valores padrão |
| Streaming | `bufio`, `io`, `os` | Processar stdin/stdout e pipes eficientemente |
| Encoding | — (Rot13) | Cifra de substituição simétrica, exemplo de transformação de dados |
| Exit codes | `os` (`os.Exit`) | Comunicar sucesso/falha ao ambiente chamador |
| Boas práticas | — | Logging, help, version, terminação graciosa |
| Interrupções | `os/signal`, `syscall` | Capturar SIGINT/SIGTSTP e executar cleanup |
| Subprocessos | `os/exec` | Iniciar e interagir com processos externos |
| TUI | `bubbletea` | Criar interfaces interativas no terminal (MVU) |
| Instalação | `go install` | Distribuir e instalar CLIs Go globalmente |
| CLI framework | `cobra` | Scaffolding e estruturação de CLIs complexas |

### Fluxo do capítulo

```
Argumentos simples (os.Args)
        ↓
Flags estruturadas (flag package)
        ↓
Streaming de dados (stdin/stdout/pipes)
        ↓
Exit codes e boas práticas
        ↓
Tratamento de interrupções (signals)
        ↓
Execução de comandos externos (os/exec)
        ↓
Interfaces TUI (Bubble Tea)
        ↓
Distribuição com go install + Cobra
```

---

## Conexões com Outros Capítulos

- **Capítulo 12 (Tempo):** Usado no Exercício 13.04 — `time.Sleep` e `time.Second`
- **Capítulo 14 (Arquivos e Sistemas):** Expansão natural — leitura/escrita de arquivos a partir da CLI; tratamento de sinais em contexto de filesystem
- **Capítulos 15-16 (SQL e Web):** CLIs são muitas vezes a interface de entrada para ferramentas que interagem com bancos de dados e APIs
