package main

import (
	"log"
	"myapp/handlers"
	"net/http"
	"os"
)


func main() {
	// It register a function to a path called deafualt serve MUX
	// Mux choose what function will be executed according to the request
	// handle func takes function and creates a handler and add to MUX

	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)

	sm := http.NewServeMux()
	sm.Handle("/",hh)
	http.ListenAndServe(":9090",sm)
	 // TO serve explixity one ıp ıp:8080
	 // Second parameter is handler
	

}
