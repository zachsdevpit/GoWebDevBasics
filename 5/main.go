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

const motorcycleTemplate = `<!DOCTYPE html>
<html>
<head>
<title>Motorcycle</title>
</head>
<body>
<h1>%s</h1>
<p>%s</p>
<a href="/">Return Home</a>
</body>
</html>`

type motorcycle struct {
	Brand string
	Color string
}

func main() {
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/{motorcycle}", handleMotorcycle)
	log.Println("Listening for Request on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, homeTemplate)
}

func handleMotorcycle(w http.ResponseWriter, r *http.Request) {
	whichMotorcycle := r.PathValue("motorcycle")
	m := getMotorcycle(whichMotorcycle)
	fmt.Fprintf(w, motorcycleTemplate, m.Brand, m.Color)
}

func getMotorcycle(brand string) *motorcycle {
	var m motorcycle
	switch brand {
	case "honda":
		m = motorcycle{
			Brand: "Honda",
			Color: "The Red One",
		}
	case "suzuki":
		m = motorcycle{
			Brand: "Suzuki",
			Color: "The Blue One",
		}
	case "kawasaki":
		m = motorcycle{
			Brand: "Kawasaki",
			Color: "The Green One",
		}
	default:
		m = motorcycle{
			Brand: "404",
			Color: "We don't have that one",
		}
	}
	return &m
}
