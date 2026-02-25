# Exercícios Práticos de Go — Capítulos 1 ao 7

> **Livro:** Go Programming: From Beginner to Professional — Samantha Coyle
> **Objetivo:** Reforçar e combinar os conceitos vistos nos capítulos 1 a 7 antes de iniciar o Capítulo 8 (Generics).
> Cada exercício indica quais capítulos são revisitados. Tente resolver antes de ver as dicas!

---

## 📦 Como organizar

Crie uma pasta `pratica/` no seu repositório com subpastas por exercício:

```
estudoGo/
└── pratica/
    ├── ex01_inventario/
    ├── ex02_calculadora/
    ├── ex03_banco/
    ├── ex04_biblioteca/
    ├── ex05_agenda/
    └── ex06_shapes/
```

Cada pasta deve ter seu próprio `go.mod` e `main.go`.

---

## Exercício 1 — Inventário de Produtos

> **Caps revisitados:** 1 (variáveis, iota), 2 (loops, condicionais), 3 (strings), 4 (maps, slices)

### Enunciado

Crie um programa que gerencie um inventário simples de produtos. O programa deve:

1. Definir uma constante de categoria usando `iota` com pelo menos 3 categorias (ex: `Eletrônico`, `Roupa`, `Alimento`)
2. Criar um `map` onde a chave é o nome do produto (`string`) e o valor é uma `struct` com:
   - `Preco float64`
   - `Quantidade int`
   - `Categoria` (use o tipo com iota)
3. Pré-popular o mapa com ao menos 5 produtos
4. Iterar sobre o mapa e imprimir todos os produtos formatados
5. Calcular e imprimir o valor total do estoque (preço × quantidade)
6. Filtrar e imprimir apenas os produtos com quantidade menor que 3 (estoque baixo)

### Exemplo de saída esperada

```
=== Inventário ===
Notebook     | Eletrônico | R$ 3500.00 | Qtd: 2
Camiseta     | Roupa      | R$   89.90 | Qtd: 15
...

Valor total em estoque: R$ 12.450,50

=== Estoque Baixo ===
Notebook  (2 unidades)
```

### Dicas

<details>
<summary>Ver dicas</summary>

- Use `const` com `iota` para as categorias
- Use `fmt.Printf` com `%-15s` para alinhar as colunas
- Percorra o mapa com `for k, v := range produtos`
- Acumule o total com `total += v.Preco * float64(v.Quantidade)`

</details>

---

## Exercício 2 — Calculadora com Tratamento de Erros

> **Caps revisitados:** 2 (switch, condicionais), 5 (funções), 6 (erros)

### Enunciado

Implemente uma calculadora que receba dois números e uma operação via argumentos do programa (`os.Args`) e retorne o resultado. O programa deve:

1. Criar uma função para cada operação: `Somar`, `Subtrair`, `Multiplicar`, `Dividir`
2. Cada função deve retornar `(float64, error)`
3. A função `Dividir` deve retornar um erro customizado se o divisor for zero
4. Criar uma função `Calcular(a, b float64, op string) (float64, error)` que usa `switch` para chamar a função correta
5. Retornar erro se a operação for desconhecida
6. No `main`, ler os argumentos, chamar `Calcular` e tratar o erro adequadamente

### Exemplo de uso

```bash
go run main.go 10 / 2   # Resultado: 5.00
go run main.go 10 / 0   # Erro: divisão por zero não é permitida
go run main.go 5 % 2    # Erro: operação '%' não suportada
```

### Dicas

<details>
<summary>Ver dicas</summary>

- Use `errors.New("mensagem")` para erros simples
- Use `fmt.Errorf("operação '%s' não suportada", op)` para erros com contexto
- Importe `"os"` e use `os.Args[1]`, `os.Args[2]`, `os.Args[3]`
- Use `strconv.ParseFloat(os.Args[1], 64)` para converter string em float

</details>

---

## Exercício 3 — Sistema Bancário Simples

> **Caps revisitados:** 1 (variáveis, tipos), 4 (slices), 5 (funções, variadic), 6 (erros), 7 (interfaces)

### Enunciado

Crie um mini sistema bancário com as seguintes regras:

