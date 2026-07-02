package main

import (
	"log"
	"net/http"

	"github.com/israeloluwasegun293-stars/Music-Brand-Site-Project/db"
	"github.com/israeloluwasegun293-stars/Music-Brand-Site-Project/handlers"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("/about", handlers.About)
	mux.HandleFunc("/bookings", handlers.Bookings)
	mux.HandleFunc("/music", handlers.Music)
	mux.HandleFunc("/services", handlers.Services)
	mux.HandleFunc("/submit-booking", handlers.Submit)

	dbCOnn, err := db.InitDb()
	if err != nil {
		log.Fatal(err)
	}
	defer dbCOnn.Close()

	log.Println("Starting server at http://localhost:8080")
	log.Println("Database initialization sucessful!")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
