package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"projet/json"
)

// Handler pour afficher les informations du vin
func wineHandler(w http.ResponseWriter, r *http.Request) {
	// Charger les données du fichier JSON
	vins, err := json.Donner("json/wine-data-set.json")
	if err != nil {
		http.Error(w, "Erreur lors du chargement des données", http.StatusInternalServerError)
		return
	}

	// Charger le template HTML
	tmpl, err := template.ParseFiles("template/index.html")
	if err != nil {
		http.Error(w, "Erreur de chargement du template", http.StatusInternalServerError)
		return
	}

	// Exécuter le template et envoyer la réponse avec les données JSON
	err = tmpl.Execute(w, vins) // On passe les données JSON (vins) au template
	if err != nil {
		http.Error(w, "Erreur lors de l'affichage de la page", http.StatusInternalServerError)
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
