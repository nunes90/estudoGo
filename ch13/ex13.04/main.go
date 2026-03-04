// ex13.04 - creating a stopwatch with a time limit
package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	timeLimit := 5 * time.Second
	fmt.Println("Press Enter to start the stopwatch...")
	_, err := fmt.Scanln() // Wait for user to press Enter
	if err != nil {
		fmt.Println("Error reading from stdin:", err)
		return
	}
	fmt.Println("Stopwatch started. Waiting for", timeLimit)

	time.Sleep(timeLimit)
	fmt.Println("Time's up! Executing the other command.")
	cmd := exec.Command("echo", "Hello")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error executing command:", err)
	}

}
