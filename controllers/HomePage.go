// Package controllers gère la logique métier de l'application, y compris les gestionnaires HTTP.
package controllers

import (
	"fmt"
	"html/template"
	"net/http"
)

// HomePage est un gestionnaire HTTP pour la page d'accueil.
func HomePage(w http.ResponseWriter, r *http.Request) {
	// Vérifier si l'URL correspond à la racine du site
	if r.URL.Path == "/" {
		// Vérifier si la méthode HTTP est POST (formulaire soumis)
		if r.Method == http.MethodPost {
			// Récupérer les données du formulaire
			text := r.FormValue("inputText")
			font := r.FormValue("fontSelect")

			// Vérifier si les caractères font partie des caractères imprimables ASCII
			for i := 0; i < len(text); i++ {
				if text[i] < ' ' && text[i] != '\n' && text[i] != '\r' || text[i] > '~' {
					http.Error(w, "Bad Request", http.StatusBadRequest)
					return
				}
			}

			// Obtenir les lignes de la police spécifiée
			lines, err := getFont(font)
			if err != nil {
				http.Error(w, "Page Not Found", http.StatusNotFound)
				return
			}

			// Convertir le texte
			convertedText := convertInput(text, lines)

			// Créer une structure de données pour stocker les variables de la page
			pageVariables := PageVariables{
				ConvertedText: convertedText,
			}

			// Utiliser le modèle pour générer la page HTML
			tmpl, err := template.ParseFiles("static/html/index.html")
			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			// Remplacer la partie {{.ConvertedText}} dans le fichier HTML par le texte converti
			err = tmpl.Execute(w, pageVariables)
			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
		} else {
			// Si la méthode HTTP n'est pas POST, simplement servir la page HTML
			tmpl, err := template.ParseFiles("static/html/index.html")
			if err != nil {
				// En cas d'erreur, servir une page d'erreur 500
				fmt.Println(err)
				tmpl, _ := template.ParseFiles("static/html/error.html")
				tmpl.Execute(w, "Error 500 : Internal server error")
				return
			}
			// Servir la page HTML sans traitement spécial
			tmpl.Execute(w, nil)
		}
	} else {
		// Si l'URL ne correspond pas à la racine, renvoyer une erreur 404
		http.Error(w, "Page Not Found", http.StatusNotFound)
	}
}
