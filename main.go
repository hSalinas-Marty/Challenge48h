package main

import (
	"fmt"
	"html/template"
	"net/http"
	"projet/choix"
	"projet/json"
)

// Handler pour afficher la liste des vins
func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Charger les données du fichier JSON
	vins, err := json.Donner("json/wine-data-set.json")
	if err != nil {
		http.Error(w, "Erreur lors du chargement des données", http.StatusInternalServerError)
		return
	}

	// Charger le template de la page d'index
	tmpl, err := template.ParseFiles("template/index.html")
	if err != nil {
		http.Error(w, "Erreur de chargement du template", http.StatusInternalServerError)
		return
	}

	// Afficher la page d'index avec la liste des vins
	tmpl.Execute(w, vins)
}

func main() {
	// Servir les fichiers statiques (CSS)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Route principale pour afficher la liste des vins
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/", winehandler.WineHandler)
	// Route pour afficher les détails d'un vin
	http.HandleFunc("/vin/", choix.VinDetailsHandler)

	// Lancer le serveur
	port := ":8080"
	fmt.Println("Serveur Go lancé sur http://localhost" + port)
	http.ListenAndServe(port, nil)
}
