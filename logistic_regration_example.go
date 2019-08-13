package main

import (
	"encoding/csv"
	"fmt"
	"github.com/kniren/gota/dataframe"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	scoreMax = 830.0
	scoreMin = 640.0
)

func main() {

	// Open the loan dataset file.
	f, err := os.Open("loan_data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Create a new CSV reader reading from the opened file.
	reader := csv.NewReader(f)
	reader.FieldsPerRecord = 2

	// Read in all of the CSV records
	rawCSVData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Create the output file.
	f, err = os.Create("clean_loan_data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Create a CSV writer.
	w := csv.NewWriter(f)

	// Sequentially move the rows writing out the parsed values.
	for idx, record := range rawCSVData {

		// Skip the header row.
		if idx == 0 {

			// Write the header to the output file.
			if err := w.Write(record); err != nil {
				log.Fatal(err)
			}
			continue
		}

		// Initialize a slice to hold our parsed values.
		outRecord := make([]string, 2)

		// Parse and standardize the FICO score.
		score, err := strconv.ParseFloat(strings.Split(record[0], "-")[0], 64)
		if err != nil {
			log.Fatal(err)
		}

		outRecord[0] = strconv.FormatFloat((score-scoreMin)/(scoreMax-scoreMin), 'f', 4, 64)

		// Parse the Interest rate class.
		rate, err := strconv.ParseFloat(strings.TrimSuffix(record[1], "%"), 64)
		if err != nil {
			log.Fatal(err)
		}

		if rate <= 12.0 {
			outRecord[1] = "1.0"

			// Write the record to the output file.
			if err := w.Write(outRecord); err != nil {
				log.Fatal(err)
			}
			continue
		}

		outRecord[1] = "0.0"

		// Write the record to the output file.
		if err := w.Write(outRecord); err != nil {
			log.Fatal(err)
		}
	}

	// Write any buffered data to the underlying writer (standard output).
	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}

	// Open the CSV file.
	loanDataFile, err := os.Open("clean_loan_data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer loanDataFile.Close()

	// Create a dataframe from the CSV file.
	loanDF := dataframe.ReadCSV(loanDataFile)

	// Use the Describe method to calculate summary statistics
	// for all of the columns in one shot.
	loanSummary := loanDF.Describe()

	// Output the summary statistics to stdout.
	fmt.Println(loanSummary)

	// Create a histogram for each of the columns in the dataset.
	for _, colName := range loanDF.Names() {

		// Create a plotter.Values value and fill it with the
		// values from the respective column of the dataframe.
		plotVals := make(plotter.Values, loanDF.Nrow())
		for i, floatVal := range loanDF.Col(colName).Float() {
			plotVals[i] = floatVal
		}

		// Make a plot and set its title.
		p, err := plot.New()
		if err != nil {
			log.Fatal(err)
		}
		p.Title.Text = fmt.Sprintf("Histogram of a %s", colName)

		// Create a histogram of our values.
		h, err := plotter.NewHist(plotVals, 16)
		if err != nil {
			log.Fatal(err)
		}

		// Normalize the histogram.
		h.Normalize(1)

		// Add the histogram to the plot.
		p.Add(h)

		// Save the plot to a PNG file.
		if err := p.Save(4*vg.Inch, 4*vg.Inch, colName+"_hist.png"); err != nil {
			log.Fatal(err)
		}
	}
}
