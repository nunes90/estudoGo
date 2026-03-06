package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan struct{})
	signal.Notify(sigs, syscall.SIGINT)
	go func() {
		for {
			s := <-sigs
			switch s {
			case syscall.SIGINT:
				fmt.Println("My process has been interrupted. Someone might of pressed CTRL-C")
				fmt.Println("Some clean up is occuring")
				done <- struct{}{}
			}
		}
	}()
	fmt.Println("Program is blocked until a signal is caught")
	done <- struct{}{}
	fmt.Println("Out of here")
}

/*
• Define a channel to send signals
• Define a channel that we can use as a flag to stop the execution
• Use Notify to send a SIGINT signal
• Create a goroutine that listens indefinitely to signals and if the signal is SIGINT, it does some
printouts and sends a message to the done channel with the true value
• Print a message stating we are waiting for the done message to be received
• Wait for the done message
• Print the final message
*/
