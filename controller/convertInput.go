package controller

import (
	"strings"
)

// ConvertInputWeb génère l'ASCII Art en fonction du texte et des lignes spécifiés.
func ConvertInputWeb(text, lines string) (string, error) {
	// Vérifie si le texte n'est pas vide
	if text != "" {
		var result strings.Builder

		// Parcourt chaque caractère du texte (char)
		for _, char := range text {
			ascii := int(char) - 32

			// Ajoute le caractère correspondant à la position (ascii*9 + i) dans la slice lines
			result.WriteString(lines[ascii*9+1 : ascii*9+10])
		}

		// Ajoute un saut de ligne après chaque ligne générée
		result.WriteString("\n")

		// Renvoie la représentation ASCII générée
		return result.String(), nil
	}

	// Si le texte est vide, renvoie une ligne vide
	return "", nil
}
