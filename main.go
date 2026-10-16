package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/israeloluwasegun293-stars/Music-Brand-Site-Project/config"
	"github.com/israeloluwasegun293-stars/Music-Brand-Site-Project/db"
	"github.com/israeloluwasegun293-stars/Music-Brand-Site-Project/handlers"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {

	addr := flag.String("addr", ":8080", "HTTP Network Usage")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO/t", log.Ldate|log.Ltime)

	errorLog := log.New(os.Stderr, "ERROR/t", log.Ldate|log.Ltime|log.Lshortfile)

	mux := http.NewServeMux()

	appConfig := &config.AppConfig{
		InfoLog:  infoLog,
		ErrorLog: errorLog,
	}

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	mux.Handle("/", handlers.Home(appConfig))
	mux.Handle("/about", handlers.About(appConfig))
	mux.Handle("/bookings", handlers.Bookings(appConfig))
	mux.Handle("/music", handlers.Music(appConfig))
	mux.Handle("/services", handlers.Services(appConfig))
	mux.Handle("/submit-booking", handlers.Submit(appConfig))

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	dbCOnn, err := db.InitDb()
	if err != nil {
		log.Fatal(err)
	}
	defer dbCOnn.Close()

	infoLog.Printf("Starting server at http://localhost%s", *addr)
	infoLog.Println("Database initialization sucessful!")
	err = srv.ListenAndServe()
	if err != nil {
		errorLog.Fatal(err)
	}
}
