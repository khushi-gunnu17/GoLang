package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {	

	fmt.Println("-----------------mod in golang--------------")

	greeter()

	// for routing
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	// running a server 
	// log.FataL = to handle the error 
	log.Fatal(http.ListenAndServe(":4000", r))

}


func greeter()  {
	
	fmt.Println("Hey there mod users.")

}


func serveHome(w http.ResponseWriter, r *http.Request)  {
	
	w.Write([]byte("<h1>Welcome to golang.</h1>"))

}