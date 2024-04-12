package main

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
)

const homeTemplate = `<!DOCTYPE html>
<html>
<head>
    <title>Home</title>
    <link rel="stylesheet" href="/static/styles.css">
</head>
<body>
    <h1>Pick A Motorcycle</h1>
<button class="Suzuki"><a href="/suzuki">Suzuki</a></button>
<button class="Honda"><a href="/honda">Honda</a></button>
<button class="Kawasaki"><a href="/kawasaki">Kawasaki</a></button>
</body>
</html>`

const motoTemplate = `<!DOCTYPE html>
<html>
<head>
    <title>{{.Brand}}</title>
    <link rel="stylesheet" href="/static/styles.css">
</head>
<body>
    <h1 class="{{.Brand}}">{{.Brand}}</h1>
    <p>{{.Model}}</p>
    <button><a href="/"> <- Go Back</a></button>
</body>
</html>`

type Motorcycle struct {
	Brand string
	Model string
}

const (
	homePath       = "/"
	motorcyclePath = "/{motorcycle}"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc(homePath, serveHome)
	http.HandleFunc(motorcyclePath, serveMotorcycle)

	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func renderTemplate(w http.ResponseWriter, templateName string, data interface{}) {
	t, err := template.New(templateName).Parse(templateName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var buf bytes.Buffer
	err = t.Execute(&buf, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	buf.WriteTo(w)
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	data := ""
	renderTemplate(w, homeTemplate, data)
}

func serveMotorcycle(w http.ResponseWriter, r *http.Request) {
	whichMotorcycle := r.PathValue("motorcycle")
	data := findMotorcycle(whichMotorcycle)
	renderTemplate(w, motoTemplate, data)
}

func findMotorcycle(brand string) *Motorcycle {
	var m Motorcycle
	switch brand {
	case "honda":
		m = Motorcycle{
			Brand: "Honda",
			Model: "The Red One",
		}
	case "suzuki":
		m = Motorcycle{
			Brand: "Suzuki",
			Model: "The Blue One",
		}
	case "kawasaki":
		m = Motorcycle{
			Brand: "Kawasaki",
			Model: "The Green One",
		}
	default:
		m = Motorcycle{
			Brand: "404",
			Model: "We don't have that one",
		}
	}
	return &m
}
