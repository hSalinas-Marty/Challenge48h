package models

type Wine struct {
	Points      int     `json:"points"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	TasterName  *string `json:"taster_name"`
	Price       float64 `json:"price"`
	Designation *string `json:"designation"`
	Variety     string  `json:"variety"`
	Region1     *string `json:"region_1"`
	Region2     *string `json:"region_2"`
	Province    string  `json:"province"`
	Country     string  `json:"country"`
	Winery      string  `json:"winery"`
}

type WineFilter struct {
	Country   *string  `json:"country"`
	Variety   *string  `json:"variety"`
	MinPrice  *float64 `json:"min_price"`
	MaxPrice  *float64 `json:"max_price"`
	MinPoints *int     `json:"min_points"`
	Page      int      `json:"page"`
	PageSize  int      `json:"page_size"`
}
