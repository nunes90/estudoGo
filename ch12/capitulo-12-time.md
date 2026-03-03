# Capítulo 12 — About Time (Sobre Tempo em Go)

> Resumo detalhado e explicativo — baseado no livro *Go Programming: From Beginner to Professional* (Samantha Coyle)

---

## Visão Geral

O Capítulo 12 ensina como o Go lida com **variáveis de tempo** usando o pacote `time` da biblioteca padrão. Ao final do capítulo, você será capaz de:

- Criar variáveis de tempo e timestamps
- Comparar dois instantes de tempo
- Calcular a duração entre dois momentos
- Manipular tempo (adicionar/subtrair duração)
- Formatar datas e horas em diferentes padrões
- Converter horários entre fusos horários

---

## 1. O Pacote `time` — Estrutura Base

Antes de qualquer coisa, para usar tempo em Go você precisa importar o pacote `time`:

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    // código aqui
}
```

O pacote `time` é parte da biblioteca padrão do Go, ou seja, não precisa instalar nada extra. Ele fornece tipos, funções e constantes para trabalhar com datas e horas.

---

## 2. Criando Tempo (`Making Time`)

### 2.1 `time.Now()` — O instante atual

A função mais usada do pacote é `time.Now()`, que retorna o momento atual como um valor do tipo `time.Time`.

```go
start := time.Now()
fmt.Println("O script começou em:", start)
```

**Saída típica:**
```
O script começou em: 2023-09-27 08:19:33.8358274 +0200 CEST m=+0.001998701
```

Essa saída parece estranha no início, mas vamos entender cada parte:

| Parte | Significado |
|---|---|
| `2023-09-27` | Data: ano-mês-dia |
| `08:19:33.8358274` | Hora com nanosegundos |
| `+0200 CEST` | Fuso horário (UTC+2, horário de verão europeu) |
| `m=+0.001998701` | **Relógio monotônico** — tempo desde o início do processo |

> **O que é o relógio monotônico (`m=...`)?**
> O Go mantém dois tipos de relógio internamente:
> - **Wall clock** (relógio de parede): o horário real do sistema, como você vê no canto da tela. Pode ser ajustado pelo sistema operacional (ex: sincronização NTP).
> - **Monotonic clock** (relógio monotônico): mede tempo *decorrido* desde que o processo iniciou. Nunca vai para trás. É o que a parte `m=+0.001998701` representa.
>
> Quando você calcula diferenças de tempo (duração), o Go usa o relógio monotônico para garantir precisão, mesmo que o relógio do sistema seja ajustado no meio da execução.

### 2.2 `time.Sleep()` — Pausando a execução

```go
time.Sleep(2 * time.Second)
```

Essa função pausa o programa pelo tempo indicado. O argumento é do tipo `time.Duration`. O exemplo acima pausa por 2 segundos.

**Exemplo completo com início e fim:**

```go
start := time.Now()
fmt.Println("Iniciou em:", start)
fmt.Println("Salvando o mundo...")
time.Sleep(2 * time.Second)
end := time.Now()
fmt.Println("Terminou em:", end)
```

### 2.3 Extraindo partes da data

O tipo `time.Time` tem métodos para extrair cada componente da data/hora individualmente:

```go
day  := time.Now().Weekday()   // Dia da semana (ex: Monday)
hour := time.Now().Hour()      // Hora do dia (0–23)
date := time.Now()
date.Year()                    // Ano (ex: 2024)
date.Month()                   // Mês como tipo Month (ex: March)
date.Day()                     // Dia do mês (1–31)
```

**Exemplo prático — tomando decisões com base no dia:**

```go
day  := time.Now().Weekday()
hour := time.Now().Hour()
fmt.Println("Dia:", day, "Hora:", hour)

