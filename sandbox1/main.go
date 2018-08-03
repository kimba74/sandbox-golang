// Channel declaration
//    chan    - Send/Receive Channel
//    <-chan  - Receive Only Channel
//    chan<-  - Send Only Channel
//
// Example:
//   chan string    - Can be used for sending and receiving strings
//   <-chan string  - Can be used for receiving strings
//   chan<- string  - Can be used for sending strings

package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var logMainProg = log.New(os.Stdout, "[--main--] ", log.Ldate|log.Ltime)
var logConsumer = log.New(os.Stdout, "[consumer] ", log.Ldate|log.Ltime)
var logProducer = log.New(os.Stdout, "[producer] ", log.Ldate|log.Ltime)

// producer function creates test strings and posts them to the provided channel.
// The channel passed to the function is a write-only channel.
func producer(channel chan<- string) {
	logProducer.Println("Started")      // TODO: Research how to check if channel closed
	defer logProducer.Println("Exited") //       (1) It looks like a channel can't be
	i := 0                              //           tested for state best practice
	for {                               //           recommends the writer closes the
		i++                                  //           channel once done writing
		channel <- fmt.Sprintf("Test-%d", i) //
		time.Sleep(1 * time.Second)
	}
}

// consumer function consumes the string messages sent through the provided channel.
// The channel passed to the function is a read-only channel.
func consumer(channel <-chan string) {
	logConsumer.Println("Started")
	defer logConsumer.Println("Exited")
	val, ok := <-channel
	for ok {
		logConsumer.Printf("Read: \"%s\"\n", val)
		val, ok = <-channel
	}
}

func main() {
	// Create channels used for this sandbox
	com := make(chan string, 2)
	sig := make(chan os.Signal)
	// Register the 'sig' channel for observing signals for the SIGTERM (15) occurence
	signal.Notify(sig, syscall.SIGTERM)
	// Run the producer and consumer of this example
	go producer(com)
	go consumer(com)
	// Performing program shutdown after receiving SIGTERM
	<-sig
	logMainProg.Println("Received SIGTERM.....shutting down") // NOTE: not shutting down cleanly
	close(com)                                                //       receiving a panic() after
	close(sig)                                                //       channels are closed. It
	time.Sleep(1 * time.Second)                               //       looks like as if the producer()
	logMainProg.Println("Shutdown process complete!")         //       tries to send on a closed channel.
}
