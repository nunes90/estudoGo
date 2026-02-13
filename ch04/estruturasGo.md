# Estruturas de Dados e Tipos em Go

Resumo completo baseado nos exercícios e atividades do Capítulo 4.

---

## 1. Arrays

Arrays em Go são coleções de **tamanho fixo** e **tipo homogêneo**. O tamanho faz parte do tipo — `[5]int` e `[10]int` são tipos diferentes e incompatíveis.

### 1.1 Declaração e Inicialização

```go
// Zero-value: todos os elementos recebem o valor zero do tipo
var arr [10]int                        // [0 0 0 0 0 0 0 0 0 0]

// Literal com valores explícitos
arr := [5]int{1, 2, 3, 4, 5}

// Tamanho inferido pelo compilador com [...]
arr := [...]int{0, 0, 0, 0, 0}        // equivale a [5]int

// Inicialização por chave (índice) — permite preencher posições específicas
arr := [...]int{9: 0}                  // array de 10 elementos, todos zero
arr := [10]int{1, 9: 10, 4: 5}        // posições 0=1, 4=5, 9=10, restante=0
```

### 1.2 Leitura e Escrita

```go
val := arr[2]       // leitura por índice
arr[3] = 99         // escrita por índice
```

### 1.3 Comparação

Arrays de **mesmo tipo e tamanho** podem ser comparados diretamente com `==`:

```go
arr1 := [5]int{0, 0, 0, 0, 0}
arr2 := [...]int{0, 0, 0, 0, 0}
fmt.Println(arr1 == arr2) // true
```

### 1.4 Iteração

```go
for i := 0; i < len(arr); i++ {
    fmt.Println(i, arr[i])
}
```

### 1.5 Passagem por Valor

Arrays são **copiados** ao serem passados como argumento de função. Alterações dentro da função **não** afetam o array original:

```go
func dobrar(arr [5]int) [5]int {
    for i := range arr {
        arr[i] *= 2
    }
    return arr   // precisa retornar e reatribuir
}
```

### 1.6 Funções Importantes

| Função      | Descrição                          |
|-------------|------------------------------------|
| `len(arr)`  | Retorna o número de elementos      |

---

## 2. Slices

Slices são **referências dinâmicas** a arrays subjacentes. Diferente de arrays, o tamanho **não** faz parte do tipo — `[]int` aceita qualquer quantidade de elementos.

### 2.1 Declaração e Inicialização

```go
// Declaração nil (len=0, cap=0)
var s []int

// Literal
s := []int{1, 2, 3, 4, 5}

// Com make — controle de length e capacity
s := make([]int, 10)       // len=10, cap=10
s := make([]int, 10, 50)   // len=10, cap=50
```

### 2.2 Append — Adicionar Elementos

`append` retorna um **novo slice** (pode ou não realocar o array subjacente):

```go
// Adicionar um elemento
s = append(s, 6)

// Adicionar múltiplos elementos
s = append(s, 7, 8, 9)

// Adicionar todos os elementos de outro slice (spread com ...)
s = append(s, outroSlice...)
```

### 2.3 Sub-slicing — Criar Slices a Partir de Outro

```go
s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

s[0:1]          // [1]            — do índice 0 até 1 (exclusivo)
s[:5]           // [1 2 3 4 5]    — primeiros 5
s[5:]           // [6 7 8 9]      — do índice 5 até o fim
s[2:7]          // [3 4 5 6 7]    — do índice 2 ao 7
s[len(s)-1:]    // [9]            — último elemento
```

### 2.4 Length vs Capacity

```go
len(s)   // número de elementos acessíveis
cap(s)   // tamanho total do array subjacente a partir do início do slice
```

### 2.5 Compartilhamento de Array Subjacente (Backing Array)

Slices podem **compartilhar** o mesmo array por baixo. Isso é crítico para entender mutações:

```go
// LINKED — s2 aponta para o mesmo backing array
s1 := []int{1, 2, 3, 4, 5}
s2 := s1       // mesma referência
s3 := s1[:]    // mesma referência
s1[3] = 99     // s2[3] e s3[3] também viram 99

// UNLINKED — append pode realocar, quebrando o vínculo
s1 = append(s1, 6)   // se cap estourou, novo array → s2 fica independente
s1[3] = 99            // s2[3] continua com valor original
```

### 2.6 Cópia Profunda (Deep Copy)

Duas formas de criar uma cópia independente:

```go
// Com copy()
s2 := make([]int, len(s1))
n := copy(s2, s1)     // n = número de elementos copiados

// Com append idiomático
s2 := append([]int{}, s1...)
```

