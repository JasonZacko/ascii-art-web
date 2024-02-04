// Package controllers gère la logique métier de l'application
package controllers

import (
	"bufio"
	"os"
)

// getFont charge les lignes d'un fichier de police spécifié.
// Il prend en paramètre le nom de la police et renvoie un tableau de lignes représentant la police.
// En cas d'erreur lors de la lecture du fichier, une erreur est renvoyée.
func getFont(font string) ([]string, error) {
	// Construire le chemin du fichier de police basé sur le nom de la police spécifié.
	fileName := "static/font/" + font + ".txt" 

	// Ouvrir le fichier de police.
	file, err := os.Open(fileName)
	if err != nil {
		// En cas d'erreur, renvoyer nil et l'erreur pour indiquer un problème.
		return nil, err
	}
	defer file.Close()

	// Lire les lignes du fichier de police.
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Ajouter chaque ligne au tableau de lignes.
		lines = append(lines, scanner.Text())
	}

	// Renvoyer le tableau de lignes représentant la police et aucune erreur.
	return lines, nil
}
