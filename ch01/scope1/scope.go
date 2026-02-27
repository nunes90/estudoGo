// Scope
/*
Regra geral

  Go sempre resolve um nome procurando do escopo mais interno
  para o mais externo: bloco → função → pacote. Usa o primeiro
  que encontrar.
*/
package main

import "fmt"

var level = "pkg"

func main() {
	fmt.Println("Main start :", level)
	if true {
		fmt.Println("Block start :", level)
		funcA()
	}
}

func funcA() {
	fmt.Println("funcA start :", level)
}