### 2.7 Deletar Elemento de um Slice

Go não tem função nativa de delete para slices. O padrão idiomático é:

```go
// Remover o elemento no índice i
s = append(s[:i], s[i+1:]...)

// Exemplo: remover índice 2
sli := []string{"A", "B", "C", "D", "E"}
sli = append(sli[:2], sli[3:]...)   // ["A", "B", "D", "E"]
```

### 2.8 Rotacionar um Slice

```go
// Mover último elemento para o início
week := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
week = append(week[6:], week[:6]...)   // ["Sun", "Mon", "Tue", ...]
```

### 2.9 Funções Importantes

| Função/Operação              | Descrição                                              |
|------------------------------|--------------------------------------------------------|
| `len(s)`                     | Número de elementos no slice                           |
| `cap(s)`                     | Capacidade total do backing array                      |
| `append(s, elems...)`        | Adiciona elementos, retorna novo slice                 |
| `copy(dst, src)`             | Copia elementos de src para dst, retorna qtd copiada   |
| `make([]T, len)`             | Cria slice com tamanho definido                        |
| `make([]T, len, cap)`        | Cria slice com tamanho e capacidade definidos           |
| `s[low:high]`                | Sub-slice do índice low até high (exclusivo)           |

---

## 3. Maps

Maps são coleções de **pares chave-valor** com acesso em tempo constante (hash table). As chaves devem ser de um tipo comparável.

### 3.1 Declaração e Inicialização

```go
// Literal
users := map[string]string{
    "305": "Sue",
    "204": "Bob",
    "631": "Jake",
}

// Variável de pacote
var users = map[string]string{
    "305": "Sue",
    "204": "Bob",
}

// Com make (mapa vazio)
m := make(map[string]int)
```

### 3.2 Leitura e Escrita

```go
// Escrita / atualização
users["073"] = "Tracy"

// Leitura simples (retorna zero-value se não existir)
name := users["305"]

// Leitura segura com comma-ok (verificar existência)
name, exists := users["305"]
if !exists {
    fmt.Println("Chave não encontrada")
}
```

### 3.3 Deletar Entrada

```go
delete(users, "073")   // remove a chave "073" do mapa
```

### 3.4 Iteração

A ordem de iteração em maps é **aleatória** (não garantida):

```go
for key, value := range users {
    fmt.Println("ID:", key, "Nome:", value)
}
```

### 3.5 Map com Struct como Chave — Padrão Set

É possível usar structs como chaves de map (desde que todos os campos sejam comparáveis). Usando `struct{}` como valor, implementa-se um **set** com custo zero de memória:

```go
type locale struct {
    language string
    country  string
}

func getLocales() map[locale]struct{} {
    supported := make(map[locale]struct{}, 5)
    // struct{}{} → struct{} é o tipo (zero bytes), {} é a inicialização
    supported[locale{"en", "US"}] = struct{}{}
    supported[locale{"fr", "FR"}] = struct{}{}
    return supported
}

func localeExists(l locale) bool {
    _, exists := getLocales()[l]
    return exists
}
```

### 3.6 Funções Importantes

| Função/Operação            | Descrição                                              |
|----------------------------|--------------------------------------------------------|
| `m[key]`                   | Lê o valor associado à chave                           |
| `m[key] = value`           | Insere ou atualiza um par chave-valor                  |
| `v, ok := m[key]`          | Leitura segura — ok indica se a chave existe           |
| `delete(m, key)`           | Remove uma entrada do mapa                             |
| `len(m)`                   | Número de entradas no mapa                             |
| `make(map[K]V)`            | Cria um mapa vazio                                     |
| `for k, v := range m`     | Itera sobre todos os pares chave-valor                 |

---

## 4. Tipos Customizados

Go permite criar tipos nomeados baseados em tipos primitivos. Tipos customizados **não são intercambiáveis** com o tipo base sem conversão explícita:

```go
type id string

var id1 id                          // zero-value: ""
var id2 id = "1234-5678"            // inicialização com valor
var id3 id
id3 = "1234-5678"

// Comparação entre valores do mesmo tipo customizado
fmt.Println(id2 == id3)             // true

// Conversão explícita para o tipo base
fmt.Println(string(id2) == "1234-5678")  // true

// ERRO: não é possível comparar diretamente id com string
// fmt.Println(id2 == "1234-5678")  // erro de compilação
```

Tipos customizados são úteis para dar **semântica** a valores e evitar confusão entre tipos que possuem a mesma representação base.

---

## 5. Structs

