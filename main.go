package main

import (
	"ascii-art-web/controllers"
	"fmt"
	"net/http"
)

func main() {
	// Définir le gestionnaire pour servir le fichier HTML
	http.HandleFunc("/", controllers.HomePage)

	// Gérer les fichiers statiques (CSS, JavaScript, etc.)
	// StripPrefix est utilisé pour retirer le préfixe "/static/" de l'URL
	// http.FileServer(http.Dir ..) créer le serveur afin de servir les fichiers à partir du chemin
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	// Démarrer le serveur sur le port 8080
	fmt.Println("Serveur démarré sur http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
