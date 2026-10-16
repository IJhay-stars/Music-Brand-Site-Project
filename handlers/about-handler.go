package handlers

import (
	"html/template"
	"net/http"

	"github.com/israeloluwasegun293-stars/Music-Brand-Site-Project/config"
)

func About(app *config.AppConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/about" {
			http.NotFound(w, r)
			return
		}

		if r.Method != http.MethodGet {
			w.Header().Set("Allow", http.MethodGet)

			http.Error(w, "Method not allowed", 405)
			return
		}
		temp, err := template.ParseFiles("templates/about.html")
		if err != nil {
			http.Error(w, "internal server error", 500)
			return
		}

		temp.Execute(w, nil)
	}
}