Structs são **tipos compostos** que agrupam campos nomeados de tipos diferentes. São a principal forma de modelar dados em Go.

### 5.1 Declaração

```go
type user struct {
    name    string
    age     int
    balance float64
    member  bool
}
```

### 5.2 Quatro Formas de Inicialização

```go
// 1. Campos nomeados (recomendado — ordem não importa)
u1 := user{
    name:    "Tracy",
    age:     51,
    balance: 98.43,
    member:  true,
}

// 2. Campos nomeados parciais (campos omitidos recebem zero-value)
u2 := user{
    age:  19,
    name: "Nick",
}
// u2.balance == 0, u2.member == false

// 3. Posicional (depende da ordem — frágil, evitar)
u3 := user{"Bob", 25, 0, false}

// 4. Zero-value + atribuição campo a campo
var u4 user
u4.name = "Sue"
u4.age = 31
```

### 5.3 Leitura e Escrita de Campos

```go
fmt.Println(u1.name)    // leitura: "Tracy"
u1.balance = 150.00     // escrita
```

### 5.4 Structs Anônimas

Structs podem ser declaradas inline, sem criar um tipo nomeado:

```go
// Declaração + inicialização inline
point1 := struct {
    x int
    y int
}{10, 10}

// Declaração inline + atribuição
point2 := struct {
    x int
    y int
}{}
point2.x = 10
point2.y = 5
```

### 5.5 Comparação de Structs

Structs são comparáveis com `==` se **todos os campos forem comparáveis**. Go permite comparar structs anônimas com structs nomeadas se os campos (nome, tipo e ordem) forem idênticos:

```go
type point struct{ x, y int }

point1 := struct{ x, y int }{10, 10}
point3 := point{10, 10}

fmt.Println(point1 == point3)  // true — mesma estrutura
```

### 5.6 Struct Embedding (Composição)

Go não tem herança, mas permite **embedding** — incluir um tipo dentro de outro. Os campos do tipo embutido são **promovidos** (acessíveis diretamente):

```go
type name string

type location struct{ x, y int }

type size struct{ width, height int }

type dot struct {
    name              // tipo customizado embutido
    location          // struct embutida
    size              // struct embutida
}
```

Formas de acessar campos promovidos:

```go
var d dot

// Acesso promovido (atalho)
d.name = "A"
d.x = 5           // promovido de location
d.width = 10      // promovido de size

// Acesso explícito (via tipo embutido)
d.location.x = 5
d.size.width = 10

// Inicialização com campos nomeados
d2 := dot{
    name:     "B",
    location: location{x: 13, y: 27},
    size:     size{width: 5, height: 7},
}
```

### 5.7 Struct Vazia — `struct{}`

`struct{}` é um tipo que ocupa **zero bytes** de memória. É usado como valor em maps para implementar sets:

```go
// Set idiomático em Go
set := map[string]struct{}{}
set["item"] = struct{}{}

_, exists := set["item"]  // exists == true
```

### 5.8 Resumo de Structs

| Operação                     | Sintaxe                                        |
|------------------------------|------------------------------------------------|
| Declarar tipo                | `type T struct { campo Tipo }`                 |
| Inicializar (nomeado)        | `T{campo: valor}`                              |
| Inicializar (posicional)     | `T{val1, val2, ...}`                           |
| Ler campo                    | `v.campo`                                      |
| Escrever campo               | `v.campo = valor`                              |
| Comparar                     | `v1 == v2` (se todos os campos forem comparáveis) |
| Embutir (embedding)          | Declarar tipo sem nome de campo dentro da struct |
| Acessar campo promovido      | `v.campoDaTipoEmbutida`                        |

---

## 6. Conversões Numéricas

Go exige **conversão explícita** entre tipos numéricos, mesmo entre variações do mesmo tipo base (ex: `int` e `int8`):

```go
var i8 int8 = math.MaxInt8   // 127
i := 128
f64 := 3.14
```

### 6.1 Tipos de Conversão

```go
// Widening (ampliação) — seguro, sem perda
int64(i8)       // 127 → 127

// Narrowing (estreitamento) — PERIGOSO, pode causar overflow
int8(i)         // 128 → -128 (overflow! wraps around)

// Int para float — seguro
float64(i8)     // 127 → 127.0

// Float para int — TRUNCA (não arredonda)
int(f64)        // 3.14 → 3
int(3.99)       // 3.99 → 3  (NÃO é 4)
```

### 6.2 Cuidados

