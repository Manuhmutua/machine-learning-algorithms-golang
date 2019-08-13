package main

import (
	"fmt"
	"github.com/gonum/floats"
	"github.com/kniren/gota/dataframe"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
	"log"
	"os"
)

func main() {
	// Open the driver dataset file.
	f, err := os.Open("fleet_data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Create a dataframe from the CSV file.
	driverDF := dataframe.ReadCSV(f)

	// Extract the distance column.
	yVals := driverDF.Col("Distance_Feature").Float()

	// clusterOne and clusterTwo will hold the values for plotting.
	var clusterOne [][]float64
	var clusterTwo [][]float64

	// Fill the clusters with data.
	for i, xVal := range driverDF.Col("Speeding_Feature").Float() {
		distanceOne := floats.Distance([]float64{yVals[i], xVal}, []float64{50.05, 8.83}, 2)
		distanceTwo := floats.Distance([]float64{yVals[i], xVal}, []float64{180.02, 18.29}, 2)
		if distanceOne < distanceTwo {
			clusterOne = append(clusterOne, []float64{xVal, yVals[i]})
			continue
		}
		clusterTwo = append(clusterTwo, []float64{xVal, yVals[i]})
	}

	// pts* will hold the values for plotting
	ptsOne := make(plotter.XYs, len(clusterOne))
	ptsTwo := make(plotter.XYs, len(clusterTwo))

	// Fill pts with data.
	for i, point := range clusterOne {
		ptsOne[i].X = point[0]
		ptsOne[i].Y = point[1]
	}

	for i, point := range clusterTwo {
		ptsTwo[i].X = point[0]
		ptsTwo[i].Y = point[1]
	}

	// Create the plot.
	p, err := plot.New()
	if err != nil {
		log.Fatal(err)
	}
	p.X.Label.Text = "Speeding"
	p.Y.Label.Text = "Distance"
	p.Add(plotter.NewGrid())

	sOne, err := plotter.NewScatter(ptsOne)
	if err != nil {
		log.Fatal(err)
	}
	sOne.GlyphStyle.Radius = vg.Points(3)
	sOne.GlyphStyle.Shape = draw.PyramidGlyph{}

	sTwo, err := plotter.NewScatter(ptsTwo)
	if err != nil {
		log.Fatal(err)
	}
	sTwo.GlyphStyle.Radius = vg.Points(3)

	// Save the plot to a PNG file.
	p.Add(sOne, sTwo)
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "fleet_data_clusters.png"); err != nil {
		log.Fatal(err)
	}

	// Output our within cluster metrics.
	fmt.Printf("\nCluster 1 Metric: %0.2f\n", withinClusterMean(clusterOne, []float64{50.05, 8.83}))
	fmt.Printf("\nCluster 2 Metric: %0.2f\n", withinClusterMean(clusterTwo, []float64{180.02, 18.29}))
}

// withinClusterMean calculates the mean distance between
// points in a cluster and the centroid of the cluster.
func withinClusterMean(cluster [][]float64, centroid []float64) float64 {

	// meanDistance will hold our result.
	var meanDistance float64

	// Loop over the points in the cluster.
	for _, point := range cluster {
		meanDistance += floats.Distance(point, centroid, 2) / float64(len(cluster))
	}

	return meanDistance
}