if day.String() == "Monday" {
    if hour >= 1 {
        fmt.Println("Realizando teste completo!")
    } else {
        fmt.Println("Realizando teste rápido!")
    }
} else {
    fmt.Println("Realizando teste rápido!")
}
```

Aqui o programa decide que tipo de teste executar com base no dia da semana — útil para scripts de automação que têm comportamentos diferentes conforme o dia.

### 2.4 Construindo nomes de arquivos com timestamp

Uma aplicação muito comum é criar nomes de log com a data atual. O desafio é que `time.Time` não é uma `string`, então precisamos **converter** as partes numéricas com `strconv.Itoa()`.

```go
import (
    "fmt"
    "strconv"
    "time"
)

func main() {
    appName := "HTTPCHECKER"
    action  := "BASIC"
    date    := time.Now()

    logFileName := appName + "_" + action + "_" +
        strconv.Itoa(date.Year()) + "_" +
        date.Month().String() + "_" +
        strconv.Itoa(date.Day()) + ".log"

    fmt.Println("Nome do log:", logFileName)
}
```

**Saída:**
```
Nome do log: HTTPCHECKER_BASIC_2024_March_16.log
```

> **Por que usar `strconv.Itoa()`?**
> Em Go, você não pode concatenar (`+`) um inteiro com uma string diretamente. `date.Year()` e `date.Day()` retornam `int`. A função `strconv.Itoa()` converte um `int` para `string`, permitindo a concatenação.
> `date.Month()` retorna um tipo especial `time.Month`, que tem o método `.String()` que já retorna o nome do mês como texto.

---

## 3. Comparando Tempo (`Comparing Time`)

O pacote `time` oferece três funções para comparar dois instantes de tempo:

| Função | O que faz |
|---|---|
| `t1.After(t2)` | Retorna `true` se `t1` é **posterior** a `t2` |
| `t1.Before(t2)` | Retorna `true` se `t1` é **anterior** a `t2` |
| `t1.Equal(t2)` | Retorna `true` se `t1` e `t2` são **iguais** |

### 3.1 `After()` — executar só depois de certo horário

```go
now := time.Now()
onlyAfter, err := time.Parse(time.RFC3339, "2020-11-01T22:08:41+00:00")
if err != nil {
    fmt.Println(err)
}

fmt.Println(now.After(onlyAfter)) // true ou false

if now.After(onlyAfter) {
    fmt.Println("Executando ações!")
} else {
    fmt.Println("Ainda não é a hora!")
}
```

> **O que é `time.Parse()`?**
> Essa função converte uma **string** de texto em um valor do tipo `time.Time`. Ela precisa de dois argumentos:
> 1. O **formato** da string (ex: `time.RFC3339`)
> 2. A **string** a ser convertida (ex: `"2020-11-01T22:08:41+00:00"`)
>
> Se a conversão falhar (formato errado), ela retorna um erro no segundo valor de retorno.

> **O que é RFC3339?**
> É um padrão internacional para representar datas e horas em texto, amplamente usado em APIs e sistemas web. O formato é: `YYYY-MM-DDTHH:MM:SS+HH:MM`.

### 3.2 `Equal()` — verificando igualdade exata

```go
now    := time.Now()
nowToo := now            // cópia do mesmo instante
time.Sleep(2 * time.Second)
later  := time.Now()     // instante diferente (2s depois)

if now.Equal(nowToo) {
    fmt.Println("São iguais!")    // imprime isso
}

if now.Equal(later) {
    fmt.Println("São iguais!")
} else {
    fmt.Println("São diferentes!") // imprime isso
}
```

**Por que `nowToo := now` cria uma cópia igual?**
Quando você atribui `nowToo := now`, está copiando o *valor* do tipo `time.Time`. Ambas as variáveis ficam com o mesmo instante gravado. Já `later := time.Now()` chama a função de novo, capturando um novo momento (2 segundos depois), por isso são diferentes.

---

## 4. Calculando Duração (`Duration Calculation`)

### 4.1 O tipo `time.Duration`

`time.Duration` é o tipo usado para representar um **intervalo de tempo** em Go. Internamente, é um inteiro de 64 bits que representa nanosegundos.

As constantes de duração disponíveis são:

```go
time.Hour         // 1 hora
time.Minute       // 1 minuto
time.Second       // 1 segundo
time.Millisecond  // 1 milissegundo (1/1000 de segundo)
time.Microsecond  // 1 microssegundo (1/1.000.000 de segundo)
time.Nanosecond   // 1 nanossegundo (1/1.000.000.000 de segundo)
```

### 4.2 `end.Sub(start)` — calculando a duração entre dois instantes

```go
start := time.Now()
// ... código que leva algum tempo ...
sum := 0
for i := 1; i < 10000000000; i++ {
    sum += i
}
end := time.Now()

