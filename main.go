package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// Struct représentant un vin
type Wine struct {
	Points              int     `json:"points"`
	Title               string  `json:"title"`
	TasterName          *string `json:"taster_name"`
	TasterTwitterHandle *string `json:"taster_twitter_handle"`
	Price               float64 `json:"price"`
	Designation         *string `json:"designation"`
	Variety             string  `json:"variety"`
	Region1             string  `json:"region_1"`
	Region2             *string `json:"region_2"`
	Province            string  `json:"province"`
	Country             string  `json:"country"`
	Winery              string  `json:"winery"`
	ImageURL            string  `json:"image_url"`
	DetailedDescription string  `json:"detailed_description"`
	Type                int     `json:"type"`
}

// Fonction pour obtenir le type de vin sous forme de chaîne
func (w Wine) TypeStr() string {
	switch w.Type {
	case 1:
		return "Rouge"
	case 2:
		return "Blanc"
	case 3:
		return "Rosé"
	case 4:
		return "Pétillant"
	case 5:
		return "Dessert"
	default:
		return "Inconnu"
	}
}

// Charger les vins à partir du fichier JSON
func LoadWines(filename string) ([]Wine, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var wines []Wine
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&wines)
	if err != nil {
		return nil, err
	}
	return wines, nil
}

// Fonction pour vérifier si une string est présente dans une autre string
func contains(substring, str string) bool {
	return strings.Contains(str, substring)
}

// Handler pour le quiz
func QuizHandler(w http.ResponseWriter, r *http.Request) {
	// Charger tous les vins depuis le fichier JSON
	wines, err := LoadWines("static/json/wine-data-example.json")
	var filtred []Wine
	if err != nil {
		log.Printf("Erreur lors du chargement des vins: %v", err)
		http.Error(w, "Could not load wine data", http.StatusInternalServerError)
		return
	}

	// Affichage des vins avant filtrage
	log.Println("Liste des vins avant filtrage:")
	for _, wine := range wines {
		log.Printf("Titre: %s, Type: %s", wine.Title, wine.TypeStr())
	}

	// Si la méthode est POST (envoi du formulaire)
	if r.Method == http.MethodPost {
		// Récupérer les valeurs du formulaire
		quizTypeStr := r.FormValue("type") // Type de vin (Rouge, Blanc, etc.)
		pays := r.FormValue("country")     // Région sélectionnée par l'utilisateur
		priceRange := r.FormValue("price") // Plage de prix sélectionnée

		// Convertir le type de vin en int
		quizType, err := strconv.Atoi(quizTypeStr)
		if err != nil {
			quizType = 0 // "Tous" si erreur de conversion
		}

		// Filtrer par type de vin
		if quizType != 0 {
			for _, filtredType := range wines {
				if quizType == 1 && filtredType.Type == 1 {
					filtred = append(filtred, filtredType)
				} else if quizType == 2 && filtredType.Type == 2 {
					filtred = append(filtred, filtredType)
				} else if quizType == 3 && filtredType.Type == 3 {
					filtred = append(filtred, filtredType)
				} else if quizType == 4 && filtredType.Type == 4 {
					filtred = append(filtred, filtredType)
				} else if quizType == 5 && filtredType.Type == 5 {
					filtred = append(filtred, filtredType)
				}
			}
		} else {
			// Si quizType est 0, tous les vins sont conservés
			filtred = wines
		}

		fmt.Println("--------------------------------------------------------")
		fmt.Println(" Apres type : ")
		fmt.Println(filtred)
		fmt.Println("--------------------------------------------------------")

		// Filtrer par pays
		if pays != "" {
			var filteredByRegion []Wine
			for _, filtredRegion := range filtred {
				if contains(filtredRegion.Country, pays) {
					filteredByRegion = append(filteredByRegion, filtredRegion)
				}
			}
			filtred = filteredByRegion
		}

		fmt.Println("--------------------------------------------------------")
		fmt.Println(" Apres Region : ")
		fmt.Println(filtred)
		fmt.Println("--------------------------------------------------------")

		// Filtrer par prix
		if priceRange != "0" {
			var filteredByPrice []Wine
			for _, filtredPrice := range filtred {
				if priceRange == "1" && filtredPrice.Price >= 0 && filtredPrice.Price <= 50 {
					filteredByPrice = append(filteredByPrice, filtredPrice)
				} else if priceRange == "2" && filtredPrice.Price > 50 && filtredPrice.Price <= 150 {
					filteredByPrice = append(filteredByPrice, filtredPrice)
				} else if priceRange == "3" && filtredPrice.Price > 150 {
					filteredByPrice = append(filteredByPrice, filtredPrice)
				}
			}
			filtred = filteredByPrice
		}

		fmt.Println("--------------------------------------------------------")
		fmt.Println(" Apres Price : ")
		fmt.Println(filtred)
		fmt.Println("--------------------------------------------------------")
	}

	// Rendre le template avec les vins filtrés
	tmplPath := filepath.Join("static", "templates", "form.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		log.Printf("Erreur de chargement du template: %v", err)
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	// Passer les vins filtrés au template
	err = tmpl.Execute(w, filtred)
	if err != nil {
		log.Printf("Erreur lors du rendu du template: %v", err)
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Démarrer le serveur
	http.HandleFunc("/", QuizHandler)
	log.Println("Démarrage du serveur sur http://localhost:1717")
	err := http.ListenAndServe(":1717", nil)
	if err != nil {
		log.Fatal("Erreur lors du démarrage du serveur : ", err)
	}
}
