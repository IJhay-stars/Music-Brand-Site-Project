package handlers

import (
	"html/template"
	"net/http"

	"github.com/israeloluwasegun293-stars/Music-Brand-Site-Project/config"
)

func Home(app *config.AppConfig) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		if r.Method != http.MethodGet {
			w.Header().Set("Allow", http.MethodGet)

			http.Error(w, "Method not allowd", 405)
			return
		}

		temp, err := template.ParseFiles("templates/index.html")
		if err != nil {
			http.Error(w, "internal server error", 500)
			return
		}
		temp.Execute(w, nil)
	}
}
