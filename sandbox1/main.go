package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

var logMainProg = log.New(os.Stdout, "[--main--] ", log.Ldate|log.Ltime)
var logConsumer = log.New(os.Stdout, "[consumer] ", log.Ldate|log.Ltime)
var logProducer = log.New(os.Stdout, "[producer] ", log.Ldate|log.Ltime)

func producer(min, max int, channel chan string) {
	logProducer.Printf("Started (min=%d, max=%d)\n", min, max)
	defer logProducer.Println("Exited")
	for i := min; i <= max; i++ {
		channel <- fmt.Sprintf("Test-%d", i)
		time.Sleep(1 * time.Second)
	}
	close(channel)
}

func consumer(channel chan string) {
	logConsumer.Println("Started")
	defer logConsumer.Println("Exited")
	val, ok := <-channel
	for ok {
		logConsumer.Printf("Read: \"%s\"\n", val)
		val, ok = <-channel
	}
}

func main() {
	ch := make(chan string, 2)
	go producer(0, 30, ch)
	consumer(ch)
}
