package main

import (
	"fmt"
	"github.com/kniren/gota/dataframe"
	"github.com/sjwhitworth/golearn/evaluation"
	"log"
	"math"
	"os"
)

func main() {

	// Open the diabetes dataset file.
	f, err := os.Open("diabetes.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Create a dataframe from the CSV file.
	// The types of the columns will be inferred.
	diabetesDF := dataframe.ReadCSV(f)

	// Define the decision tree model.
	tree := trees.NewID3DecisionTree(param)

	// Perform the cross validation.
	cfs, err := evaluation.GenerateCrossFoldValidationConfusionMatrices(diabetesDF, tree, 5)
	if err != nil {
		panic(err)
	}

	// Calculate the metrics.
	mean, variance := evaluation.GetCrossValidatedMetric(cfs, evaluation.GetAccuracy)
	stdev := math.Sqrt(variance)

	// Output the results to standard out.
	fmt.Printf("%0.2f\t\t%.2f (+/- %.2f)\n", param, mean, stdev*2)
}
