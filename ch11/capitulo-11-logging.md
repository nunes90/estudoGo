# Capítulo 11 — Bug-Busting Debugging Skills

> **Livro:** Go Programming – From Beginner to Professional (2ª ed.) — Samantha Coyle
> **Páginas:** 341–367

---

## Visão Geral

Este capítulo aborda metodologias práticas de debugging em Go. O objetivo é apresentar técnicas proativas para reduzir a introdução de bugs e, quando eles ocorrem, identificar sua localização com eficiência. Ao final, o leitor é capaz de formatar saídas de impressão, inspecionar tipos e valores de variáveis, utilizar o pacote `log` da biblioteca padrão e aplicar estratégias de depuração em ambientes restritos.

---

## 1. Introdução: O que são bugs e por que surgem?

**Debugging** é o processo de determinar a causa de um comportamento não intencional em um programa. A autora lista as principais causas de bugs:

| Causa | Explicação |
|---|---|
| **Testing como afterthought** | Testar somente ao final do desenvolvimento dificulta isolar onde o bug foi introduzido |
| **Melhorias de requisitos** | Mudanças em produção podem impactar áreas não relacionadas; testes unitários mitigam isso |
| **Prazos irrealistas** | Levam a atalhos nas boas práticas, menor design e menos testes |
| **Erros não tratados** | Ignorar retornos de erro permite que estados inválidos se propaguem silenciosamente |

**Relação de causa e efeito chave:** a ausência de testes unitários combinada com alterações em produção resulta em bugs difíceis de rastrear. A solução é integrar testes ao ciclo de desenvolvimento desde o início.

---

## 2. Métodos para Código Livre de Bugs

O capítulo apresenta cinco pilares preventivos (Figura 11.1):

```
┌───────────────────┐  ┌──────────────┐  ┌──────────────────┐
│  Code             │  │  Test Often  │  │  Write Unit      │
│  Incrementally    │  │              │  │  Tests           │
└───────────────────┘  └──────────────┘  └──────────────────┘
          ┌───────────────────────┐  ┌──────────────────┐
          │  Handle All Errors    │  │  Perform Logging │
          └───────────────────────┘  └──────────────────┘
```

### 2.1 Codificar Incrementalmente e Testar com Frequência
Desenvolver em pequenos incrementos e testar cada parte ao completá-la. Isso restringe o escopo de busca quando um bug aparece, pois há menos código novo entre dois estados conhecidamente corretos.

### 2.2 Escrever Testes Unitários
Um teste unitário recebe uma entrada conhecida e valida que a saída esperada é produzida. Se o teste passa antes de uma mudança e falha depois, conclui-se que a mudança introduziu o bug. Times de desenvolvimento exigem que testes passem antes de aceitar novos commits.

### 2.3 Tratar Todos os Erros
Discutido em detalhes no Capítulo 6. Ignorar valores de erro retornados por funções pode causar resultados inesperados e tornar o debugging muito mais difícil.

### 2.4 Realizar Logging
Logging permite determinar o estado do programa antes de uma falha. Os tipos comuns de log são: `debug`, `info`, `warn`, `error`, `fatal` e `trace`. O foco do capítulo é o **debug logging**.

> **Atenção ao desempenho:** em aplicações de alta carga, um volume excessivo de logs pode degradar a performance. Quanto mais usuários simultâneos, mais logs são gerados — e, em casos extremos, isso pode tornar a aplicação sem resposta.

### 2.5 Formatação com `fmt`
O pacote `fmt` é a porta de entrada para exibir dados no console ou em arquivos durante o debugging.

---

## 3. Formatação com o Pacote `fmt`

### 3.1 `fmt.Println()`
- Coloca espaços entre os argumentos.
- Adiciona automaticamente `\n` ao final.
- Imprime cada tipo no seu formato padrão (strings como estão, inteiros em decimal).

```go
fmt.Println("Hello:", fname, lname)
// Saída: Hello: Edward Scissorhands
```

### 3.2 `fmt.Printf()` — Verbs (Verbos de Formato)
Formata strings usando **verbos de formato** (inspirados na linguagem C). A variável substitui o verbo correspondente na ordem em que aparecem.

```go
fmt.Printf("Hello %s, good morning", fname)
// Saída: Hello Edward, good morning
```

**Tabela de verbos principais (Figura 11.3):**

| Verbo | Significado |
|---|---|
| `%d` | Inteiro em base 10 (decimal) |
| `%f` | Número de ponto flutuante, largura e precisão padrão |
| `%t` | Tipo `bool` |
| `%s` | String |
| `%v` | Valor no formato padrão |
| `%b` | Representação em base 2 (binário) |
| `%x` | Representação hexadecimal |
| `%T` | **Tipo** da variável (capital T) |
| `%#v` | **Representação Go** da variável (útil para structs, maps, slices) |