duration := end.Sub(start)  // retorna time.Duration
```

O método `Sub()` subtrai dois valores `time.Time` e retorna um `time.Duration`. A partir daí, você pode extrair a duração em diferentes resoluções:

```go
fmt.Println("Duração em horas:       ", duration.Hours())
fmt.Println("Duração em minutos:     ", duration.Minutes())
fmt.Println("Duração em segundos:    ", duration.Seconds())
fmt.Println("Duração em nanosegundos:", duration.Nanoseconds())
```

> **Atenção:** `Hours()`, `Minutes()` e `Seconds()` retornam `float64`. `Nanoseconds()` retorna `int64`.
> Se você precisar de dias, semanas ou meses, calcule a partir dessas resoluções (ex: `duration.Hours() / 24` para dias).

### 4.3 Comparando duração com um prazo (SLA)

Uma situação prática comum: você tem um **prazo máximo** para executar uma operação e quer verificar se foi cumprido.

```go
// Define prazo: 600ms * 10 = 6000ms = 6 segundos
deadlineSeconds := time.Duration((600 * 10) * time.Millisecond)

start := time.Now()
fmt.Println("Prazo da transação:", deadlineSeconds)
fmt.Println("Transação iniciada em:", start)

sum := 0
for i := 1; i < 25000000000; i++ {
    sum += i
}
end := time.Now()

duration := end.Sub(start)

// Converte duration para time.Duration usando resolução Nanosecond
transactionTime := time.Duration(duration.Nanoseconds()) * time.Nanosecond

fmt.Println("Transação concluída em:", end, duration)

if transactionTime <= deadlineSeconds {
    fmt.Println("Performance OK! Transação concluída em", transactionTime)
} else {
    fmt.Println("Problema de performance! Transação demorou", transactionTime)
}
```

> **Por que converter para `time.Duration` novamente?**
> O método `duration.Nanoseconds()` retorna um `int64` (número puro). Para compará-lo diretamente com `deadlineSeconds` (que é um `time.Duration`), precisamos convertê-lo de volta para `time.Duration`. Fazemos isso multiplicando pelo `time.Nanosecond`, que é `1` no sistema interno do Go — essencialmente "rotulando" o número como sendo em nanosegundos.

---

## 5. Manipulando Tempo (`Managing Time`)

O pacote `time` oferece duas funções para **manipular** um instante:

| Função | O que faz |
|---|---|
| `t.Add(d)` | Retorna `t + d` (adiciona uma duração) |
| `t.Sub(t2)` | Retorna `t - t2` (diferença entre dois instantes) |

### 5.1 `Add()` — somando tempo

```go
timeToManipulate := time.Now()
toBeAdded := time.Duration(10 * time.Second)  // 10 segundos

