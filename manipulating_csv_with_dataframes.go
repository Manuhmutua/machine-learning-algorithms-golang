package main

import (
	"fmt"
	"github.com/kniren/gota/dataframe"
	"log"
	"os"
)

func main() {
	// Open the CSV file.
	irisFile, err := os.Open("iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer irisFile.Close()

	// Create a dataframe from the CSV file.
	// The types of the columns will be inferred.
	irisDF := dataframe.ReadCSV(irisFile)

	// As a sanity check, display the records to stdout.
	// Gota will format the dataframe for pretty printing.
	fmt.Println(irisDF)

	// Create a filter for the dataframe.
	filter := dataframe.F{
		Colname: "species",
		Comparator: "==",
		Comparando: "Iris-versicolor",
	}

	// Filter the dataframe to see only the rows where
	// the iris species is "Iris-versicolor".
	versicolorDF := irisDF.Filter(filter)
	if versicolorDF.Err != nil {
		log.Fatal(versicolorDF.Err)
	}

	// Filter the dataframe again, but only select out the
	// sepal_width and species columns.
	versicolorDF = irisDF.Filter(filter).Select([]string{"sepal_width", "species"})

	// Filter and select the dataframe again, but only display
	// the first three results.
	versicolorDF = irisDF.Filter(filter).Select([]string{"sepal_width", "species"}).Subset([]int{0, 1, 2})
}