> **Diferença crítica:** `%t` (minúsculo) = booleano; `%T` (maiúsculo) = tipo da variável.

### 3.3 Opções Adicionais de Formatação

**Controle de precisão decimal:** `%.nf` arredonda o float para `n` casas decimais.

```
3.74567 + %.2f = 3.75
3.74567 + %.3f = 3.746
```

**Controle de largura total:** `%10.2f` define largura total de 10 caracteres com 2 casas decimais, alinhado à direita com padding de espaços.

**Alinhamento à esquerda:** use o flag `-` após `%`:

```go
fmt.Printf("%-10.2f\n", v)  // alinhado à esquerda
```

**Exercício 11.02** demonstra a impressão de valores de 1 a 255 em decimal, binário e hexadecimal com larguras fixas:
```go
fmt.Printf("Decimal: %3.d Base Two: %8.b Hex: %2.x\n", i, i, i)
```

---

## 4. Debugging Básico

Após dominar a formatação, o debugging básico envolve quatro técnicas (Figura 11.8):

### 4.1 Imprimir Marcadores no Código
`print` statements estrategicamente posicionados indicam **onde** o programa estava quando o bug ocorreu. O processo é iterativo: coloca-se um marcador, verifica-se se o código chega até ele e, caso não seja a origem do bug, move-se o marcador para outro ponto.

```go
fmt.Println("We are in function calculateGPA")
```

### 4.2 Imprimir o Tipo da Variável
Usando o verbo `%T` (maiúsculo):

```go
fmt.Printf("fname is of type %T\n", fname)
// Saída: fname is of type string
```

### 4.3 Imprimir o Valor da Variável
Usando `%v` para o valor padrão ou `%#v` para a representação Go completa (ideal para structs, maps e slices):

```go
fmt.Printf("fname value %#v\n", fname)
// Saída: fname value "Joe"

fmt.Printf("p value %#v\n", p)
// Saída: p value main.person{lname:"Lincoln", age:210, salary:25000}
```

O verbo `%#v` é especialmente valioso pois produz sintaxe que pode ser copiada diretamente no código Go.

### 4.4 Realizar Debug Logging
Quando se deseja logar para um arquivo (e não apenas para o terminal):

```go
log.Printf("fname value %#v\n", fname)
```

---

## 5. Logging com o Pacote `log`

### 5.1 Por que Logar?
Logging é uma infraestrutura do programa que captura eventos mesmo quando não há erros. É especialmente crítico em produção, onde o código pode se comportar diferente do ambiente de desenvolvimento (mais carga, dados malformados, etc.). Sem logging adequado, pode ser impossível reproduzir e entender um bug de produção.

### 5.2 Funções do Pacote `log`
O pacote `log` da biblioteca padrão de Go espelha as funções do `fmt`, mas adiciona automaticamente **data e hora** de execução:

```go
import "log"

log.Println("Demo app")           // → 2019/11/10 23:00:00 Demo app
log.Printf("%s is here!", name)   // → 2019/11/10 23:00:00 Thanos is here!
log.Print("Run")                  // → 2019/11/10 23:00:00 Run
```

### 5.3 Customizando Flags do Logger
A função `log.SetFlags()` permite incluir informações adicionais no prefixo de cada mensagem de log. Os flags podem ser combinados com o operador `|`:

| Flag | Descrição |
|---|---|
| `log.Ldate` | Data no fuso horário local (`2009/01/23`) |
| `log.Ltime` | Hora no fuso horário local (`01:23:23`) |
| `log.Lmicroseconds` | Resolução em microssegundos |
| `log.Llongfile` | Nome completo do arquivo e número da linha (`/a/b/c/d.go:23`) |
| `log.Lshortfile` | Nome do arquivo e número da linha (sobrescreve `Llongfile`) |
| `log.LUTC` | Usa UTC em vez do fuso local |
| `log.LstdFlags` | Valores iniciais padrão (`Ldate \| Ltime`) |

**Exemplo com múltiplos flags:**
```go
log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
log.Println("Demo app")
// Saída: 2019/04/30 08:15:57.835521 /go/src/myprojects/scratch/main.go:10: Demo app
```

**Causa e efeito:** ao ativar `Llongfile`, cada mensagem de log passa a exibir o caminho completo do arquivo e a linha, permitindo localizar exatamente a origem do log sem inspecionar o código manualmente.

---

## 6. Logging de Erros Fatais

### 6.1 Funções Fatal
As funções `Fatal()`, `Fatalf()` e `Fatalln()` funcionam como `Print`, `Printf` e `Println`, respectivamente, com uma diferença crítica: após logar, elas chamam `os.Exit(1)`.

**Consequência importante:** como `os.Exit(1)` não executa funções `defer`, o uso de `Fatal` deve ser reservado para situações onde o programa não pode continuar de forma segura — por exemplo, corrupção de dados iminente ou falha catastrófica de configuração.

