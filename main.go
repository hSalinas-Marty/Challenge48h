package main

import (
	jsoncha "Challenge48h/json"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

// Structure pour gérer les traductions
type Translation struct {
	Language      string `json:"language"` // Ajout de la langue pour la gestion dans le HTML
	Title         string `json:"title"`
	Description   string `json:"description"`
	WineListTitle string `json:"wineListTitle"`
	Prices        string `json:"prices"`
	Variety       string `json:"variety"`
	Region        string `json:"region"`
	Taster        string `json:"taster"`
	Twitter       string `json:"twitter"`
	Winery        string `json:"winery"`
}

// Fonction pour charger les traductions depuis un fichier JSON
func loadTranslations(lang string) (Translation, error) {
	var translations Translation
	file, err := os.Open(fmt.Sprintf("translations/%s.json", lang))
	if err != nil {
		return translations, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&translations)
	if err != nil {
		return translations, err
	}
	translations.Language = lang // Ajout de la langue sélectionnée
	return translations, nil
}

// Handler pour afficher les informations du vin
func wineHandler(w http.ResponseWriter, r *http.Request) {
	// Déterminer la langue depuis le paramètre URL
	lang := r.URL.Query().Get("lang")
	if lang == "" {
		lang = "fr" // Langue par défaut
	}

	// Charger les traductions
	translations, err := loadTranslations(lang)
	if err != nil {
		// Log the error for debugging
		log.Printf("Erreur lors du chargement des traductions: %v", err)
		http.Error(w, "Erreur lors du chargement des traductions", http.StatusInternalServerError)
		return
	}

	// Charger le fichier des vins en fonction de la langue
	var wineFile string
	switch lang {
	case "en":
		wineFile = "json/wine-data-en.json"
	default:
		wineFile = "json/wine-data-fr.json"
	}

	// Charger les données du fichier JSON
	vins, err := jsoncha.Donner(wineFile)
	if err != nil {
		// Log the error for debugging
		log.Printf("Erreur lors du chargement des données des vins: %v", err)
		http.Error(w, "Erreur lors du chargement des données", http.StatusInternalServerError)
		return
	}

	// Charger le template HTML
	tmpl, err := template.ParseFiles("template/index.html")
	if err != nil {
		// Log the error for debugging
		log.Printf("Erreur de chargement du template: %v", err)
		http.Error(w, "Erreur de chargement du template", http.StatusInternalServerError)
		return
	}

	// Passer les données de vin et de traduction au template
	data := struct {
		Vins         interface{}
		Translations Translation
	}{
		Vins:         vins,
		Translations: translations,
	}

	// Exécuter le template et envoyer la réponse avec les données JSON et les traductions
	err = tmpl.Execute(w, data)
	if err != nil {
		// Log the error for debugging
		log.Printf("Erreur lors de l'affichage du template: %v", err)
		http.Error(w, "Erreur lors de l'affichage de la page", http.StatusInternalServerError)
		return
	}
}

func main() {
	// Servir les fichiers statiques (CSS, images)
	fs := http.FileServer(http.Dir("template/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Route principale
	http.HandleFunc("/", wineHandler)

	// Lancer le serveur
	fmt.Println("Serveur démarré sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
