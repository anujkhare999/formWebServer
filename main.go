package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "wrong url entered", http.StatusBadRequest)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/form" {
		http.Error(w, "wrong url entered", http.StatusBadRequest)
		return
	}

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parse form error: %v\n", err)
		return
	}

	fmt.Fprintf(w, "Post request successful\n")
	Name := r.FormValue("name")
	Address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", Name)
	fmt.Fprintf(w, "Addres = %s\n", Address)
}

func main() {
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
