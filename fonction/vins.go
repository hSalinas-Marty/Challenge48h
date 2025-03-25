package choix

import (
	"html/template"
	"net/http"
	"projet/json"
)

// Handler pour afficher les détails d'un vin spécifique
func VinDetailsHandler(w http.ResponseWriter, r *http.Request) {
	// Extraire le titre du vin à partir de l'URL (après /vin/)
	vinTitle := r.URL.Path[len("/vin/"):]

	// Charger les données du fichier JSON
	vins, err := json.Donner("json/wine-data-set.json")
	if err != nil {
		http.Error(w, "Erreur lors du chargement des données", http.StatusInternalServerError)
		return
	}

	// Chercher le vin correspondant au titre
	var vinDetail json.Vins
	found := false
	for _, vin := range vins {
		if vin.Title == vinTitle {
			vinDetail = vin
			found = true
			break
		}
	}

	// Si le vin n'est pas trouvé
	if !found {
		http.Error(w, "Vin non trouvé", http.StatusNotFound)
		return
	}

	// Charger le template de la page de détails
	tmpl, err := template.ParseFiles("template/details_vins.html")
	if err != nil {
		http.Error(w, "Erreur de chargement du template", http.StatusInternalServerError)
		return
	}

	// Exécuter le template avec les données du vin
	err = tmpl.Execute(w, vinDetail)
	if err != nil {
		http.Error(w, "Erreur lors de l'affichage de la page", http.StatusInternalServerError)
	}
}
