package main

import (
	"fmt"
	"log"
	"net/http"
)

const homeTemplate = `<!DOCTYPE html>
<html>
<head>
<title>Home</title>
</head>
<body>
<h1>Pick A Motorcycle</h1>
<a href="/kawasaki">Kawasaki</a>
<a href="/honda">Honda</a>
<a href="/suzuki">Suzuki</a>
</body>
</html>`

type motorcycle struct {
	Brand string
	Color string
}

func main() {
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/kawasaki", handleKawasaki)
	http.HandleFunc("/honda", handleHonda)
	http.HandleFunc("/suzuki", handleSuzuki)
	log.Println("Listening for Request on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	log.Println("A request at / came through")
	fmt.Fprintf(w, homeTemplate)
}

func handleKawasaki(w http.ResponseWriter, r *http.Request) {
	log.Println("Handle Kawasaki Ran")
	var m motorcycle
	m.Brand = "Kawasaki"
	m.Color = "Green"

	fmt.Fprintf(w, `<h1>Brand: %s<h1><p>Color: %s<p><a href="/">Return Home</a>`, m.Brand, m.Color)
}

func handleHonda(w http.ResponseWriter, r *http.Request) {
	log.Println("Handle Honda Ran")
	var m motorcycle
	m.Brand = "Honda"
	m.Color = "Red"

	fmt.Fprintf(w, `<h1>Brand: %s<h1><p>Color: %s<p><a href="/">Return Home</a>`, m.Brand, m.Color)
}

func handleSuzuki(w http.ResponseWriter, r *http.Request) {
	log.Println("Handle Suzuki Ran")
	var m motorcycle
	m.Brand = "Suzuki"
	m.Color = "Blue"

	fmt.Fprintf(w, `<h1>Brand: %s<h1><p>Color: %s<p><a href="/">Return Home</a>`, m.Brand, m.Color)
}
