package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var logger = log.New(os.Stdout, "greeter: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)

func ServerAPI() {
	handler := &eventHandler{}
	practicerouter := mux.NewRouter().PathPrefix("/greeter").Subrouter()
	practicerouter.Methods("GET").Path("/hello").HandlerFunc(handler.handleGenericHello)
	practicerouter.Methods("GET").Path("/hello/{name}").HandlerFunc(handler.handleSpecificHello)
	http.ListenAndServe(":12345", practicerouter)
}

func main() {
	log.SetPrefix("startup: ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
	log.Println("--== Let's get started ==--")
	handler := &eventHandler{}
	practicerouter := mux.NewRouter().PathPrefix("/greeter").Subrouter()
	practicerouter.Methods("GET").Path("/hello").HandlerFunc(handler.handleGenericHello)
	practicerouter.Methods("GET").Path("/hello/{name}").HandlerFunc(handler.handleSpecificHello)
	log.Fatal(http.ListenAndServe(":12345", practicerouter))

	// sig := make(chan os.Signal, 1)
	// signal.Notify(sig, syscall.SIGINT, syscall.SIGIO)
	//
	// s := <-sig
	//
	// switch s {
	// case syscall.SIGINT:
	// 	fmt.Println("SIGINT received")
	// case syscall.SIGIO:
	// 	fmt.Println("SIGIO received")
	// default:
	// 	fmt.Printf("Signal received: %s\n", s)
	// }
}
