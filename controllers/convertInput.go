// Package controllers gère la logique métier de l'application
package controllers

import (
	"strings"
)

// convertInput prend une ligne de texte et un tableau de lignes représentant une police ASCII.
// Il convertit la ligne de texte en une représentation ASCII artistique en utilisant la police fournie.
func convertInput(line string, lines []string) string {
	// Diviser la ligne de texte en un tableau de lignes en utilisant le saut de ligne comme délimiteur.
	arr := strings.Split(line, "\n")

	// Initialiser une variable strings.Builder pour construire le résultat final.
	var result strings.Builder
	
	// Vérifier si la ligne de texte n'est pas vide.
	if line != "" {
		// Parcourir chaque ligne dans le tableau résultant de la division.
		for _, str := range arr {
			// Diviser chaque ligne en un tableau de lignes en utilisant "\n" comme délimiteur.
			arr2 := strings.Split(str, "\\n")
			for _, str2 := range arr2 {
				// Parcourir chaque caractère dans la ligne divisée.
				for i := 1; i < 9; i++ {
					for _, char := range str2 {
						// Convertir le caractère en valeur ASCII.
						ascii := int(char) - 32
						// Calculer l'index dans le tableau de lignes basé sur la police ASCII et l'indice actuel.
						index := ascii*9 + i

						// Vérifier si l'index est dans la plage valide des lignes.
						if index >= 0 && index < len(lines) {
							// Ajouter la ligne correspondante au résultat.
							result.WriteString(lines[index])
						} else {
							// Ignorer l'index invalide et ne rien imprimer dans le terminal.
						}
					}
					// Ajouter un saut de ligne à la fin de chaque ligne convertie.
					result.WriteString("\n")
				}
				// Ajouter une ligne vide entre chaque ligne convertie.
				
			}
		}
	}

	// Renvoyer la représentation ASCII complète générée à partir de la ligne de texte fournie.
	return result.String()
}
