Solução 2: O Jeito Novo — Workspaces (go.work)
Em vez de sujar cada go.mod com replace, você cria um arquivo acima de tudo que diz "estes módulos locais se conhecem":

```Go
workspace-demo/
├── go.work         ← ✨ NOVO! Arquivo de workspace
├── printer/
│   ├── printer.go
│   └── go.mod
└── app/
    ├── main.go
    └── go.mod      ← LIMPO, sem replace!
```

O `go.work` contém:

```Go
go 1.21

use (
    ./printer
    ./app
)
```

Agora o Go sabe: "quando `app` pedir por `github.com/lucas/printer`, olhe primeiro na pasta `./printer` local". Sem tocar em nenhum `go.mod`!
