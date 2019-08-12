package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

func main() {
	// Initialize a "vector" via a slice.
	//var myvector []float64

	// Add a couple of components to the vector.
	//myvector = append(myvector, 11.0)
	//myvector = append(myvector, 5.2)

	// Create a new vector value.
	myvector := mat.NewVecDense(2, []float64{11.0, 5.2})

	// Output the results to stdout.
	fmt.Println(myvector)
}
