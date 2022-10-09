package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/impin2rex/go_microservices/details"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving Root Route")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Application is up and & running")
}

func detailsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Fetching the details")
	hostname, err := details.GetHostName()
	if err != nil {
		panic(err)
	}
	IP, _ := details.GetIP()
	response := map[string]string{
		"hostname": hostname,
		"ip":       IP.String(),
	}
	json.NewEncoder(w).Encode(response)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Checking application health")
	response := map[string]string{
		"status":    "UP",
		"timestamp": time.Now().String(),
	}
	json.NewEncoder(w).Encode(response)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", rootHandler)
	r.HandleFunc("/health", healthHandler)
	r.HandleFunc("/details", detailsHandler)
	log.Printf("Server has started at http://localhost:4000")
	log.Fatal(http.ListenAndServe(":4000", r))
}
