package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"log"
	"math"
)

func main() {

	// Create two matrices of the same size, a and b.
	a := mat.NewDense(3, 3, []float64{1, 2, 3, 0, 4, 5, 0, 0, 6})
	b := mat.NewDense(3, 3, []float64{8, 9, 10, 1, 4, 2, 9, 0, 2})

	// Create a third matrix of a different size.
	c := mat.NewDense(3, 2, []float64{3, 2, 1, 4, 0, 8})

	// Add a and b.
	d := mat.NewDense(3, 3, nil)
	d.Add(a, b)
	fd := mat.Formatted(d, mat.Prefix("            "))
	fmt.Printf("d = a + b = %0.4v\n\n", fd)

	// Multiply a and c.
	f := mat.NewDense(3, 2, nil)
	f.Mul(a, c)
	ff := mat.Formatted(f, mat.Prefix("          "))
	fmt.Printf("f = a c = %0.4v\n\n", ff)

	// Raising a matrix to a power.
	g := mat.NewDense(3, 3, nil)
	g.Pow(a, 5)
	fg := mat.Formatted(g, mat.Prefix("          "))
	fmt.Printf("g = a^5 = %0.4v\n\n", fg)

	// Apply a function to each of the elements of a.
	h := mat.NewDense(3, 3, nil)
	sqrt := func(_, _ int, v float64) float64 { return math.Sqrt(v) }
	h.Apply(sqrt, a)
	fh := mat.Formatted(h, mat.Prefix("              "))
	fmt.Printf("h = sqrt(a) = %0.4v\n\n", fh)

	// Create a new matrix a.
	a = mat.NewDense(3, 3, []float64{1, 2, 3, 0, 4, 5, 0, 0, 6})

	// Compute and output the transpose of the matrix.
	ft := mat.Formatted(a.T(), mat.Prefix(" "))
	fmt.Printf("a^T = %v\n\n", ft)

	// Compute and output the determinant of a.
	deta := mat.Det(a)
	fmt.Printf("det(a) = %.2f\n\n", deta)

	// Compute and output the inverse of a.
	aInverse := mat.NewDense(3, 3, nil)
	if err := aInverse.Inverse(a); err != nil {
		log.Fatal(err)
	}
	fi := mat.Formatted(aInverse, mat.Prefix(" "))
	fmt.Printf("a^-1 = %v\n\n", fi)

}