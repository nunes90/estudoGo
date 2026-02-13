// activity4.05 - Creating a locale checker
package main

import (
	"fmt"
	"os"
	"strings"
)

type locale struct {
	language string
	country  string
}

func getLocales() map[locale]struct{} {
	supportedLocales := make(map[locale]struct{}, 5)
	// `struct{}{}` - o `struct{}` e o tipo (zero bytes),
	// o `{}` externo e a inicialização do valor (um struct vazio)
	supportedLocales[locale{"en", "US"}] = struct{}{}
	supportedLocales[locale{"en", "CN"}] = struct{}{}
	supportedLocales[locale{"fr", "CN"}] = struct{}{}
	supportedLocales[locale{"fr", "FR"}] = struct{}{}
	supportedLocales[locale{"ru", "RU"}] = struct{}{}
	return supportedLocales
}

func localeExists(l locale) bool {
	_, exists := getLocales()[l]
	return exists
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No locale provided")
		os.Exit(1)
	}

	localeParts := strings.Split(os.Args[1], "_")
	if len(localeParts) != 2 {
		fmt.Printf("Invalid locale format: %v\n", os.Args[1])
		os.Exit(1)
	}

	passedLocale := locale{
		language: localeParts[0],
		country:  localeParts[1],
	}

	if !localeExists(passedLocale) {
		fmt.Printf("Locale not supported: %v\n", os.Args[1])
		os.Exit(1)
	}
	fmt.Println("Locale passed is supported")
}
