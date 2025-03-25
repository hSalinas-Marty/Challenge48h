package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"

    "cite-du-vin/internal/models"
    "cite-du-vin/internal/repository"
)

func GetWinesHandler(repo repository.WineRepository) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Parse query parameters for filtering
        filter := &models.WineFilter{
            Page:     1,
            PageSize: 10,
        }

        // Extract country filter
        if country := r.URL.Query().Get("country"); country != "" {
            filter.Country = &country
        }

        // Extract variety filter
        if variety := r.URL.Query().Get("variety"); variety != "" {
            filter.Variety = &variety
        }

        // Extract price filters
        if minPrice := r.URL.Query().Get("min_price"); minPrice != "" {
            if price, err := strconv.ParseFloat(minPrice, 64); err == nil {
                filter.MinPrice = &price
            }
        }

        if maxPrice := r.URL.Query().Get("max_price"); maxPrice != "" {
            if price, err := strconv.ParseFloat(maxPrice, 64); err == nil {
                filter.MaxPrice = &price
            }
        }

        // Extract page and page size
        if page := r.URL.Query().Get("page"); page != "" {
            if pageNum, err := strconv.Atoi(page); err == nil {
                filter.Page = pageNum
            }
        }

        if pageSize := r.URL.Query().Get("page_size"); pageSize != "" {
            if size, err := strconv.Atoi(pageSize); err == nil {
                filter.PageSize = size
            }
        }

        // Get filtered wines
        wines, total, err := repo.GetWines(filter)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        // Prepare response
        response := struct {
            Wines      []models.Wine `json:"wines"`
            Total     int            `json:"total"`
            Page      int            `json:"page"`
            PageSize  int            `json:"page_size"`
        }{
            Wines:     wines,
            Total:    total,
            Page:     filter.Page,
            PageSize: filter.PageSize,
        }

        // Send JSON response
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(response)
    }
}

func GetWineByIDHandler(repo repository.WineRepository) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Get ID from URL parameters
        vars := mux.Vars(r)
        id := vars["id"]

        // Fetch wine
        wine, err := repo.GetWineByID(id)
        if err != nil {
            http.Error(w, err.Error(), http.StatusNotFound)
            return
        }

        // Send JSON response
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(wine)
    }
}