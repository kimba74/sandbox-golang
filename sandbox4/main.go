package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Create a channel with size 1 for capturing signals
	sig := make(chan os.Signal, 1)

	// Configure 'signal' to forward SIGINT and SIGIO to the afore create channel
	// If no signal is configured then all signals are captured and placed in the channel
	signal.Notify(sig, syscall.SIGINT, syscall.SIGIO)

	// Wait until the channel has a signal then assign it to 's'
	s := <-sig

	// Evaluate which signal was captured
	switch s {
	case syscall.SIGINT:
		fmt.Println("SIGINT received")
	case syscall.SIGIO:
		fmt.Println("SIGIO received")
	default:
		fmt.Printf("Signal received: %s\n", s)
	}
}