fmt.Println("Tempo original:", timeToManipulate)
fmt.Printf("%v depois: %v\n", toBeAdded, timeToManipulate.Add(toBeAdded))
```

**Saída:**
```
Tempo original: 2023-10-18 08:49:53.1499273 +0200 CEST m=+0.001994601
10s depois: 2023-10-18 08:50:03.1499273 +0200 CEST m=+10.001994601
```

### 5.2 Subtraindo tempo com `Add()` negativo

Para *remover* uma duração, use `Add()` com um valor negativo. `Sub()` não é feito para isso diretamente.

```go
toBeAdded := time.Duration(-10 * time.Minute)  // menos 10 minutos
fmt.Printf("%v depois: %v\n", toBeAdded, timeToManipulate.Add(toBeAdded))
```

**Saída:**
```
-10m0s depois: 2023-10-18 08:40:36.5950116 +0200 CEST
```

Isso mostra qual era o horário **10 minutos atrás**.

---

## 6. Formatando Tempo (`Formatting Time`)

Essa é a parte mais importante e também a mais confusa do capítulo.

### 6.1 O problema: datas feias

Por padrão, ao imprimir um `time.Time`, o Go exibe algo assim:
```
2023-09-27 13:50:58.2715452 +0200 CEST m=+0.002992801
```

Isso é difícil de ler para humanos. O Go permite formatar esse valor de maneiras mais amigáveis.

### 6.2 A data de referência mágica do Go

Go usa uma abordagem **única e confusa no início**: para definir o formato de saída, você escreve a **data de referência** no formato desejado. Essa data de referência é fixa:

```
Mon Jan 2 15:04:05 -0700 MST 2006
  0   1  2    3   4   5      6
```

Cada número tem um significado:
| Posição | Valor | Representa |
|---|---|---|
| 0 | `Mon` | Dia da semana abreviado |
| 1 | `Jan` | Mês abreviado |
| 2 | `2` | Dia do mês |
| 3 | `15` | Hora (formato 24h) |
| 4 | `04` | Minutos |
| 5 | `05` | Segundos |
| 6 | `2006` | Ano |
| — | `-0700` | Fuso horário |

> **Por que essa data específica?**
> Porque 1-2-3-4-5-6 em ordem: janeiro=1, 2=dia, 15:04:05=hora/min/seg, -0700=timezone, 2006=ano. É um mnemônico que o Go usa internamente.

### 6.3 Constantes de formato pré-definidas

O Go já oferece formatos prontos:

```go
time.ANSIC       // "Mon Jan _2 15:04:05 2006"
time.RFC3339     // "2006-01-02T15:04:05Z07:00"
time.UnixDate    // "Mon Jan _2 15:04:05 MST 2006"
time.RFC822      // "02 Jan 06 15:04 MST"
time.Kitchen     // "3:04PM"
// ... entre outros
```

### 6.4 `time.Format()` — formatando para string

```go
t := time.Now()
fmt.Println(t.Format(time.ANSIC))
// Saída: Thu Oct 17 13:56:03 2023
```

Você também pode criar seu próprio formato:

```go
fmt.Println(t.Format("02/01/2006 15:04:05"))
// Saída: 17/10/2023 13:56:03
```

### 6.5 `time.Parse()` — convertendo string para `time.Time`

```go
t1, err := time.Parse(time.RFC3339, "2019-09-27T22:18:11+00:00")
if err != nil {
    fmt.Println(err)
}
fmt.Println("RFC3339:", t1)
// Saída: RFC3339: 2019-09-27 22:18:11 +0000 +0000
```

**Cuidado:** se o formato não bater com a string, o erro é gerado e o valor retornado é o tempo zero (`0001-01-01 00:00:00`).

```go
// ERRO: formato UnixDate não aceita este formato de string
t2, err := time.Parse(time.UnixDate, "2019-09-27T22:18:11+00:00")
// err != nil
// t2 = 0001-01-01 00:00:00 +0000 UTC  (valor zero — inválido)
```

> **UnixDate** espera algo como: `Mon Sep 27 18:24:05 2019`
> **ANSIC** espera algo como: `Thu Oct 17 13:56:03 2023`

### 6.6 `time.Date()` — criando uma data específica

Você pode criar um instante de tempo específico passando cada componente:

```go
// Sintaxe: time.Date(year, month, day, hour, min, sec, nanosec, location)
date := time.Date(2019, 9, 27, 18, 50, 48, 324359102, time.UTC)
fmt.Println(date)
// Saída: 2019-09-27 18:50:48.324359102 +0000 UTC
```

### 6.7 `AddDate()` — adicionando anos, meses e dias

```go
date     := time.Date(2019, 9, 27, 18, 50, 48, 324359102, time.UTC)
nextDate := date.AddDate(1, 2, 3)  // +1 ano, +2 meses, +3 dias
fmt.Println(nextDate)
// Saída: 2020-11-30 18:50:48.324359102 +0000 UTC
```

> **Diferença entre `Add()` e `AddDate()`:**
> - `Add()` recebe um `time.Duration` (nanosegundos, segundos, minutos, horas)
> - `AddDate()` recebe anos, meses e dias como inteiros — mais intuitivo para datas

---

## 7. Fusos Horários (`Time Zones`)

### 7.1 `time.LoadLocation()` — carregando um fuso horário

```go
losAngeles, err := time.LoadLocation("America/Los_Angeles")
if err != nil {
    fmt.Println(err)
}
```

Os nomes de fusos horários seguem o padrão IANA (ex: `"America/Sao_Paulo"`, `"Europe/London"`, `"Asia/Tokyo"`).

### 7.2 `.In()` — convertendo para outro fuso horário

```go
current    := time.Now()
losAngeles, _ := time.LoadLocation("America/Los_Angeles")

