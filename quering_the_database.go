package main

import (
	"database/sql"
	"fmt"
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

	// Iterate over the rows, sending the results to
	// standard out.
	for rows.Next() {

		var (
			sLength float64
			sWidth float64
			pLength float64
			pWidth float64
		)

		if err := rows.Scan(&sLength, &sWidth, &pLength, &pWidth); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%.2f, %.2f, %.2f, %.2f\n", sLength, sWidth, pLength, pWidth)
	}

	// Check for errors after we are done iterating over rows.
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
