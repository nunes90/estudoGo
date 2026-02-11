# Estruturas de Coleção em Go

Resumo baseado nos exercícios e atividades do Capítulo 4.

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

### 3.5 Funções Importantes

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

Go permite criar tipos nomeados baseados em tipos primitivos:

```go
type id string

var id1 id                          // zero-value: ""
var id2 id = "1234-5678"            // inicialização com valor

// Comparação entre valores do mesmo tipo customizado
fmt.Println(id1 == id2)             // false

// Conversão explícita para o tipo base
fmt.Println(string(id2) == "1234-5678")  // true
```

Tipos customizados são úteis para dar **semântica** a valores e evitar confusão entre tipos que possuem a mesma representação base.

---

## 5. Comparativo Rápido

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
