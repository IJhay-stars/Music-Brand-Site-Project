package handlers

import (
	"net/http"

	"github.com/israeloluwasegun293-stars/Music-Brand-Site-Project/config"
	"github.com/israeloluwasegun293-stars/Music-Brand-Site-Project/db"
)

type Info struct {
	Fname string
	Lname string
	Email string
	Date  string
}

func Submit(app *config.AppConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/submit-booking" {
			http.NotFound(w, r)
			return
		}

		if r.Method != http.MethodPost {
			w.Header().Set("Allow", http.MethodPost)

			http.Error(w, "Method not allowed", 405)
			return
		}
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Bad Request", 400)
			return
		}

		var s Info

		s.Fname = r.FormValue("Firstname")
		s.Lname = r.FormValue("Lastname")
		s.Email = r.FormValue("Email")
		s.Date = r.FormValue("Date")

		b := `INSERT INTO bookings(FirstName, LastName, Email, Date)VALUES(?,?,?,?);`

		_, err = db.DB.Exec(b, s.Fname, s.Lname, s.Email, s.Date)
		if err != nil {
			http.Error(w, "Internal server error", 500)
			return
		}
		w.Write([]byte("Booking submitted sucessfully!"))
	}
}
