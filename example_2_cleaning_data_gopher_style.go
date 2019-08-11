package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

func main() {
	// Open the iris dataset file.
	f, err := os.Open("iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Create a new CSV reader reading from the opened file.
	//reader := csv.NewReader(f)

	// Assume we don't know the number of fields per line. By setting
	// FieldsPerRecord negative, each row may have a variable
	// number of fields.
	//reader.FieldsPerRecord = -1
	//
	//// Read in all of the CSV records.
	//rawCSVData, err := reader.ReadAll()
	//if err != nil {
	//	log.Fatal(err)
	//}

	// Create a new CSV reader reading from the opened file.
	reader := csv.NewReader(f)
	reader.FieldsPerRecord = -1

	// rawCSVData will hold our successfully parsed rows.
	var rawCSVData [][]string

	// Read in the records one by one.
	for {

		// Read in a row. Check if we are at the end of the file.
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		// Append the record to our dataset.
		rawCSVData = append(rawCSVData, record)
	}



}
