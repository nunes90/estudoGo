# Errors em Go

## Visao Geral

Go adota uma filosofia unica para tratamento de erros: erros sao valores, nao excecoes. Diferente de linguagens como Java, Python ou C# que utilizam mecanismos de try/catch/throw, Go trata erros de forma explicita, onde funcoes retornam valores de erro como parte do seu retorno multiplo e o chamador e responsavel por verificar e tratar cada erro individualmente. Essa abordagem torna o fluxo de erro visivel, previsivel e rastreavel no codigo. Nao existe tratamento implicito ou propagacao automatica de erros — tudo e deliberado.

---

## Categorias de Erros

### Syntax Errors (Erros de Sintaxe)

Erros de sintaxe sao violacoes das regras gramaticais da linguagem, detectados pelo compilador antes de qualquer execucao. Como Go e uma linguagem compilada, o binario simplesmente nao e gerado enquanto houver erros de sintaxe. O compilador aponta a linha exata e a natureza do problema, tornando esses erros os mais faceis de identificar e corrigir. Exemplos incluem parenteses ou chaves nao fechados, uso incorreto de palavras-chave, declaracoes mal formadas e tentativas de acessar tipos ou membros nao exportados de outros pacotes (como um tipo que comeca com letra minuscula). Em Go, a regra de exportacao baseada em capitalizacao faz parte da sintaxe visivel: nomes que comecam com maiuscula sao publicos, com minuscula sao privados ao pacote.

### Runtime Errors (Erros de Tempo de Execucao)

Erros de runtime ocorrem quando o programa esta em execucao e tenta realizar uma operacao invalida. O codigo compila sem problemas, mas falha ao rodar. O caso mais classico e o acesso a um indice fora dos limites de um slice ou array, que causa um panic automatico com a mensagem "index out of range". Outros exemplos incluem desreferenciamento de ponteiro nulo, divisao por zero e envio para um canal fechado. Esses erros sao perigosos porque so se manifestam em tempo de execucao e podem depender de condicoes especificas que nao foram testadas. A mitigacao envolve uso de construcoes seguras como `range` para iteracao, verificacoes explicitas de limites, e cobertura de testes abrangente.

### Semantic Errors (Erros Semanticos / Erros de Logica)

Erros semanticos sao os mais insidiosos. O codigo compila, executa sem crashar, mas produz resultados incorretos. O programa faz algo diferente do que o programador pretendia. Um exemplo tipico e usar um operador de comparacao errado, como `>` quando deveria ser `>=`, fazendo com que um valor de borda seja excluido indevidamente. Esses erros nao geram nenhuma mensagem — o programa simplesmente se comporta de forma inesperada. Por isso, sao os mais dificeis de detectar, exigindo que o programador compare sistematicamente o comportamento observado com o esperado. As principais ferramentas para encontra-los sao testes unitarios, testes de borda, revisao de codigo e depuracao cuidadosa.

---

## Error Interface Type (Tipo Interface error)

No coracao do sistema de erros de Go esta a interface `error`, uma interface nativa da linguagem com um unico metodo: `Error() string`. Qualquer tipo que implemente esse metodo satisfaz automaticamente a interface `error`. Isso significa que erros em Go sao simplesmente valores — podem ser passados como argumentos, retornados de funcoes, armazenados em variaveis, comparados e compostos como qualquer outro valor da linguagem.

A forma mais simples de criar um erro e atraves da funcao `errors.New()`, que recebe uma string descritiva e retorna um valor implementando a interface `error`. O tipo concreto retornado e privado ao pacote `errors`, inacessivel diretamente. Para erros com formatacao dinamica, usa-se `fmt.Errorf()`, que permite interpolar valores na mensagem de erro.

### Sentinel Errors (Erros Sentinela)

Uma pratica idiomatica em Go e definir erros sentinela como variaveis no nivel do pacote, usando a convencao de prefixo `Err` (por exemplo, `ErrNotFound`, `ErrInvalidInput`). Esses erros sao valores unicos que podem ser comparados por identidade, permitindo que o chamador determine exatamente qual tipo de falha ocorreu e tome a acao apropriada. A biblioteca padrao de Go usa extensivamente esse padrao — `io.EOF`, `sql.ErrNoRows`, `os.ErrNotExist` sao exemplos conhecidos. Erros sentinela sao especialmente uteis quando diferentes falhas exigem tratamentos distintos.

---

## Error Handling (Tratamento de Erros)

### Padrao Idiomatico

O padrao fundamental de tratamento de erros em Go e o retorno multiplo: uma funcao retorna o resultado desejado junto com um valor `error`. O chamador verifica se o erro e `nil` (indicando sucesso) ou nao `nil` (indicando falha). Esse padrao e onipresente na biblioteca padrao e no ecossistema Go. A convencao e que o erro seja sempre o ultimo valor de retorno.

A verificacao de erros e feita imediatamente apos a chamada, tipicamente com `if err != nil`. Ignorar um erro retornado e considerado uma pratica perigosa e e uma fonte comum de bugs. O linter `errcheck` pode ser usado para detectar erros nao verificados.

### Validacao com Erros

Funcoes de validacao seguem naturalmente o padrao de retorno de erro. Cada condicao invalida resulta em um retorno antecipado com um erro especifico (frequentemente um sentinel error). Quando todas as validacoes passam, a funcao retorna `nil` como erro, sinalizando sucesso. Isso permite composicao clara: o chamador executa validadores sequencialmente, tratando o primeiro erro encontrado.

