package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8000\n")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func helloHandler(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/hello" {
		http.Error(writer, "404", http.StatusNotFound)
		return
	}

	if request.Method != "GET" {
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(writer, "hello!")

}

func formHandler(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Fprintf(writer, "Form Submitted!")
	name := request.FormValue("name")
	address := request.FormValue("address")
	fmt.Fprintf(writer, "Name = %s\n", name)
	fmt.Fprintf(writer, "Address = %s\n", address)

}