| Conversão        | Risco                                          |
|------------------|------------------------------------------------|
| Widening         | Nenhum — o valor cabe no tipo maior            |
| Narrowing        | **Overflow silencioso** — o valor "dá a volta" |
| Float → Int      | **Truncamento** — parte decimal é descartada   |
| Int → Float      | Possível perda de precisão para valores grandes|

---

## 7. Interface Vazia — `interface{}` / `any`

### 7.1 O que é `interface{}`?

Em Go, uma interface define **quais métodos** um tipo deve ter. `interface{}` é uma interface que **não exige nenhum método** — portanto, **todos os tipos a satisfazem**.

Desde Go 1.18, existe o alias `any`, que é equivalente a `interface{}`:

```go
var v1 interface{} = 42
var v2 any = "hello"          // idêntico a interface{}
```

### 7.2 Por que `interface{}` aceita qualquer tipo?

Qualquer tipo em Go — primitivos, structs, slices, maps, tipos customizados — possui zero ou mais métodos. Como `interface{}` exige **zero métodos**, todo tipo automaticamente a satisfaz. É por isso que funções como `fmt.Println` aceitam qualquer argumento.

### 7.3 Limitação: acesso bloqueado

Quando um valor é colocado numa variável `interface{}`, Go **bloqueia o acesso** aos métodos, campos e valor interno. O valor continua intacto por dentro, mas o compilador não permite usá-lo diretamente:

```go
var v interface{} = "hello"

// v + " world"    // ERRO: operação não permitida em interface{}
// len(v)          // ERRO: interface{} não tem método len
```

### 7.4 Slice de `interface{}`

Para armazenar valores de tipos diferentes numa mesma coleção:

```go
func getData() []interface{} {
    return []interface{}{
        1,
        3.14,
        "hello",
        true,
        struct{}{},
    }
}
```

---

## 8. Type Assertions

Type assertion é o mecanismo para **recuperar o tipo concreto** de um valor armazenado em `interface{}` / `any`. A sintaxe é `valor.(Tipo)`.

### 8.1 Forma Segura — comma-ok

Nunca causa panic. Retorna o valor convertido e um booleano indicando sucesso:

```go
func doubler(v interface{}) (string, error) {
    // Tenta int
    if i, ok := v.(int); ok {
        return fmt.Sprint(i * 2), nil
    }
    // Tenta string
    if s, ok := v.(string); ok {
        return s + s, nil
    }
    return "", errors.New("unsupported type")
}

doubler(5)       // "10", nil
doubler("Go")    // "GoGo", nil
doubler(3.14)    // "", error
```

### 8.2 Forma Direta — sem comma-ok (PERIGOSA)

Se a asserção falhar, Go **levanta um panic**:

```go
var v interface{} = "hello"

s := v.(string)    // OK: s = "hello"
i := v.(int)       // PANIC: interface conversion error
```

### 8.3 Type Switch

Alternativa mais limpa quando há múltiplos tipos possíveis. A variável `t` assume o tipo do caso correspondente:

```go
func doubler(v interface{}) (string, error) {
    switch t := v.(type) {
    case string:
        return t + t, nil                    // t é string
    case bool:
        if t {
            return "truetrue", nil
        }
        return "falsefalse", nil
    case float32, float64:                   // caso com múltiplos tipos
        // t continua como interface{} — precisa de assertion adicional
        if f, ok := t.(float64); ok {
            return fmt.Sprint(f * 2), nil
        }
        return fmt.Sprint(t.(float32) * 2), nil
    case int:
        return fmt.Sprint(t * 2), nil        // t é int
    default:
        return "", errors.New("unsupported type")
    }
}
```

### 8.4 Type Switch sem Variável (apenas classificação)

Quando só se quer identificar o tipo, sem usar o valor tipado:

```go
func getTypeName(v interface{}) string {
    switch v.(type) {
    case int, int32, int64:
        return "int"
    case float64, float32:
        return "float"
    case bool:
        return "bool"
    case string:
        return "string"
    default:
        return "unknown"
    }
}
```

### 8.5 Como funciona por baixo

Go **não remove nada** do valor ao colocá-lo em `interface{}`. O compilador apenas bloqueia o acesso porque não consegue fazer verificações de tipo em tempo de compilação. Type assertion instrui o Go a realizar essas verificações **em tempo de execução**:

```
Valor concreto "hello" (string)
        │
        ▼
  interface{} / any     ← compilador bloqueia acesso ao tipo real
        │
        ▼
  Type assertion         ← "eu sei que é string, verifique em runtime"
   v.(string)
        │
        ▼
  Valor "hello" como string  ← Go verifica e libera (ou panic/false)
```

