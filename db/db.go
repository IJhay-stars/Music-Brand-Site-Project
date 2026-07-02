package db

import (
	"database/sql"
	"log"

	_"modernc.org/sqlite"
)

	var DB *sql.DB
func InitDb() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "Music-Brand.db")
	if err != nil {
		log.Println("Error Opening the Database")
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		log.Default().Println("Error the db file")
		return nil, err
	}

	query := `CREATE TABLE IF NOT EXISTS bookings (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    FirstName TEXT,
	LastName  TEXT,
	Email  TEXT,
	Date  TEXT
);`

	_, err = db.Exec(query)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	DB = db

	return db, nil
}
