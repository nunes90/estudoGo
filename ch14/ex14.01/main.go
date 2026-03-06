// ex14.01 - simulating a cleanup
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	sigs := make(chan os.Signal, 1)

	done := make(chan struct{})

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTSTP)

	go func() {
		for {
			s := <-sigs

			switch s {
			case syscall.SIGINT:
				fmt.Println()
				fmt.Println("My process has been interrupted. Someone might have pressed CTRL-C")
				fmt.Println("Some clean up is occuring")
				cleanup()
				done <- struct{}{}
			case syscall.SIGTSTP:
				fmt.Println()
				fmt.Println("Someone pressed CTRL-Z")
				fmt.Println("Some clean up is occuring")
				cleanup()
				done <- struct{}{}
			}
		}
	}()
	fmt.Println("Program is blocked until a signal is caught(ctrl-z, ctrl-c)")
	done <- struct{}{}
	fmt.Println("Out of here")
}

func cleanup() {
	fmt.Println("Simulating clean up")
	for i := 0; i <= 10; i++ {
		fmt.Println("Deleting Files.. Not really.", i)
		time.Sleep(1 * time.Second)
	}
}