1. Defina uma interface `ContaBancaria` com os métodos:
   - `Depositar(valor float64) error`
   - `Sacar(valor float64) error`
   - `Saldo() float64`
   - `Extrato() []string`

2. Implemente dois tipos que satisfaçam essa interface:
   - `ContaCorrente`: permite saldo negativo até um limite de `-500.00`
   - `ContaPoupanca`: não permite saldo negativo (erro se tentar sacar mais do que tem)

3. Crie uma função `RealizarTransferencia(origem, destino ContaBancaria, valor float64) error` que:
   - Saca da conta de origem
   - Deposita na conta de destino
   - Desfaz o saque se o depósito falhar

4. No `main`, crie uma conta de cada tipo, faça algumas operações e imprima o extrato

### Exemplo de saída esperada

```
=== Conta Corrente (Lucas) ===
Depósito: +R$ 1000.00
Saque:    -R$ 200.00
Saldo atual: R$ 800.00

=== Transferência ===
Transferência de R$ 300.00 realizada com sucesso

=== Extrato — Conta Poupança (Maria) ===
Depósito: +R$ 300.00
Saldo atual: R$ 300.00
```

### Dicas

<details>
<summary>Ver dicas</summary>

- O campo `extrato []string` pode ser um slice que você vai fazendo `append` a cada operação
- Use `fmt.Sprintf("Depósito: +R$ %.2f", valor)` para formatar as entradas do extrato
- A função de transferência deve ser do tipo que aceite a interface, não os tipos concretos

</details>

---

## Exercício 4 — Biblioteca de Livros

> **Caps revisitados:** 3 (strings), 4 (slices, maps), 5 (funções variadic), 7 (interfaces)

### Enunciado

Implemente um gerenciador de biblioteca:

1. Crie uma interface `Pesquisavel` com o método `Contem(termo string) bool`

2. Crie um tipo `Livro` com os campos: `Titulo`, `Autor`, `Genero string` e `AnoPublicacao int`

3. Implemente `Contem` em `Livro` — retorna `true` se o termo aparece no título ou no nome do autor (case-insensitive, use `strings.ToLower`)

4. Crie um tipo `Biblioteca` que é um slice de `Livro` e implemente:
   - `Adicionar(livros ...Livro)` — variadic, adiciona um ou mais livros
   - `Pesquisar(termo string) []Livro` — retorna todos que satisfazem `Contem`
   - `PorGenero(genero string) []Livro` — filtra por gênero
   - `MaisRecentes(n int) []Livro` — retorna os n livros com maior ano

5. No `main`, popule a biblioteca com ao menos 6 livros e demonstre todas as funcionalidades

### Dicas

<details>
<summary>Ver dicas</summary>

- Para `MaisRecentes`, ordene o slice com `sort.Slice` antes de retornar os primeiros `n`
- Use `strings.Contains(strings.ToLower(b.Titulo), strings.ToLower(termo))`
- A função variadic `Adicionar(livros ...Livro)` permite chamar `bib.Adicionar(l1, l2, l3)`

</details>

---

## Exercício 5 — Agenda de Contatos com Validação

> **Caps revisitados:** 1, 3, 4, 5, 6 (erros customizados com tipo próprio)

### Enunciado

Crie uma agenda de contatos com validação robusta:

1. Defina um tipo de erro customizado `ErrValidacao` com os campos `Campo string` e `Mensagem string`, e implemente o método `Error() string`

2. Crie um tipo `Contato` com: `Nome`, `Email`, `Telefone string`

3. Implemente uma função `ValidarContato(c Contato) []error` que retorna uma lista de erros de validação:
   - Nome não pode estar vazio
   - Email deve conter `@` e `.`
   - Telefone deve ter entre 8 e 15 caracteres (apenas dígitos)

4. Crie um tipo `Agenda` (map de string para Contato) com os métodos:
   - `Adicionar(c Contato) error` — valida antes de adicionar, retorna o primeiro erro encontrado
   - `Buscar(nome string) (Contato, error)` — retorna erro se não encontrar
   - `Listar() []Contato` — retorna todos os contatos ordenados por nome

