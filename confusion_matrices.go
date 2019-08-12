package main

import (
	"fmt"
	"gonum.org/v1/gonum/integrate"
	"gonum.org/v1/gonum/stat"
)

func main() {
	// Define our scores and classes.
	scores := []float64{0.1, 0.35, 0.4, 0.8}
	classes := []bool{true, false, true, false}
	cutoffs := []float64{0.0, 0.0, 0.0, 0.0}

	// Calculate the true positive rates (recalls) and
	// false positive rates.
	tpr, fpr, _ := stat.ROC(cutoffs, scores, classes, nil)
	// Compute the Area Under Curve.
	auc := integrate.Trapezoidal(fpr, tpr)

	// Output the results to standard out.
	fmt.Printf("true positive rate: %v\n", tpr)
	fmt.Printf("false positive rate: %v\n", fpr)
	fmt.Printf("auc: %v\n", auc)
}
