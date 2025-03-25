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

	tmpl, err := template.ParseFiles("static/templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, wines)
}

// Structure to hold the query and matched wines
type SearchResults struct {
	Query        string
	MatchedWines []structure.Wine
}

func SearchWines(w http.ResponseWriter, r *http.Request) {
	query := strings.TrimSpace(r.FormValue("query"))

	wines, err := LoadWines("static/json/wine-data-example.json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Préparer les données pour le template
	data := struct {
		Query        string
		MatchedWines []structure.Wine
		ShowAll      bool
	}{
		Query:        query,
		MatchedWines: wines, // Par défaut, montre tous les vins
		ShowAll:      true,
	}

	// Si une recherche a été effectuée
	if query != "" {
		words := strings.Fields(strings.ToLower(query))
		var matchedWines []structure.Wine

		// Détection des filtres
		filters := make(map[string]string)

		// Identifier les filtres dans la requête
		for _, word := range words {
			lowerWord := strings.ToLower(word)
			for _, wine := range wines {
				if strings.ToLower(wine.Country) == lowerWord {
					filters["country"] = lowerWord
					break
				}
				if strings.ToLower(wine.Region1) == lowerWord {
					filters["region1"] = lowerWord
					break
				}
				if strings.ToLower(wine.Province) == lowerWord {
					filters["province"] = lowerWord
					break
				}
				if strings.ToLower(wine.Variety) == lowerWord {
					filters["variety"] = lowerWord
					break
				}
				if strings.ToLower(wine.Winery) == lowerWord {
					filters["winery"] = lowerWord
					break
				}
			}
		}

		// Appliquer les filtres
		for _, wine := range wines {
			match := true

			if country, ok := filters["country"]; ok && strings.ToLower(wine.Country) != country {
				match = false
			}
			if region, ok := filters["region1"]; ok && match && strings.ToLower(wine.Region1) != region {
				match = false
			}
			if province, ok := filters["province"]; ok && match && strings.ToLower(wine.Province) != province {
				match = false
			}
			if variety, ok := filters["variety"]; ok && match && strings.ToLower(wine.Variety) != variety {
				match = false
			}
			if winery, ok := filters["winery"]; ok && match && strings.ToLower(wine.Winery) != winery {
				match = false
			}

			if match {
				matchedWines = append(matchedWines, wine)
			}
		}

		// Filtrer par mots-clés
		if len(words) > 0 {
			var keywordResults []structure.Wine

			for _, wine := range matchedWines {
				matchesAll := true

				for _, word := range words {
					lowerWord := strings.ToLower(word)

					// Ignorer les mots utilisés comme filtres
					if _, ok := filters["country"]; ok && lowerWord == filters["country"] {
						continue
					}
					if _, ok := filters["region1"]; ok && lowerWord == filters["region1"] {
						continue
					}
					if _, ok := filters["province"]; ok && lowerWord == filters["province"] {
						continue
					}
					if _, ok := filters["variety"]; ok && lowerWord == filters["variety"] {
						continue
					}
					if _, ok := filters["winery"]; ok && lowerWord == filters["winery"] {
						continue
					}

					// Rechercher dans les champs pertinents
					found := strings.Contains(strings.ToLower(wine.Title), lowerWord) ||
						strings.Contains(strings.ToLower(wine.Description), lowerWord) ||
						strings.Contains(strings.ToLower(wine.DetailedDescription), lowerWord) ||
						(wine.Designation != nil && strings.Contains(strings.ToLower(*wine.Designation), lowerWord))

					if !found {
						matchesAll = false
						break
					}
				}

				if matchesAll {
					keywordResults = append(keywordResults, wine)
				}
			}

			matchedWines = keywordResults
		}

		data.MatchedWines = matchedWines
		data.ShowAll = false
	}

	// Charger et exécuter le template
	tmpl, err := template.ParseFiles("static/templates/form.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