5. No `main`, tente adicionar contatos válidos e inválidos, mostrando as mensagens de erro

### Exemplo de saída esperada

```
Erro ao adicionar: campo 'Email' inválido: deve conter '@' e '.'
Erro ao adicionar: campo 'Telefone' inválido: deve ter entre 8 e 15 dígitos
Contato 'Lucas Nunes' adicionado com sucesso!
```

### Dicas

<details>
<summary>Ver dicas</summary>

- `ErrValidacao` deve implementar a interface `error` com o método `Error() string`
- Use `unicode.IsDigit(r)` dentro de um loop para checar se são só dígitos
- Para ordenar, use `sort.Slice(lista, func(i, j int) bool { return lista[i].Nome < lista[j].Nome })`

</details>

---

## Exercício 6 — Calculadora de Formas Geométricas (Desafio)

> **Caps revisitados:** 1, 2, 5, 6, 7 (interfaces — foco principal)

### Enunciado

Este exercício foca em interfaces e polimorfismo:

1. Crie uma interface `Forma` com os métodos:
   - `Area() float64`
   - `Perimetro() float64`
   - `Descricao() string`

2. Implemente pelo menos 4 tipos que satisfaçam `Forma`:
   - `Circulo` (campo `Raio float64`)
   - `Retangulo` (campos `Largura, Altura float64`)
   - `Triangulo` (campos `Base, Altura, LadoA, LadoB, LadoC float64`)
   - `Quadrado` (campo `Lado float64`)

3. Crie uma função `ImprimirInfo(f Forma)` que imprime área, perímetro e descrição

4. Crie uma função `MaiorArea(formas []Forma) Forma` que retorna a forma com maior área

5. Crie uma função `TotalArea(formas []Forma) float64` que soma todas as áreas

6. No `main`, crie um slice com instâncias de cada forma e use todas as funções

### Fórmulas

| Forma | Área | Perímetro |
|-------|------|-----------|
| Círculo | π × r² | 2 × π × r |
| Retângulo | l × a | 2 × (l + a) |
| Triângulo | (b × h) / 2 | a + b + c |
| Quadrado | l² | 4 × l |

### Dicas

<details>
<summary>Ver dicas</summary>

- Use `math.Pi` para π e `math.Pow(r, 2)` para r²
- `MaiorArea` pode inicializar com `formas[0]` e iterar a partir do índice 1
- Lembre-se: `Quadrado` pode ser implementado reaproveitando as fórmulas de `Retangulo` com `Largura == Altura`

</details>

---

## 🎯 Desafio Final — Mini Sistema Integrado

> **Todos os caps: 1–7**

Combine os exercícios anteriores criando um programa que:

1. Usa a `Biblioteca` do Ex.4 como base de dados
2. Adiciona uma interface `Relatorio` com o método `GerarRelatorio() string`
3. Implementa `Relatorio` tanto em `Biblioteca` quanto em `Agenda`
4. Trata todos os erros com tipos customizados (Ex.5)
5. Exibe um menu no terminal com as opções de gerenciar livros ou contatos

**Esse desafio não tem resposta — é seu projeto pessoal para o repositório!** 🚀

---

## 📚 Referência Rápida — Conceitos por Capítulo

| Capítulo | Tópico Principal | Palavras-chave |
|----------|-----------------|----------------|
| 1 | Variáveis e Tipos | `var`, `const`, `iota`, `:=` |
| 2 | Controle de Fluxo | `if`, `for`, `switch`, `break` |
| 3 | Tipos Primitivos | `string`, `int`, `float64`, `bool`, `rune` |
| 4 | Tipos Complexos | `array`, `slice`, `map`, `struct` |
| 5 | Funções | `func`, múltiplos retornos, variadic, closures |
| 6 | Tratamento de Erros | `error`, `errors.New`, `fmt.Errorf`, tipo de erro customizado |
| 7 | Interfaces | `interface`, polimorfismo, `type assertion` |

---

> **Próximo passo:** Capítulo 8 — Generics 🎉
> Após concluir esses exercícios, você estará bem preparado para aprender como tornar suas funções e tipos reutilizáveis com parâmetros de tipo genérico!
