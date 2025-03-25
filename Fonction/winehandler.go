package Fonction

import (
	"html/template"
	"net/http"
	"projet/json"
	"strconv" // Permet de convertir string en int
)

// Handler pour afficher la liste des vins
func WineHandler(w http.ResponseWriter, r *http.Request) {
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
	err = tmpl.Execute(w, vins)
	if err != nil {
		http.Error(w, "Erreur lors du rendu du template", http.StatusInternalServerError)
	}
}

// Handler pour afficher les détails d'un vin spécifique
func VinDetailsHandler(w http.ResponseWriter, r *http.Request) {
	// Extraire l'ID du vin à partir de l'URL
	vinIDStr := r.URL.Path[len("/vin/"):]
	vinID, err := strconv.Atoi(vinIDStr) // Convertir l'ID en int
	if err != nil {
		http.Error(w, "ID de vin invalide", http.StatusBadRequest)
		return
	}

	// Charger les données du fichier JSON
	vins, err := json.Donner("json/wine-data-set.json")
	if err != nil {
		http.Error(w, "Erreur lors du chargement des données", http.StatusInternalServerError)
		return
	}

	// Trouver le vin correspondant à l'ID
	var vinDetail json.Vins
	for _, vin := range vins {
		if vin.ID == vinID {
			vinDetail = vin
			break
		}
	}

	// Vérifier si le vin existe
	if vinDetail.Title == "" {
		http.Error(w, "Vin non trouvé", http.StatusNotFound)
		return
	}

	// Charger le template HTML de la page de détails
	tmpl, err := template.ParseFiles("template/details_vins.html")
	if err != nil {
		http.Error(w, "Erreur de chargement du template", http.StatusInternalServerError)
		return
	}

	// Exécuter le template avec les données du vin spécifique
	err = tmpl.Execute(w, vinDetail)
	if err != nil {
		http.Error(w, "Erreur lors du rendu du template", http.StatusInternalServerError)
	}
}
