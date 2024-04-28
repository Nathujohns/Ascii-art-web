package main

import (
	"fmt"
	"log"
	"net/http"

	"ascii/src/server"
)

func main() {
	serverAddress := ":8080" // Define server port as a variable
	fmt.Println("Server is active at http://localhost" + serverAddress + ". Press Ctrl+C to terminate.")

	// Setting up the server handlers
	http.HandleFunc("/", server.HomePageHandler)          // Updated to HomePageHandler
	http.HandleFunc("/ascii-art", server.ASCIIArtHandler) // Updated to ASCIIArtHandler
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	// Start the HTTP server and log if there's a fatal error
	log.Fatal(http.ListenAndServe(serverAddress, nil))
}
