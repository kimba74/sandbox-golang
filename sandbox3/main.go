package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// Create a new logger to be used in the sandbox3 application
//                     +---------------------------------- Writer for the log output
//                     |            +--------------------- Prefix to use fo this logger's messages
//                     |            |             +------- Format for additional log info (e.g. date, time, file)
//                     |            |             |
var logger = log.New(os.Stdout, "greeter: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)

func main() {
	// Configure default/static logger
	log.SetPrefix("startup: ")                                               // Set prefix to use
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile) // Set format of additional info
	// Logging that that the program execution has started
	log.Println("--== Let's get started ==--")
	// Initialize the Gorilla Web Toolkit and configure the handlers
	handler := &eventHandler{}
	practicerouter := mux.NewRouter().PathPrefix("/greeter").Subrouter()
	practicerouter.Methods("GET").Path("/hello").HandlerFunc(handler.handleGenericHello)
	practicerouter.Methods("GET").Path("/hello/{name}").HandlerFunc(handler.handleSpecificHello)
	// Start a HTTP listener on port 12345 and configure the Gorilla Web Toolkit router
	// as the main router for incoming traffic.
	log.Fatal(http.ListenAndServe(":12345", practicerouter))
}