### Metodos com Erros em Structs

Metodos definidos em structs podem encapsular logica de validacao retornando erros. Cada metodo valida um aspecto especifico dos dados e retorna um erro apropriado quando a validacao falha. Esse padrao permite separar responsabilidades de validacao e compor verificacoes de forma modular.

---

## Error Wrapping (Encapsulamento de Erros)

Go permite encapsular erros para adicionar contexto conforme eles propagam pela pilha de chamadas. Usando `fmt.Errorf()` com o verbo `%w`, e possivel envolver um erro original com informacao adicional sobre onde e por que a falha ocorreu. Cada camada da aplicacao adiciona seu contexto, resultando em cadeias de erro descritivas que facilitam o diagnostico.

A funcao `errors.Is()` permite inspecionar toda a cadeia de erros encapsulados para verificar se um erro especifico esta presente em qualquer nivel de profundidade. Mesmo apos multiplas camadas de encapsulamento, e possivel identificar o erro original. Isso e fundamental para tomar decisoes baseadas no tipo de erro sem perder o contexto acumulado.

A funcao `errors.As()` complementa `errors.Is()`: enquanto `Is` verifica identidade, `As` verifica se algum erro na cadeia corresponde a um tipo especifico, permitindo acessar campos e metodos do erro concreto. Juntas, essas funcoes formam a base da inspecao de erros em Go moderno.

---

## Panic

`panic` e um mecanismo nativo que interrompe imediatamente o fluxo normal de execucao. Quando invocado, a funcao atual para de executar, todas as funcoes `defer` pendentes na goroutine sao executadas em ordem LIFO (Last In, First Out), e o programa termina com uma mensagem de erro e stack trace. `panic` aceita qualquer valor como argumento, mas tipicamente recebe um erro ou uma string descritiva.

### Panic e Defer

Uma propriedade fundamental e que funcoes registradas com `defer` sao garantidamente executadas mesmo durante um panic. Se a funcao A chama a funcao B que entra em panico, primeiro as funcoes defer de B sao executadas, depois as de A, e entao o programa termina. Isso garante que operacoes de limpeza — fechar arquivos, liberar conexoes, desbloquear mutexes — acontecam mesmo em cenarios de falha catastrofica. O `defer` e o mecanismo que torna o panic seguro do ponto de vista de gerenciamento de recursos.

### Quando Usar Panic

`panic` deve ser reservado para situacoes verdadeiramente excepcionais e irrecuperaveis: violacoes de invariantes internas, estados que nunca deveriam ocorrer, ou falhas durante a inicializacao do programa que tornam a continuacao impossivel. Para erros esperados (entrada invalida, arquivo nao encontrado, falha de rede, timeout), o padrao correto e retornar erros. Misturar estrategias de tratamento — retorno de erro em algumas funcoes e panic em outras para o mesmo tipo de situacao — torna o codigo inconsistente e imprevisivel. Uma excecao aceita e o uso de `panic` em funcoes auxiliares prefixadas com `Must` (como `template.Must` ou `regexp.MustCompile`), que sao chamadas tipicamente durante a inicializacao com valores conhecidos em tempo de compilacao.

---

## Recover

`recover` e a contraparte do `panic`. E uma funcao nativa que captura um panic em andamento, impedindo que o programa termine. `recover` so funciona quando chamado dentro de uma funcao `defer` — em qualquer outro contexto, retorna `nil` e nao tem efeito.

### Mecanica do Recover

Quando `recover` captura um panic, ele retorna o valor originalmente passado ao `panic`. A funcao que panicou retorna ao seu chamador como se tivesse terminado normalmente, com valores zero para seus retornos. A execucao continua a partir do ponto seguinte a chamada da funcao que panicou. Isso permite que o programa continue operando mesmo apos uma situacao excepcional.

### Recover com Named Return Values

Uma tecnica idiomatica e combinar `recover` com valores de retorno nomeados. Ao declarar uma funcao com retorno nomeado `(err error)`, uma funcao defer com recover pode atribuir o valor recuperado a `err`, convertendo efetivamente um panic em um retorno de erro convencional. O chamador recebe o erro normalmente, sem saber que internamente houve um panic. Essa tecnica e util para criar APIs seguras que isolam panics internos do codigo cliente.

### Recover Seletivo

O valor retornado por `recover` pode ser inspecionado para determinar a natureza do panic. Comparando o valor recuperado com erros sentinela ou verificando seu tipo, e possivel tomar acoes diferentes para cada cenario de falha. Panics nao reconhecidos podem ser re-panicados com `panic(r)` para nao mascarar problemas inesperados. Essa abordagem seletiva permite tratamento granular sem engolir erros silenciosamente.

---

## Boas Praticas

- Erros sao valores: trate-os como tal, passando, comparando e compondo
- Sempre verifique erros retornados: ignorar um retorno de erro e fonte comum de bugs
- Adicione contexto ao propagar erros: use error wrapping com `%w`
- Use `errors.Is()` e `errors.As()` para inspecionar erros encapsulados
- Prefira retornar erros ao inves de panic: panic e para o inesperado, erros sao para o esperado
- Use recover com cautela: apenas em limites de API, handlers HTTP ou para evitar que uma goroutine derrube o programa inteiro
- Defina erros sentinela no nivel do pacote com prefixo `Err`
- Nao misture estrategias: seja consistente entre retorno de erro e panic
- Nao use panic para validacao de entrada do usuario
- Faca recover seletivo: re-panice o que nao reconhece
