// main.go
package main

import (
	"log"
	"net/http"
	"projet/server"
)

const port = ":1717"

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", server.WineHandler) // Route de la page d'accueil

	log.Println("Démarrage du serveur http://localhost", port)
	err := http.ListenAndServe(port, nil) // Lancer le serveur sur le port 8080
	if err != nil {
		log.Fatal("Erreur lors du démarrage du serveur : ", err)
	}
}
