// Package controllers gère la logique métier de l'application, y compris les gestionnaires HTTP.
package controllers

// PageVariables est une structure utilisée pour stocker les variables nécessaires pour afficher une page.
type PageVariables struct {
	ConvertedText string
}

// myError est une structure définie pour représenter une erreur personnalisée avec un code et un message d'erreur.
type myError struct {
	errorInt int
	errorMsg string
}
