package repository

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "sync"

    "cite-du-vin/internal/models"
)

type WineRepository interface {
    GetWines(filter *models.WineFilter) ([]models.Wine, int, error)
    GetWineByID(id string) (*models.Wine, error)
}

type MemoryWineRepository struct {
    wines []models.Wine
    mu    sync.RWMutex
}

func NewMemoryWineRepository(dataPath string) (*MemoryWineRepository, error) {
    repo := &MemoryWineRepository{}
    
    // Read JSON file
    data, err := ioutil.ReadFile(dataPath)
    if err != nil {
        return nil, fmt.Errorf("error reading wine dataset: %v", err)
    }

    // Unmarshal JSON
    if err := json.Unmarshal(data, &repo.wines); err != nil {
        return nil, fmt.Errorf("error parsing wine dataset: %v", err)
    }

    return repo, nil
}

func (r *MemoryWineRepository) GetWines(filter *models.WineFilter) ([]models.Wine, int, error) {
    r.mu.RLock()
    defer r.mu.RUnlock()

    // Create a copy of wines to filter
    filteredWines := make([]models.Wine, 0, len(r.wines))
    
    for _, wine := range r.wines {
        // Apply filters
        if filter.Country != nil && wine.Country != *filter.Country {
            continue
        }
        if filter.Variety != nil && wine.Variety != *filter.Variety {
            continue
        }
        if filter.MinPrice != nil && wine.Price < *filter.MinPrice {
            continue
        }
        if filter.MaxPrice != nil && wine.Price > *filter.MaxPrice {
            continue
        }
        if filter.MinPoints != nil && wine.Points < *filter.MinPoints {
            continue
        }

        filteredWines = append(filteredWines, wine)
    }

    // Pagination
    pageSize := 10
    if filter.PageSize > 0 {
        pageSize = filter.PageSize
    }
    page := filter.Page
    if page < 1 {
        page = 1
    }

    start := (page - 1) * pageSize
    end := start + pageSize
    if start > len(filteredWines) {
        start = len(filteredWines)
    }
    if end > len(filteredWines) {
        end = len(filteredWines)
    }

    return filteredWines[start:end], len(filteredWines), nil
}

func (r *MemoryWineRepository) GetWineByID(id string) (*models.Wine, error) {
    r.mu.RLock()
    defer r.mu.RUnlock()

    // This is a placeholder. In a real implementation, you'd need a way to uniquely identify wines
    for _, wine := range r.wines {
        if wine.Title == id {
            return &wine, nil
        }
    }

    return nil, fmt.Errorf("wine not found")
}