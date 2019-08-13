package main

import (
	"fmt"
	"github.com/kniren/gota/dataframe"
	"github.com/sajari/regression"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"log"
	"os"
	"strconv"
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
	fmt.Println("Partial Autocorrelation:")
	for i := 1; i < 10; i++ {

		// Calculate the partial autocorrelation.
		pac := pacf(passengers, i)
		fmt.Printf("Lag %d period: %0.2f\n", i, pac)
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
		pts[i-1] = pacf(passengers, i)
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
	if err := p.Save(8*vg.Inch, 4*vg.Inch, "p_acf.png"); err != nil {
		log.Fatal(err)
	}

}

// pacf calculates the partial autocorrelation for a series
// at the given lag.
func pacf(x []float64, lag int) float64 {

	// Create a regresssion.Regression value needed to train
	// a model using github.com/sajari/regression.
	var r regression.Regression
	r.SetObserved("x")

	// Define the current lag and all of the intermediate lags.
	for i := 0; i < lag; i++ {
		r.SetVar(i, "x"+strconv.Itoa(i))
	}

	// Shift the series.
	xAdj := append(x[lag:len(x)])

	// Loop over the series creating the data set
	// for the regression.
	for i, xVal := range xAdj {

		// Loop over the intermediate lags to build up
		// our independent variables.
		laggedVariables := make([]float64, lag)
		for idx := 1; idx <= lag; idx++ {

			// Get the lagged series variables.
			laggedVariables[idx-1] = x[lag+i-idx]
		}

		// Add these points to the regression value.
		r.Train(regression.DataPoint(xVal, laggedVariables))
	}

	// Fit the regression.
	r.Run()

	return r.Coeff(lag)
}