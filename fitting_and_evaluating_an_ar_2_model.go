package main

import (
	"encoding/csv"
	"fmt"
	"github.com/kniren/gota/dataframe"
	"github.com/sajari/regression"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {

	// Open the CSV file.
	passengersFile, err := os.Open("diff_series.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer passengersFile.Close()

	// Create a dataframe from the CSV file.
	passengersDF := dataframe.ReadCSV(passengersFile)

	// Get the time and passengers as a slice of floats.
	passengers := passengersDF.Col("differenced_passengers").Float()

	// Calculate the coefficients for lag 1 and 2 and
	// our error.
	coeffs, intercept := autoregressive(passengers, 2)

	// Output the AR(2) model to stdout.
	fmt.Printf("\nlog(x(t)) - log(x(t-1)) = %0.6f + lag1*%0.6f + lag2*%0.6f\n\n", intercept, coeffs[0], coeffs[1])

	//////////

	// Open the log differenced dataset file.
	transFile, err := os.Open("diff_series.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer transFile.Close()

	// Create a CSV reader reading from the opened file.
	transReader := csv.NewReader(transFile)

	// Read in all of the CSV records
	transReader.FieldsPerRecord = 2
	transData, err := transReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Loop over the data predicting the transformed
	// observations.
	var transPredictions []float64
	for i, _ := range transData {

		// Skip the header and the first two observations
		// (because we need two lags to make a prediction).
		if i == 0 || i == 1 || i == 2 {
			continue
		}

		// Parse the first lag.
		lagOne, err := strconv.ParseFloat(transData[i-1][1], 64)
		if err != nil {
			log.Fatal(err)
		}

		// Parse the second lag.
		lagTwo, err := strconv.ParseFloat(transData[i-2][1], 64)
		if err != nil {
			log.Fatal(err)
		}

		// Predict the transformed variable with our trained AR model.
		transPredictions = append(transPredictions, 0.008159+0.234953*lagOne-0.173682*lagTwo)
	}

	// Open the original dataset file.
	origFile, err := os.Open("AirPassengers.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer origFile.Close()

	// Create a CSV reader reading from the opened file.
	origReader := csv.NewReader(origFile)

	// Read in all of the CSV records
	origReader.FieldsPerRecord = 2
	origData, err := origReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// pts* will hold the values for plotting.
	ptsObs := make(plotter.XYs, len(transPredictions))
	ptsPred := make(plotter.XYs, len(transPredictions))

	// Reverse the transformation and calculate the MAE.
	var mAE float64
	var cumSum float64
	for i := 4; i <= len(origData)-1; i++ {

		// Parse the original observation.
		observed, err := strconv.ParseFloat(origData[i][1], 64)
		if err != nil {
			log.Fatal(err)
		}

		// Parse the original date.
		date, err := strconv.ParseFloat(origData[i][0], 64)
		if err != nil {
			log.Fatal(err)
		}

		// Get the cumulative sum up to the index in
		// the transformed predictions.
		cumSum += transPredictions[i-4]

		// Calculate the reverse transformed prediction.
		predicted := math.Exp(math.Log(observed) + cumSum)

		// Accumulate the MAE.
		mAE += math.Abs(observed-predicted) / float64(len(transPredictions))

		// Fill in the points for plotting.
		ptsObs[i-4].X = date
		ptsPred[i-4].X = date
		ptsObs[i-4].Y = observed
		ptsPred[i-4].Y = predicted
	}

	// Output the MAE to standard out.
	fmt.Printf("\nMAE = %0.2f\n\n", mAE)

	// Create the plot.
	p, err := plot.New()
	if err != nil {
		log.Fatal(err)
	}
	p.X.Label.Text = "time"
	p.Y.Label.Text = "passengers"
	p.Add(plotter.NewGrid())

	// Add the line plot points for the time series.
	lObs, err := plotter.NewLine(ptsObs)
	if err != nil {
		log.Fatal(err)
	}
	lObs.LineStyle.Width = vg.Points(1)

	lPred, err := plotter.NewLine(ptsPred)
	if err != nil {
		log.Fatal(err)
	}
	lPred.LineStyle.Width = vg.Points(1)
	lPred.LineStyle.Dashes = []vg.Length{vg.Points(5), vg.Points(5)}

	// Save the plot to a PNG file.
	p.Add(lObs, lPred)
	p.Legend.Add("Observed", lObs)
	p.Legend.Add("Predicted", lPred)
	if err := p.Save(10*vg.Inch, 4*vg.Inch, "passengers_ts.png"); err != nil {
		log.Fatal(err)
	}

}

// autoregressive calculates an AR model for a series
// at a given order.
func autoregressive(x []float64, lag int) ([]float64, float64) {

	// Create a regresssion.Regression value needed to train
	// a model using github.com/sajari/regression.
	var r regression.Regression
	r.SetObserved("x")

	// Define the current lag and all of the intermediate lags.
	for i := 0; i < lag; i++ {
		r.SetVar(i, "x"+strconv.Itoa(i))
	}

	// Shift the series.
	xAdj := x[lag:len(x)]

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

	// coeff hold the coefficients for our lags.
	var coeff []float64
	for i := 1; i <= lag; i++ {
		coeff = append(coeff, r.Coeff(i))
	}

	return coeff, r.Coeff(0)
}
