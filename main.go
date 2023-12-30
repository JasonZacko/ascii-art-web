package main

import (
	"controller/controller"
	"fmt"
	"html/template"
	"net/http"
)

// PageData est une structure pour stocker les données à afficher dans le template HTML
type PageData struct {
	Title    string
	ASCIIArt string
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	// ... Reste de votre code ...

	http.HandleFunc("/", HandleIndex)
	http.HandleFunc("/ascii-art", HandleASCIIArt)

	http.ListenAndServe(":8080", nil)
}

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	// Affichez le formulaire et la page principale
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := PageData{
		Title:    "ASCII Art Web",
		ASCIIArt: "", // Laissez cette valeur vide pour l'instant
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func HandleASCIIArt(w http.ResponseWriter, r *http.Request) {
	// Récupérez les données POST et générez l'ASCII Art
	r.ParseForm()
	text := r.FormValue("text")
	banner := r.FormValue("banner")

	result, err := controller.ConvertInputWeb(text, banner)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Affichez le résultat sur la page principale
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := PageData{
		Title:    "ASCII Art Web",
		ASCIIArt: result,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
