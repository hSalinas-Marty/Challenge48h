package main

import (
	"fmt"
	"html/template"
	"net/http"

	"json/importjson"
)

// Handler pour afficher les informations du vin
func wineHandler(w http.ResponseWriter, r *http.Request) {
	// Charger les données du fichier JSON
	wines, err := importjson.LoadWineData("data/vin.json")
	if err != nil {
		http.Error(w, "Erreur lors du chargement des données", http.StatusInternalServerError)
		return
	}

	// Charger le template HTML
	tmpl, err := template.ParseFiles("templates/vin.html")
	if err != nil {
		http.Error(w, "Erreur de chargement du template", http.StatusInternalServerError)
		return
	}

	// Envoyer les données JSON au template
	tmpl.Execute(w, wines)
}

func main() {
	// Servir les fichiers statiques (CSS, images)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Route principale
	http.HandleFunc("/vin", wineHandler)

	// Lancer le serveur
	port := ":8080"
	fmt.Println("Serveur Go lancé sur http://localhost" + port)
	http.ListenAndServe(port, nil)
}
