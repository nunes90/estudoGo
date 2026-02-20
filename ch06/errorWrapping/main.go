package main

import (
	"errors"
	"fmt"
	"os"
)

func readFile(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("readFile: %w", err)
	}
	return data, nil
}

func loadConfig() error {
	_, err := readFile("config.json")
	if err != nil {
		return fmt.Errorf("loadConfig: %w", err)
	}
	return nil
}

func startServer() error {
	err := loadConfig()
	if err != nil {
		return fmt.Errorf("startServer: %w", err)
	}
	return nil
}

func main() {
	err := startServer()
	if err != nil {
		fmt.Println(err)
		// startServer: loadConfig: readFile: open config.json: no such file or directory

		// mesmo com 3 camadas de wrapping, você consegue verificar o erro original:
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("O arquivo não existe!")
		}
	}
}
