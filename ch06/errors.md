# Errors in Go - Resumo de Estudo (Chapter 06)

## Visao Geral

Go trata erros como valores, nao como excecoes. Diferente de linguagens como Java ou Python que usam try/catch, Go adota uma abordagem explicita onde funcoes retornam erros como parte de seus valores de retorno e o chamador e responsavel por verifica-los. Essa filosofia torna o fluxo de erro visivel e previsivel no codigo.

---

## Tipos de Erros

### Syntax Errors (Erros de Sintaxe)

Erros de sintaxe sao detectados pelo compilador antes da execucao do programa. Ocorrem quando o codigo viola as regras gramaticais da linguagem Go. Exemplos incluem parenteses nao fechados, ponto e virgula faltando, ou uso incorreto de palavras-chave. Como Go e uma linguagem compilada, esses erros impedem completamente a geracao do binario. O compilador indica a linha e a natureza do problema, facilitando a correcao. Um exemplo classico e tentar acessar um tipo nao exportado de outro pacote, como `errors.errorString{}`, que resulta em erro de compilacao pois `errorString` comeca com letra minuscula e e privado ao pacote `errors`.

### Runtime Errors (Erros de Tempo de Execucao)

Erros de runtime ocorrem durante a execucao do programa, quando uma operacao invalida e tentada. O programa compila sem problemas, mas falha ao rodar. O exemplo mais comum e o acesso a um indice fora dos limites de um slice ou array. Se um slice tem 3 elementos (indices 0 a 2) e o codigo tenta acessar o indice 10, o programa entra em panico (panic) com a mensagem "index out of range". Esses erros sao particularmente perigosos porque so se manifestam em tempo de execucao, possivelmente em cenarios especificos que nao foram testados. A solucao e usar construcoes seguras como `for i := range slice` ao inves de indices manuais, e implementar verificacoes de limites quando necessario.

### Semantic Errors (Erros Semanticos / Erros de Logica)

Erros semanticos sao os mais sutis. O codigo compila e executa sem crashar, mas produz resultados incorretos. O programa faz algo diferente do que o programador pretendia. Um exemplo classico e usar `>` quando deveria ser `>=` em uma condicao. Se a intencao e que o valor 2 satisfaca a condicao, usar `km > 2` esta errado pois 2 nao e maior que 2, devendo ser `km >= 2`. Esses erros nao geram mensagens de erro; o programa simplesmente se comporta de forma inesperada. Sao os mais dificeis de detectar porque exigem que o programador compare o comportamento observado com o comportamento esperado. Testes unitarios e revisao de codigo sao as principais ferramentas para encontra-los.

---

## Error Interface Type (Tipo Interface error)

Em Go, `error` e uma interface nativa da linguagem definida como:

```
type error interface {
    Error() string
}
```

Qualquer tipo que implemente o metodo `Error() string` satisfaz a interface `error`. Isso significa que erros em Go sao simplesmente valores que podem ser passados, armazenados e comparados como qualquer outro valor.

A funcao `errors.New()` e a forma mais simples de criar um erro. Ela recebe uma string descritiva e retorna um valor que implementa a interface `error`. O tipo concreto retornado e `*errors.errorString`, que e um tipo privado do pacote `errors`, inacessivel diretamente por outros pacotes.

### Sentinel Errors (Erros Sentinela)

Uma pratica idiomatica em Go e definir erros sentinela como variaveis no nivel do pacote, usando a convencao de prefixo `Err`. Por exemplo, `ErrInvalidLastName` ou `ErrHourlyRate`. Esses erros podem ser comparados por identidade, permitindo que o chamador saiba exatamente qual tipo de falha ocorreu e tome a acao apropriada.

---

## Error Handling (Tratamento de Erros)

### Padrao Idiomatico

O padrao fundamental de tratamento de erros em Go e o retorno multiplo: uma funcao retorna o resultado desejado junto com um valor `error`. O chamador verifica se o erro e `nil`. Se nao for `nil`, houve uma falha e o erro deve ser tratado.

Esse padrao e onipresente na biblioteca padrao. Funcoes como `strconv.Atoi()` retornam `(int, error)`, e o chamador sempre deve verificar o erro antes de usar o resultado.

### Validacao com Erros

Funcoes de validacao seguem o padrao de verificar condicoes e retornar erros especificos para cada tipo de falha. Por exemplo, uma funcao que calcula pagamento pode validar a taxa horaria e as horas trabalhadas, retornando erros sentinela diferentes para cada violacao. Quando tudo esta valido, retorna `nil` como erro, indicando sucesso.

### Metodos com Erros em Structs

Metodos em structs podem implementar validacao retornando erros. Cada metodo valida um aspecto dos dados do struct e retorna um erro apropriado quando a validacao falha. Isso permite compor multiplas verificacoes, chamando cada validador e verificando o erro retornado sequencialmente.

