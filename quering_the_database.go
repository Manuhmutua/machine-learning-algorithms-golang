package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func main() {
	// Get the postgres connection URL. I have it stored in
	// an environmental variable.
	pgURL := os.Getenv("PGURL")
	if pgURL == "" {
		log.Fatal("PGURL empty")
	}

	// Open a database value. Specify the postgres driver
	// for databases/sql.
	db, err := sql.Open("postgres", pgURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	// Query the database.
	rows, err := db.Query(`
    SELECT 
        sepal_length as sLength, 
        sepal_width as sWidth, 
        petal_length as pLength, 
        petal_width as pWidth 
    FROM iris
    WHERE species = $1`, "Iris-setosa")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
}
