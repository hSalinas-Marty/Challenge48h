package server

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"projet/structure"
	"strings"
)

// Charger les vins depuis un fichier JSON.
func LoadWines(filename string) ([]structure.Wine, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("erreur d'ouverture du fichier : %w", err)
	}
	defer file.Close()

	var wines []structure.Wine
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&wines); err != nil {
		return nil, fmt.Errorf("erreur de décodage JSON : %w", err)
	}

	return wines, nil
}

// Handlers pour la page d'accueil.
func WineHandler(w http.ResponseWriter, r *http.Request) {
	wines, err := LoadWines("static/json/wine-data-example.json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("static/templates/test.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, wines)
}

// Fonction qui effectue la recherche dans les descriptions des vins
func SearchWines(w http.ResponseWriter, r *http.Request) {
	// Récupérer la requête de l'utilisateur
	query := r.FormValue("query")

	// Charger les vins à partir du fichier JSON
	wines, err := LoadWines("static/json/wine-data-example.json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Créer une liste pour stocker les vins correspondants
	var matchedWines []structure.Wine

	// Vérifier si la requête de l'utilisateur correspond à une description de vin
	for _, wine := range wines {
		// Convertir en minuscules pour une recherche insensible à la casse
		if strings.Contains(strings.ToLower(wine.DetailedDescription), strings.ToLower(query)) {
			matchedWines = append(matchedWines, wine)
		}
	}

	// Passer les vins trouvés à la vue HTML
	tmpl, err := template.ParseFiles("static/templates/search_results.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Afficher les résultats
	tmpl.Execute(w, matchedWines)
}
