package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	
	r := mux.NewRouter()
	r.HandleFunc("/", Output)
	http.Handle("/", r)
	


	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}

// Output generates numbers randomly
func Output(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	
	fmt.Fprintf(w, " \n Number : %v", randomNumbers())
	fmt.Fprintf(w, " \n \n Version 1.0.0")
}

func randomNumbers() int {
	return rand.Intn(5)
}
