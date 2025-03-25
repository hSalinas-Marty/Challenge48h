package fonction

import (
	"fmt"
	"html/template"
	"net/http"
	"projet/json"
	"strconv"
)

func WineHandler(w http.ResponseWriter, r *http.Request) {
	// Charger les données du fichier JSON
	vins, err := json.Donner("json/wine-data-set.json")
	if err != nil {
		// Lorsque tu utilises http.Error, ça envoie automatiquement un WriteHeader
		http.Error(w, "Erreur lors du chargement des données", http.StatusInternalServerError)
		return // Ne fais pas de WriteHeader manuellement après
	}

	// Sélectionner un sous-ensemble des vins pour "Les Vins du Moment" (par exemple, les 5 premiers)
	var vinsDuMoment []json.Vins
	for i := 0; i < 5 && i < len(vins); i++ {
		vinsDuMoment = append(vinsDuMoment, vins[i])
	}

	// Afficher les vins du moment dans la console pour vérifier
	fmt.Println(vinsDuMoment)

	// Charger le template HTML de la page d'index
	tmpl, err := template.ParseFiles("template/index.html")
	if err != nil {
		http.Error(w, "Erreur de chargement du template", http.StatusInternalServerError)
		return
	}

	// Passer les vins du moment et les autres données au template
	err = tmpl.Execute(w, struct {
		VinsDuMoment []json.Vins
	}{
		VinsDuMoment: vinsDuMoment, // Vins du moment
	})
	if err != nil {
		http.Error(w, "Erreur lors du rendu du template", http.StatusInternalServerError)
		return // Ne fais pas d'autres appels WriteHeader après
	}
}

// Handler pour afficher les détails d'un vin spécifique
func VinDetailsHandler(w http.ResponseWriter, r *http.Request) {
	// Extraire l'ID du vin à partir de l'URL
	vinIDStr := r.URL.Path[len("/vin/"):]

	// Convertir l'ID en int
	vinID, err := strconv.Atoi(vinIDStr)
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
		http.Error(w, "Erreur de chargement du template Charger le template HTML de la page de détails", http.StatusInternalServerError)
		return
	}

	// Exécuter le template avec les données du vin spécifique
	err = tmpl.Execute(w, vinDetail)
	if err != nil {
		http.Error(w, "Erreur lors du rendu du template Exécuter le template avec les données du vin spécifique", http.StatusInternalServerError)
		return
	}
}