---

## Error Wrapping (Encapsulamento de Erros)

Go permite encapsular erros para adicionar contexto conforme eles propagam pela pilha de chamadas. Usando `fmt.Errorf()` com o verbo `%w`, e possivel envolver um erro original com informacao adicional sobre onde e por que a falha ocorreu.

Quando um erro propaga por varias camadas (por exemplo, `readFile` -> `loadConfig` -> `startServer`), cada camada adiciona seu contexto. O resultado e uma cadeia de erros como: `startServer: loadConfig: readFile: open config.json: no such file or directory`.

A funcao `errors.Is()` permite inspecionar toda a cadeia de erros encapsulados para verificar se um erro especifico esta presente em qualquer nivel. Mesmo apos tres camadas de encapsulamento, `errors.Is(err, os.ErrNotExist)` consegue identificar que o erro original foi um arquivo nao encontrado. Isso e fundamental para tomar decisoes baseadas no tipo de erro sem perder o contexto adicionado por cada camada.

---

## Panic

`panic` e um mecanismo que interrompe imediatamente o fluxo normal de execucao de uma funcao. Quando `panic` e chamado, a funcao atual para de executar, todas as funcoes `defer` pendentes na goroutine sao executadas em ordem LIFO (Last In, First Out), e entao o programa termina com uma mensagem de erro e stack trace.

`panic` aceita qualquer valor como argumento, mas tipicamente recebe um erro ou uma string descritiva. Apos o panic, nenhuma linha de codigo subsequente na funcao (ou nas funcoes chamadoras) e executada, a menos que haja um `recover`.

### Panic e Defer

Uma caracteristica importante e que funcoes registradas com `defer` sao executadas mesmo durante um panic. Se uma funcao A chama uma funcao B que entra em panico, primeiro as funcoes defer de B sao executadas, depois as de A, e entao o programa termina. Isso garante que operacoes de limpeza (fechar arquivos, liberar recursos) acontecam mesmo em cenarios de falha catastrofica.

### Quando Usar Panic

`panic` deve ser reservado para situacoes verdadeiramente excepcionais e irrecuperaveis, como violacoes de invariantes internas do programa ou estados que nunca deveriam ocorrer. Para erros esperados (entrada invalida, arquivo nao encontrado, falha de rede), o padrao correto e retornar erros. Usar panic para validacao de entrada e considerado uma ma pratica, pois mesclar estrategias de tratamento de erro (retorno de erro vs. panic) torna o codigo inconsistente e imprevisivel.

---

## Recover

`recover` e a contraparte do `panic`. E uma funcao nativa que captura um panic em andamento, impedindo que o programa termine. `recover` so funciona quando chamado dentro de uma funcao `defer`. Se nao houver panic em andamento, `recover` retorna `nil`.

### Mecanica do Recover

Quando `recover` captura um panic, ele retorna o valor passado ao `panic`. A funcao que entrou em panico retorna ao seu chamador como se tivesse retornado normalmente (com valores zero para os retornos). A execucao continua a partir do ponto seguinte a chamada da funcao que panicou.

### Recover com Named Return Values

Uma tecnica poderosa e combinar `recover` com valores de retorno nomeados (named return values). Ao declarar uma funcao com retorno nomeado `(err error)`, uma funcao defer com recover pode atribuir o valor recuperado a `err`, convertendo efetivamente um panic em um retorno de erro. Isso permite que o chamador trate a situacao como um erro normal, sem saber que internamente houve um panic. Essa tecnica e util para criar APIs seguras que nao propagam panics para o codigo cliente.

### Recover Seletivo

O valor retornado por `recover` pode ser inspecionado para determinar qual tipo de panic ocorreu. Comparando o valor recuperado com erros sentinela conhecidos, e possivel tomar acoes diferentes dependendo da natureza do problema. Isso permite um tratamento granular de diferentes cenarios de falha, mantendo a capacidade de continuar a execucao do programa.

---

## Resumo das Boas Praticas

- Erros sao valores: trate-os como tal, passando-os, comparando-os e compondo-os
- Sempre verifique erros retornados: ignorar um retorno de erro e uma fonte comum de bugs
- Adicione contexto ao propagar erros: use error wrapping com `%w` para manter a cadeia de contexto
- Use `errors.Is()` para inspecionar erros encapsulados: nao dependa de comparacao de strings
- Prefira retornar erros ao inves de usar panic: panic e para o inesperado, erros sao para o esperado
- Use recover com cautela: apenas em limites de API ou para evitar que uma goroutine derrube o programa
- Defina erros sentinela no nivel do pacote: com prefixo `Err` para consistencia e reutilizacao
- Nao misture estrategias: seja consistente entre retorno de erro e panic dentro do mesmo modulo
