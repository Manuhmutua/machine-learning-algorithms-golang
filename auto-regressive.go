package main

import (
	"fmt"
	"github.com/kniren/gota/dataframe"
	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"log"
	"math"
	"os"
)

func main() {
	// Open the CSV file.
	passengersFile, err := os.Open("AirPassengers.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer passengersFile.Close()

	// Create a dataframe from the CSV file.
	passengersDF := dataframe.ReadCSV(passengersFile)

	// Get the time and passengers as a slice of floats.
	passengers := passengersDF.Col("AirPassengers").Float()

	// Loop over various values of lag in the series.
	fmt.Println("Autocorrelation:")
	for i := 1; i < 10; i++ {

		// Shift the series.
		adjusted := passengers[i:len(passengers)]
		lag := passengers[0 : len(passengers)-i]

		// Calculate the autocorrelation.
		ac := stat.Correlation(adjusted, lag, nil)
		fmt.Printf("Lag %d period: %0.2f\n", i, ac)
	}

	// Create a new plot, to plot our autocorrelations.
	p, err := plot.New()
	if err != nil {
		log.Fatal(err)
	}

	p.Title.Text = "Autocorrelations for AirPassengers"
	p.X.Label.Text = "Lag"
	p.Y.Label.Text = "ACF"
	p.Y.Min = 0
	p.Y.Max = 1

	w := vg.Points(3)

	// Create the points for plotting.
	numLags := 10
	pts := make(plotter.Values, numLags)

	// Loop over various values of lag in the series.
	for i := 1; i < numLags; i++ {

		// Calculate the autocorrelation.
		pts[i-1] = acf(passengers, i)
	}

	// Add the points to the plot.
	bars, err := plotter.NewBarChart(pts, w)
	if err != nil {
		log.Fatal(err)
	}
	bars.LineStyle.Width = vg.Length(0)
	bars.Color = plotutil.Color(1)

	// Save the plot to a PNG file.
	p.Add(bars)
	if err := p.Save(8*vg.Inch, 4*vg.Inch, "acf.png"); err != nil {
		log.Fatal(err)
	}
}

// acf calculates the autocorrelation for a series
// at the given lag.
func acf(x []float64, lag int) float64 {

	// Shift the series.
	xAdj := x[lag:]
	xLag := x[0 : len(x)-lag]

	// numerator will hold our accumulated numerator, and
	// denominator will hold our accumulated denominator.
	var numerator float64
	var denominator float64

	// Calculate the mean of our x values, which will be used
	// in each term of the autocorrelation.
	xBar := stat.Mean(x, nil)

	// Calculate the numerator.
	for idx, xVal := range xAdj {
		numerator += (xVal - xBar) * (xLag[idx] - xBar)
	}

	// Calculate the denominator.
	for _, xVal := range x {
		denominator += math.Pow(xVal-xBar, 2)
	}

	return numerator / denominator
}
