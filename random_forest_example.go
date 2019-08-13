package main

import (
	"fmt"
	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/ensemble"
	"github.com/sjwhitworth/golearn/evaluation"
	"log"
	"math"
)

func main() {
	// Read in the iris data set into golearn "instances".
	irisData, err := base.ParseCSVToInstances("iris.csv", true)
	if err != nil {
		log.Fatal(err)
	}

	// Assemble a random forest with 10 trees and 2 features per tree,
	// which is a sane default (number of features per tree is normally set
	// to sqrt(number of features)).
	rf := ensemble.NewRandomForest(10, 2)

	// Use cross-fold validation to successively train and evaluate the model
	// on 5 folds of the data set.
	cv, err := evaluation.GenerateCrossFoldValidationConfusionMatrices(irisData, rf, 5)
	if err != nil {
		log.Fatal(err)
	}

	// Get the mean, variance and standard deviation of the accuracy for the
	// cross validation.
	mean, variance := evaluation.GetCrossValidatedMetric(cv, evaluation.GetAccuracy)
	stdev := math.Sqrt(variance)

	// Output the cross metrics to standard out.
	fmt.Printf("\nAccuracy\n%.2f (+/- %.2f)\n\n", mean, stdev*2)
}