### 8.6 Cuidados com Type Assertions

| Aspecto                      | Detalhe                                                    |
|------------------------------|-----------------------------------------------------------|
| Forma `v.(T)`               | Causa **panic** se o tipo não bater                        |
| Forma `v, ok := v.(T)`      | Retorna `ok = false` — **nunca** causa panic               |
| Em type switch multi-caso   | A variável mantém tipo `interface{}` — precisa re-assertar |
| Verificação em runtime      | Erros de tipo viram **erros em tempo de execução**         |
| Recomendação                | **Sempre** usar a forma comma-ok ou type switch            |

---

## 9. Comparativo Rápido — Coleções

| Característica     | Array               | Slice                    | Map                       |
|--------------------|---------------------|--------------------------|---------------------------|
| Tamanho            | Fixo (parte do tipo)| Dinâmico                 | Dinâmico                  |
| Tipo em Go         | `[N]T`              | `[]T`                    | `map[K]V`                 |
| Zero-value         | Array de zeros      | `nil`                    | `nil`                     |
| Comparável com `==`| Sim                 | Não                      | Não                       |
| Passagem p/ função | Por valor (cópia)   | Por referência (ponteiro)| Por referência (ponteiro) |
| Acesso             | Por índice `[i]`    | Por índice `[i]`         | Por chave `[key]`         |
| Inicialização      | `var`, literal, `[...]` | `var`, literal, `make` | Literal, `make`          |
| Adicionar elemento | Não é possível      | `append()`               | `m[key] = val`            |
| Remover elemento   | Não é possível      | `append(s[:i], s[i+1:]...)` | `delete(m, key)`     |
| Iterar             | `for i` / `for range` | `for i` / `for range` | `for k, v := range`      |

---

## 10. Mapa de Exercícios e Atividades

| Arquivo        | Diretório       | Conceito Principal                                        |
|----------------|-----------------|-----------------------------------------------------------|
| ex4.01         | `ex4.01/`       | Array: declaração e zero-value                            |
| ex4.02         | `ex4.02/`       | Array: comparação com `==`, `[...]`                       |
| ex4.03         | `ex4.03/`       | Array: inicialização por índice (keyed)                   |
| ex4.04         | `ex4.04/`       | Array: leitura de elementos por índice                    |
| ex4.05         | `ex4.05/`       | Array: escrita de elementos por índice                    |
| ex4.06         | `ex4.06/`       | Array: loop `for i` com mutação in-place                  |
| ex4.07         | `ex4.07/`       | Array: passagem por valor para funções                    |
| ex4.08         | `ex4.08/`       | Slice: nil slice + `append` com `os.Args`                 |
| ex4.09         | `ex4.09/`       | Slice: `append` múltiplo e spread `...`                   |
| ex4.10         | `ex4.10/`       | Slice: sub-slicing `s[low:high]`                          |
| ex4.11         | `ex4.11/`       | Slice: `make`, `len`, `cap`                               |
| ex4.12         | `ex4.12/`       | Slice: backing array compartilhado, `copy`, deep copy     |
| ex4.13         | `ex4.13/`       | Map: literal e adição de entradas                         |
| ex4.14         | `ex4.14/`       | Map: comma-ok e iteração com `range`                      |
| ex4.15         | `ex4.15/`       | Map: `delete`, variável de pacote                         |
| ex4.16         | `ex4.16/`       | Tipo customizado sobre primitivo (`type id string`)       |
| ex4.17         | `ex4.17/`       | Struct: declaração e 4 formas de inicialização            |
| ex4.18         | `ex4.18/`       | Struct: anônimas e comparação                             |
| ex4.19         | `ex4.19/`       | Struct: embedding e campos promovidos                     |
| ex4.20         | `ex4.20/`       | Conversões numéricas (widening, narrowing, truncamento)   |
| ex4.21         | `ex4.21/`       | Type assertion: forma comma-ok com `interface{}`          |
| ex4.22         | `ex4.22/`       | Type switch: multi-caso e re-asserção                     |
| activity4.01   | `activity4.01/` | Array: preencher com loop                                 |
| activity4.02   | `activity4.02/` | Map: lookup com `os.Args` e comma-ok                      |
| activity4.03   | `activity4.03/` | Slice: rotação com `append` + spread                      |
| activity4.04   | `activity4.04/` | Slice: remoção de elemento (idioma Go)                    |
| activity4.05   | `activity4.05/` | Map: struct como chave, padrão set com `struct{}`         |
| activity4.06   | `activity4.06/` | Type switch sem binding, `[]interface{}`                  |