fmt.Println("Hora local:          ", current.Format(time.ANSIC))
fmt.Println("Hora em Los Angeles: ", current.In(losAngeles).Format(time.ANSIC))
```

**Saída:**
```
Hora local:           Fri Oct 18 08:14:48 2019
Hora em Los Angeles:  Thu Oct 17 23:14:48 2019
```

O método `In()` **não altera** o valor do tempo — o instante é o mesmo. O que muda é a **representação visual** naquele fuso horário.

---

## 8. Exercícios do Capítulo

### Exercício 12.01 — Retornando um timestamp formatado

Cria uma função que retorna a hora atual no formato ANSIC:

```go
func whatstheclock() string {
    return time.Now().Format(time.ANSIC)
}

func main() {
    fmt.Println(whatstheclock())
}
// Saída: Thu Oct 17 13:56:03 2023
```

### Exercício 12.02 — Calculando duração de execução

Cria uma função reutilizável que recebe dois `time.Time` e retorna uma string legível com horas, minutos e segundos decorridos:

```go
func elapsedTime(start time.Time, end time.Time) string {
    elapsed := end.Sub(start)
    hours   := strconv.Itoa(int(elapsed.Hours()))
    minutes := strconv.Itoa(int(elapsed.Minutes()))
    seconds := strconv.Itoa(int(elapsed.Seconds()))
    return "Tempo total decorrido: " + hours + " hora(s) e " +
           minutes + " minuto(s) e " + seconds + " segundo(s)!"
}

func main() {
    start := time.Now()
    time.Sleep(2 * time.Second)
    end := time.Now()
    fmt.Println(elapsedTime(start, end))
}
// Saída: Tempo total decorrido: 0 hora(s) e 0 minuto(s) e 2 segundo(s)!
```

> **Por que `int(elapsed.Hours())` e não direto?**
> `elapsed.Hours()` retorna `float64`. `strconv.Itoa()` só aceita `int`. Por isso convertemos com `int(...)`, que descarta a parte decimal (arredonda para baixo).

### Exercício 12.03 — Hora em outro fuso horário

Cria uma função que recebe o nome de um fuso horário e retorna a hora local e a hora naquele fuso:

```go
func timeDiff(timezone string) (string, string) {
    current    := time.Now()
    remoteZone, err := time.LoadLocation(timezone)
    if err != nil {
        fmt.Println(err)
    }
    remoteTime := current.In(remoteZone)
    fmt.Println("Hora atual:            ", current.Format(time.ANSIC))
    fmt.Println("Hora em", timezone, ":", remoteTime.Format(time.ANSIC))
    return current.Format(time.ANSIC), remoteTime.Format(time.ANSIC)
}

