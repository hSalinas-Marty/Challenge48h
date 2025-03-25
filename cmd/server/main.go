package main

import (
    "fmt"
    "log"
    "net/http"
    "os"

    "github.com/gorilla/mux"
    "github.com/rs/cors"

    "cite-du-vin/internal/handlers"
    "cite-du-vin/internal/repository"
)

func main() {
    // Initialize wine repository
    wineRepo, err := repository.NewMemoryWineRepository("./data/wine-data-set.json")
    if err != nil {
        log.Fatalf("Failed to initialize wine repository: %v", err)
    }

    // Create router
    r := mux.NewRouter()

    // Wine routes
    r.HandleFunc("/api/wines", handlers.GetWinesHandler(wineRepo)).Methods("GET")
    r.HandleFunc("/api/wines/{id}", handlers.GetWineByIDHandler(wineRepo)).Methods("GET")

    // CORS configuration
    c := cors.New(cors.Options{
        AllowedOrigins:   []string{"http://localhost:3000", "http://localhost:8080"},
        AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
        AllowedHeaders:   []string{"Accept", "Content-Type", "Authorization"},
        AllowCredentials: true,
    })

    // Create server
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    // Apply CORS middleware
    handler := c.Handler(r)

    // Start server
    fmt.Printf("Server starting on port %s\n", port)
    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), handler))
}