```go
log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
log.Println("Start of our app")
err := errors.New("Application Aborted!")
if err != nil {
    log.Fatalln(err)  // loga e encerra com os.Exit(1)
}
log.Println("End of our app")  // esta linha NÃO é executada
```

### 6.2 Funções Panic vs Fatal

| Função | Comportamento |
|---|---|
| `log.Panic()`, `Panicf()`, `Panicln()` | Loga e então dispara um `panic` — **recuperável** via `defer` |
| `log.Fatal()`, `Fatalf()`, `Fatalln()` | Loga e chama `os.Exit(1)` — **não recuperável**, `defer` não executa |

**Quando usar Fatal:** quando o programa chegou a um estado onde continuar causaria corrupção de dados ou comportamento indesejado. Também útil em utilitários de linha de comando que precisam sinalizar falha ao chamador via exit code.

### 6.3 Atividade 11.01 — Validação de CPF/SSN com Logging
A atividade prática demonstra o uso do pacote `log` para rastrear o processo de validação de números de Seguro Social (SSN). As validações incluem:
- Comprimento exatamente 9 dígitos (`ErrInvalidSSNLength`)
- Somente dígitos numéricos (`ErrInvalidSSNNumbers`)
- Prefixo diferente de `000` (`ErrInvalidSSNPrefix`)
- Regra especial para SSNs iniciados com `9` (`ErrInvalidDigitPlace`)

O programa **não interrompe** ao encontrar um SSN inválido — ele loga o erro e continua processando os demais, demonstrando o padrão de log sem parada.

---

## 7. Debugging em Ambientes Live ou Restritos

Ambientes de produção têm restrições que impedem modificar o código durante a execução. A autora apresenta as seguintes estratégias:

| Estratégia | Descrição |
|---|---|
| **Entender o ambiente** | Antes de debugar, compreender o setup de deployment, configurações de rede e restrições de segurança |
| **Remote debugging com Delve** | Delve é o debugger oficial para Go; permite conectar a um processo em execução, inspecionar variáveis e definir breakpoints remotamente |
| **Observabilidade com `pprof`** | Ferramenta de profiling embutida no Go que coleta estatísticas de runtime sem alterar o código já deployado |
| **Utilizar níveis de log** | Diferentes log levels permitem ajustar o volume de informação logada conforme o ambiente — mas atenção ao não expor dados sensíveis |
| **Debuggers de IDE** | VS Code e JetBrains GoLand oferecem debugging visual com breakpoints, watch expressions e step-through; limitado a ambientes de desenvolvimento |
| **Feature flags e canary releases** | Ativar/desativar funcionalidades seletivamente em produção permite observar o impacto de mudanças em um subconjunto de usuários antes de uma release ampla |

> **Insight chave:** debugging é uma arte. O que funciona em um ambiente pode não funcionar em outro. A abordagem de "funciona na minha máquina" frequentemente leva a horas de investigação em ambientes de CI/CD. Testar mudanças em pequenos incrementos é a estratégia mais confiável.

---

## 8. Resumo e Relações de Causa e Efeito

```
Falta de testes           → Bugs difíceis de localizar
Erros ignorados           → Estados inválidos propagados silenciosamente
Prazos irrealistas        → Atalhos que introduzem bugs
Log insuficiente          → Impossível reproduzir bugs de produção
fmt.Printf com %T/%#v     → Inspeção de tipos e valores sem debugger
log.SetFlags(Llongfile)   → Log inclui arquivo e linha exata da origem
log.Fatal()               → Encerramento imediato sem defer
log.Panic()               → Pânico recuperável via defer
Delve + pprof             → Debugging e profiling em produção sem modificar código
```

---

## 9. Exercícios e Atividades do Capítulo

| Exercício/Atividade | Tema |
|---|---|
| Exercício 11.01 | Uso de `fmt.Println` com múltiplos argumentos |
| Exercício 11.02 | Impressão de valores decimais, binários e hexadecimais com larguras fixas |
| Exercício 11.03 | Uso de `%#v` para imprimir a representação Go de strings, slices, maps e structs |
| Atividade 11.01 | Programa completo de validação de SSN usando o pacote `log` para rastrear erros |

---

## 10. Referências e Links Úteis

- Documentação do pacote `fmt`: https://pkg.go.dev/fmt#hdr-Printing
- Código-fonte dos flags do pacote `log`: https://go.dev/src/log/log.go?s=8483:8506#L28
- Código do capítulo no GitHub: https://github.com/PacktPublishing/Go-Programming-From-Beginner-to-Professional-Second-Edition-/tree/main/Chapter11

---

*Resumo elaborado com base na leitura integral do Capítulo 11 (pp. 341–367).*
