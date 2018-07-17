package main

import (
	"flag"
	"fmt"
)

var (
	hostName string
	hostIP   string
)

func main() {
	// Parse the flags passed to the application
	flag.Parse()
	// Print out the flags received
	fmt.Printf("Flag 'hostName': %s\n", hostName)
	fmt.Printf("Flag 'hostIP'  : %s\n", hostIP)
}

// init will be executed before the main method
func init() {
	flag.StringVar(&hostName, "hostName", "localhost", "This is the 'Host Name' flag")
	flag.StringVar(&hostIP, "hostIP", "127.0.0.1", "This is the 'Host IP' flag")
}