func main() {
    fmt.Println(timeDiff("America/Los_Angeles"))
    fmt.Println(timeDiff("America/Sao_Paulo"))
}
```

---

## 9. Resumo dos Conceitos-Chave

| Conceito | Função/Tipo | Para que serve |
|---|---|---|
| Hora atual | `time.Now()` | Captura o instante presente |
| Tipo de tempo | `time.Time` | Representa um ponto no tempo |
| Tipo de duração | `time.Duration` | Representa um intervalo de tempo |
| Pausar execução | `time.Sleep(d)` | Pausa pelo tempo `d` |
| Comparar (depois) | `t.After(t2)` | `true` se `t` é posterior a `t2` |
| Comparar (antes) | `t.Before(t2)` | `true` se `t` é anterior a `t2` |
| Comparar (igual) | `t.Equal(t2)` | `true` se `t` é igual a `t2` |
| Calcular duração | `t.Sub(t2)` | Retorna `time.Duration` entre dois instantes |
| Adicionar duração | `t.Add(d)` | Soma `d` ao instante `t` |
| Adicionar data | `t.AddDate(y,m,d)` | Adiciona anos, meses e dias |
| Criar data | `time.Date(...)` | Cria um instante específico |
| Formatar | `t.Format(layout)` | Converte `time.Time` para `string` |
| Parsear | `time.Parse(layout, s)` | Converte `string` para `time.Time` |
| Fuso horário | `time.LoadLocation(nome)` | Carrega um fuso horário pelo nome |
| Converter fuso | `t.In(loc)` | Representa o instante em outro fuso |
| Converter int→string | `strconv.Itoa(n)` | Necessário para concatenar inteiros com strings |

---

## 10. Armadilhas Comuns e Dicas

**1. Não confunda `Sub()` e `Add()`**
- `t1.Sub(t2)` → diferença entre dois instantes → retorna `time.Duration`
- `t.Add(d)` → soma uma duração a um instante → retorna `time.Time`

**2. `time.Parse()` é sensível ao formato**
Se o formato não corresponder exatamente à string, retorna erro e tempo zero. Sempre verifique o `err`.

**3. `Month()` retorna um tipo especial**
`date.Month()` retorna `time.Month`, não um `int`. Use `int(date.Month())` para obter o número e `.String()` para o nome.

**4. Para concatenar com strings, converta antes**
`strconv.Itoa(date.Year())` para inteiros e `date.Month().String()` para o mês como texto.

**5. O formato de layout do Go é fixo**
Sempre use a data de referência `Mon Jan 2 15:04:05 MST 2006`. Não invente outros números.

**6. Fusos horários usam o banco IANA**
Exemplos: `"America/Sao_Paulo"`, `"America/New_York"`, `"Europe/Berlin"`, `"Asia/Tokyo"`.

---

## 11. Referência Rápida — Formatos de Layout

```go
// Formatos pré-definidos (constantes do pacote time)
time.ANSIC       = "Mon Jan _2 15:04:05 2006"
time.UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
time.RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
time.RFC822      = "02 Jan 06 15:04 MST"
time.RFC822Z     = "02 Jan 06 15:04 -0700"
time.RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
time.RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
time.RFC3339     = "2006-01-02T15:04:05Z07:00"
time.Kitchen     = "3:04PM"
time.Stamp       = "Jan _2 15:04:05"
time.DateTime    = "2006-01-02 15:04:05"
time.DateOnly    = "2006-01-02"
time.TimeOnly    = "15:04:05"

// Formatos customizados (exemplos)
"02/01/2006"          // → 17/10/2023
"2006-01-02"          // → 2023-10-17
"15:04:05"            // → 13:56:03
"02/01/2006 15:04:05" // → 17/10/2023 13:56:03
"Monday, 02 January 2006" // → Tuesday, 17 October 2023
```

---

*Resumo elaborado com base no Capítulo 12 do livro "Go Programming: From Beginner to Professional" (2ª Edição, Samantha Coyle — Packt Publishing)*
