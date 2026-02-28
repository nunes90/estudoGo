# Capítulo 9 — Usando Go Modules para Definir um Projeto

## Visão Geral

O Capítulo 9 aborda o sistema de **módulos do Go** — o mecanismo oficial para estruturar projetos, gerenciar dependências e controlar versionamento. Módulos foram introduzidos oficialmente no **Go 1.11** e substituem as práticas anteriores baseadas em `GOPATH` e diretórios `vendor`.

---

## O que é um Módulo?

Um módulo Go é uma **coleção de pacotes** agrupados sob um caminho versionado comum. Ele funciona como uma unidade autocontida e encapsulada que:

- Organiza e estrutura o código do projeto
- Gerencia dependências externas com controle de versão
- Promove reutilização e manutenibilidade do código

---

## Componentes-Chave dos Módulos

### 1. Arquivo `go.mod`

O arquivo de configuração principal do módulo. Contém:

- **Module path** — o caminho/identificador do módulo (ex: `module mymodule`)
- **Dependencies** (`require`) — lista de dependências com versões específicas
- **Replace directives** (opcional) — substituições de dependências para uso local ou testes
- **Exclude directives** (opcional) — exclusão de versões problemáticas

Exemplo:

```go
module mymodule

require (
    github.com/some/dependency v1.2.3
    github.com/another/dependency v2.0.0
)

replace (
    github.com/dependency/v3 => github.com/dependency/v4
)

exclude (
    github.com/some/dependency v2.0.0
)
```

### 2. Arquivo `go.sum`

Gerado e mantido automaticamente pelo Go toolchain. Contém **checksums criptográficos** (SHA-256) de todas as dependências para garantir:

- Integridade dos pacotes baixados
- Proteção contra adulteração ou corrupção
- Reprodutibilidade do ambiente de desenvolvimento

### 3. Versionamento Semântico

Cada módulo recebe um identificador de versão único (tags ou commit hashes), seguindo o modelo de **Semantic Versioning** ([semver.org](https://semver.org)):

- `MAJOR.MINOR.PATCH` (ex: `v1.2.3`)

---

## Benefícios dos Módulos

| Benefício | Descrição |
|-----------|-----------|
| **Gerenciamento preciso de dependências** | Controle exato das versões necessárias, sem conflitos |
| **Versionamento e reprodutibilidade** | Garante o mesmo ambiente de desenvolvimento para todos |
| **Colaboração melhorada** | Limites claros do código, facilitando contribuições |
| **Segurança de dependências** | Checksums no `go.sum` protegem contra pacotes adulterados |
| **Isolamento e modularidade** | Isola o projeto do workspace global, promovendo componentes reutilizáveis |

---

## Exercícios do Capítulo

### Exercício 09.01 — Criando seu primeiro módulo

Criação do módulo `bookutil` com um pacote `author` contendo operações sobre capítulos de livros:

```bash
mkdir bookutil && cd bookutil
go mod init bookutil
```

**Conceitos praticados:**
- Inicialização de módulo com `go mod init`
- Criação de pacotes internos ao módulo
- Structs exportadas, construtores (`NewAuthor`) e métodos
- Importação de pacotes internos: `import "bookutil/author"`

> **Nota:** O nome do módulo não precisa ser igual ao nome do pacote. Um módulo pode conter vários pacotes. Boas práticas de nomenclatura seguem padrões como `github.com/<projeto>/`.

### Exercício 09.02 — Usando um módulo externo

Uso do pacote `github.com/google/uuid` para gerar identificadores universalmente únicos:

```bash
go mod init myuuidapp
go get github.com/google/uuid
```

**Conceitos praticados:**
- Adição de dependências externas com `go get`
- Atualização automática do `go.mod` com a linha `require`
- Geração automática do `go.sum` com checksums
- Uso de pacotes de terceiros no código

### Atividade 9.01 — Consumindo múltiplos módulos

Combinação de dois módulos externos em um único projeto:
- `github.com/google/uuid` — para gerar UUIDs
- `rsc.io/quote` — para obter citações aleatórias

**Conceito-chave:** Projetos podem consumir quantas dependências externas forem necessárias, todas gerenciadas pelo sistema de módulos.

### Exercício 09.03 — Trabalhando com Workspaces

Demonstração do fluxo antes e depois dos **Go Workspaces** (Go 1.18+):

**Antes (com `replace`):**
```bash
go mod edit -replace github.com/sicoyle/printer=../printer
go mod tidy
```

**Depois (com workspaces):**
```bash
go work init
go work use ./printer
go run othermodule/main.go
```

---

## Múltiplos Módulos em um Projeto

É possível estruturar um projeto com **submódulos independentes**, cada um com seu próprio `go.mod`:

```
myproject/
├── mainmodule/
│   ├── main.go
│   ├── go.mod
│   └── go.sum
├── secondmodule/
│   ├── othermain.go
│   ├── go.mod
│   └── go.sum
└── thirdmodule/
    ├── othermain.go
    ├── go.mod
    └── go.sum
```

**Quando usar submódulos:**
- Componentes têm dependências diferentes
- Requisitos de versionamento distintos para o mesmo pacote
- Componentes precisam ser reutilizáveis em outros projetos
- Manutenção, testes e desenvolvimento separados fazem sentido

---

## Go Workspaces (`go.work`)

Introduzido no **Go 1.18**, workspaces simplificam o desenvolvimento local com múltiplos módulos:

- Elimina a necessidade de editar manualmente múltiplos `go.mod` com `replace`
- Define um arquivo `go.work` que aponta para os módulos locais
- Especialmente útil para projetos grandes ou que abrangem múltiplos repositórios

**Comandos principais:**
```bash
go work init          # Inicializa o workspace
go work use ./modulo  # Adiciona um módulo local ao workspace
```

---

## Comandos Go Essenciais para Módulos

| Comando | Função |
|---------|--------|
| `go mod init <nome>` | Inicializa um novo módulo |
| `go mod tidy` | Sincroniza dependências (adiciona faltantes, remove não usadas) |
| `go get <pacote>` | Adiciona ou atualiza uma dependência |
| `go mod edit -replace` | Adiciona diretiva de substituição |
| `go work init` | Inicializa um workspace |
| `go work use <dir>` | Adiciona módulo local ao workspace |

---

## Quando Usar Módulos Externos

Módulos de terceiros são recomendados para:

- Promover **reutilização** e eficiência de código
- Expandir a funcionalidade do projeto
- Delegar o gerenciamento de dependências
- Aproveitar **código open source** com confiabilidade e documentação comprovadas
- Desenvolvimento colaborativo com a comunidade

> **Cuidado:** Sempre avalie se a dependência externa é confiável, bem mantida e alinhada com os objetivos de longo prazo do projeto.

---

## Resumo dos Conceitos Principais

1. **Módulo** = coleção versionada de pacotes Go sob um caminho comum
2. **`go.mod`** = blueprint do módulo (caminho, versão Go, dependências, diretivas)
3. **`go.sum`** = checksums criptográficos para garantir integridade das dependências
4. **Versionamento semântico** = sistema `MAJOR.MINOR.PATCH` para controle de versões
5. **Workspaces** = facilidade para desenvolver com múltiplos módulos locais sem editar `go.mod`
6. **Submódulos** = módulos independentes dentro de um projeto maior, cada um com seu `go.mod`
