# Exercício 09.03 — Go Workspaces (Passo a Passo)

## 🎯 Objetivo
Entender **por que** workspaces existem e **como** usá-los.

---

## Parte 1: Preparando os Módulos

### Passo 1 — Criar a estrutura de pastas

```bash
mkdir workspace-demo
cd workspace-demo
mkdir printer app
```

### Passo 2 — Criar o módulo `printer`

```bash
cd printer
go mod init github.com/lucas/printer
```

Criar o arquivo `printer.go`:

```go
// printer/printer.go
package printer

import (
    "fmt"
    "github.com/google/uuid"
)

func PrintNewUUID() string {
    id := uuid.New()
    return fmt.Sprintf("UUID Gerado: %s", id)
}
```

Instalar a dependência do uuid:

```bash
go mod tidy
cd ..
```

### Passo 3 — Criar o módulo `app`

```bash
cd app
go mod init app
```

Criar o arquivo `main.go`:

```go
// app/main.go
package main

import (
    "fmt"
    "github.com/lucas/printer"
)

func main() {
    msg := printer.PrintNewUUID()
    fmt.Println(msg)
}
```

---

## Parte 2: Vendo o Problema

### Passo 4 — Tentar rodar (VAI FALHAR!)

```bash
# Dentro da pasta app/
go mod tidy
```

❌ **Erro esperado:**
```
go: finding module for package github.com/lucas/printer
cannot find module providing package github.com/lucas/printer
```

**Por quê?** O Go tenta baixar `github.com/lucas/printer` da internet,
mas esse módulo só existe LOCALMENTE na pasta ao lado.

---

## Parte 3: Solução Antiga (replace)

### Passo 5 — Editar go.mod com replace

```bash
# Dentro da pasta app/
go mod edit -replace github.com/lucas/printer=../printer
```

Veja como ficou o `app/go.mod`:

```
module app

go 1.21

replace github.com/lucas/printer => ../printer
```

### Passo 6 — Agora sim, funciona!

```bash
go mod tidy
go run main.go
```

✅ **Saída esperada:**
```
UUID Gerado: 7a533339-58b6-4396-b7f7-d0a50216bf88
```

### ⚠️ Problema com essa abordagem:
- Você "sujou" o go.mod com um caminho relativo (`../printer`)
- Se publicar esse go.mod, vai quebrar para outras pessoas
- Com muitos módulos, fica difícil gerenciar tantos `replace`

---

## Parte 4: Solução Nova — Go Workspaces! 🚀

### Passo 7 — Limpar o go.mod do app

Substitua o conteúdo inteiro de `app/go.mod` por:

```
module app

go 1.21
```

(Removemos o `replace` e o `require`)

### Passo 8 — Criar o workspace

```bash
# Volte para a pasta raiz workspace-demo/
cd ..

# Inicializar o workspace
go work init

# Adicionar os dois módulos ao workspace
go work use ./printer
go work use ./app
```

### Passo 9 — Ver o arquivo go.work gerado

```bash
cat go.work
```

```
go 1.21

use (
    ./printer
    ./app
)
```

### Passo 10 — Rodar o app pelo workspace

```bash
# Da pasta raiz workspace-demo/
go run ./app
```

✅ **Saída esperada:**
```
UUID Gerado: 5ff596a2-7c0e-41fe-b0b1-256b28a35b76
```

🎉 **Funcionou SEM nenhum `replace` no go.mod!**

---

## Estrutura Final

```
workspace-demo/
├── go.work          ← Arquivo de workspace (NÃO vai pro Git do app)
├── printer/
│   ├── printer.go
│   ├── go.mod       ← module github.com/lucas/printer
│   └── go.sum
└── app/
    ├── main.go
    └── go.mod       ← module app (LIMPO, sem replace!)
```

---

## Comparação Visual

```
╔══════════════════════════════════════════════════════════════╗
║              ANTES (Go < 1.18) — Com replace                ║
╠══════════════════════════════════════════════════════════════╣
║                                                              ║
║  app/go.mod:                                                 ║
║  ┌──────────────────────────────────────────────┐            ║
║  │ module app                                    │            ║
║  │ replace github.com/lucas/printer => ../printer│  ← SUJO!  ║
║  │ require github.com/lucas/printer v0.0.0       │            ║
║  └──────────────────────────────────────────────┘            ║
║                                                              ║
║  ⚠️ Cada go.mod precisa de replace para cada módulo local    ║
║  ⚠️ Precisa remover antes de publicar                        ║
╚══════════════════════════════════════════════════════════════╝

╔══════════════════════════════════════════════════════════════╗
║              DEPOIS (Go >= 1.18) — Com workspace            ║
╠══════════════════════════════════════════════════════════════╣
║                                                              ║
║  go.work:                          app/go.mod:               ║
║  ┌─────────────────────┐          ┌──────────────┐           ║
║  │ use (               │          │ module app   │ ← LIMPO!  ║
║  │     ./printer       │          │ go 1.21      │           ║
║  │     ./app           │          └──────────────┘           ║
║  │ )                   │                                     ║
║  └─────────────────────┘                                     ║
║                                                              ║
║  ✅ Um único arquivo gerencia tudo                           ║
║  ✅ go.mod dos módulos ficam limpos                          ║
║  ✅ go.work fica fora do controle de versão                  ║
╚══════════════════════════════════════════════════════════════╝
```

---

## Resumo dos Comandos de Workspace

| Comando               | O que faz                                        |
|-----------------------|--------------------------------------------------|
| `go work init`        | Cria o arquivo `go.work` na pasta atual          |
| `go work use ./pasta` | Adiciona um módulo local ao workspace            |
| `go work sync`        | Sincroniza as dependências do workspace          |
| `go work edit`        | Edita o go.work programaticamente                |

---

## 💡 Regra de Ouro

> **`go.work` é para desenvolvimento LOCAL.**
> Ele **não deve** ser commitado no repositório Git.
> Quando o módulo `printer` for publicado de verdade no GitHub,
> o `app` vai baixá-lo normalmente, sem precisar de workspace.
