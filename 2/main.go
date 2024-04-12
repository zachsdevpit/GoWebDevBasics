package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleHome)

	log.Println("Listening for Request on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	log.Println("A request at / was made to the Server")
	fmt.Fprintf(w, "A Response was written to the Client")
}
