package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type eventHandler struct{}

func (h *eventHandler) handleGenericHello(response http.ResponseWriter, request *http.Request) {
	logger.Println("GET /hello -- received")
	fmt.Fprintln(response, "Hi there!!!")
}

func (h *eventHandler) handleSpecificHello(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	logger.Printf("GET /hello/%s -- received\n", vars["name"])
	response.WriteHeader(http.StatusOK)
	fmt.Fprintf(response, "Hi there, %s!!!\n", vars["name"])
}
