package json

import (
	"encoding/json"
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
	file, err := ioutil.ReadFile(nom)
	if err != nil {
		return vins, err
	}

	// Convertir le JSON en struct Go
	err = json.Unmarshal(file, &vins)
	if err != nil {
		return vins, err
	}

	return vins, nil
}
