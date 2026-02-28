# Solucao1: O jeito antigo

Você edita manualmente o go.mod do app dizendo "quando precisar de github.com/lucas/printer, use a pasta local ../printer":

```Go
// app/go.mod
module app

go 1.21

replace github.com/lucas/printer => ../printer
```

Funciona, mas tem problemas:

- Se você tem 5 módulos locais, precisa de 5 linhas replace em cada go.mod
- Precisa lembrar de remover os replace antes de fazer commit/publicar
- É fácil esquecer e quebrar o projeto
