package main

import (
	"fmt"
	"math/rand"

	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/gonum/stat/sampleuv"
)

func main() {
	rand.Seed(123)

	data := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	numSamples := 1000

	medians := make([]float64, numSamples)

	for i := 0; i < numSamples; i++ {
		resample := sampleuv.Bootstrap(nil, data, len(data), func() Float64)
		medians[i] = stat.Quantile(0.5, stat.Empirical, resample, nil)
	}

	mean, std := stat.MeanStdDev(medians, nil)
	fmt.Println("Mean of medians:", mean)
	fmt.Println("Standard error of the median:", std)

}
