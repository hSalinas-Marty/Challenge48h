package json

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Structure pour stocker les données JSON
type Vins struct {
	Points        int     `json:"points"`
	Title         string  `json:"title"`
	Description   string  `json:"description"`
	TasterName    *string `json:"taster_name"`
	TasterTwitter *string `json:"taster_twitter_handle"`
	Price         int     `json:"price"`
	Designation   *string `json:"designation"`
	Variety       string  `json:"variety"`
	Region1       string  `json:"region_1"`
	Region2       string  `json:"region_2"`
	Province      string  `json:"province"`
	Country       string  `json:"country"`
	Winery        string  `json:"winery"`
}

// Fonction pour lire les données JSON
func Donner(nom string) ([]Vins, error) {
	var vins []Vins

	// Lire le fichier JSON
	fmt.Println("Tentative de lecture du fichier:", nom) // Ajout d'un log pour vérifier le chemin

	file, err := ioutil.ReadFile(nom)
	if err != nil {
		// Log détaillé sur l'erreur de lecture du fichier
		return vins, fmt.Errorf("Erreur lors de la lecture du fichier JSON: %v", err)
	}

	// Convertir le JSON en struct Go
	err = json.Unmarshal(file, &vins)
	if err != nil {
		// Log détaillé sur l'erreur de désérialisation
		return vins, fmt.Errorf("Erreur lors du décodage du fichier JSON: %v", err)
	}

	// Filtrer les vins pour ne garder que ceux avec toutes les informations
	vins = FiltrerVins(vins)

	return vins, nil
}

// Fonction pour filtrer les vins en excluant ceux qui ont des valeurs nulles
func FiltrerVins(vins []Vins) []Vins {
	var vinsFiltres []Vins
	for _, vin := range vins {
		// Vérifier si tous les champs obligatoires sont non nuls
		if vin.Title != "" && vin.Description != "" && vin.Variety != "" && vin.Region1 != "" &&
			vin.Province != "" && vin.Country != "" && vin.Winery != "" && vin.Price > 0 {

			// Ajouter le vin à la liste des vins filtrés si toutes les informations sont présentes
			vinsFiltres = append(vinsFiltres, vin)
		} else {
			// Log ou message pour vérifier les vins rejetés
			fmt.Println("Vin rejeté:", vin.Title)
		}
	}

	fmt.Printf("Nombre de vins après filtrage: %d\n", len(vinsFiltres)) // Affiche le nombre de vins restants après le filtrage
	return vinsFiltres
}
