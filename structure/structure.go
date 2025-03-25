package structure

// Wine représente un vin avec ses propriétés
type Wine struct {
	Points              int     `json:"points"`
	Title               string  `json:"title"`
	Description         string  `json:"description"`
	TasterName          *string `json:"taster_name"`           // Utilisation d'un pointeur pour gérer les valeurs nulles
	TasterTwitter       *string `json:"taster_twitter_handle"` // Utilisation d'un pointeur pour gérer les valeurs nulles
	Price               float64 `json:"price"`                 // Utilisation de float64 pour gérer les prix décimaux
	Designation         *string `json:"designation"`           // Utilisation d'un pointeur pour gérer les valeurs nulles
	Variety             string  `json:"variety"`
	Region1             string  `json:"region_1"`
	Region2             *string `json:"region_2"` // Utilisation d'un pointeur pour gérer les valeurs nulles
	Province            string  `json:"province"`
	Country             string  `json:"country"`
	Winery              string  `json:"winery"`
	ImageURL            string  `json:"image_url"`            // URL de l'image
	DetailedDescription string  `json:"detailed_description"` // Description détaillée
}

type WineSearch struct {
	Title               string `json:"title"`
	Description         string `json:"description"`
	DetailedDescription string `json:"detailed_description"`
}
