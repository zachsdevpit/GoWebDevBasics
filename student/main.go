package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", theFunctionName)
	http.HandleFunc("/whatever", anotherFunc)

	startServer()
}

func startServer() {
	log.Println("server has started listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func theFunctionName(theWriter http.ResponseWriter, theRequest *http.Request) {
	log.Println("a client request to / was made")
	fmt.Fprintf(theWriter, `<h1>Sup Foo</h1><a href="/whatever">Go to other path</a>`)
}

func anotherFunc(w http.ResponseWriter, r *http.Request) {
	log.Println("anotherFunc called cause request to /whatever")
	fmt.Fprintf(w, `<h1>The other page</h1><a href="/">Go Home</a>`)
}
