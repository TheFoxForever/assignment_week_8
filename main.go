package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"

	"gonum.org/v1/gonum/stat"
)

func generateSample(sampleSize int, mean, sd float64, rng *rand.Rand) []float64 {
	sample := make([]float64, sampleSize)
	for i := 0; i < sampleSize; i++ {
		sample[i] = rng.NormFloat64()*sd + mean
	}
	return sample
}

func bootstrapSample(data []float64, rng *rand.Rand) []float64 {
	n := len(data)
	resample := make([]float64, n)
	for i := range resample {
		resample[i] = data[rng.Intn(n)]
	}
	return resample
}

func main() {
	B := 100
	sampleSize := 100
	popMean := 100.0
	popSD := 10.0

	rng := rand.New(rand.NewSource(9999)) // For reproducible results

	// Generate a single sample
	thisSample := generateSample(sampleSize, popMean, popSD, rng)

	bootstrapMeans := make([]float64, B)
	bootstrapMedians := make([]float64, B)

	for b := 0; b < B; b++ {
		thisBootstrapSample := bootstrapSample(thisSample, rng)
		sort.Float64s(thisBootstrapSample)
		bootstrapMeans[b] = stat.Mean(thisBootstrapSample, nil)
		bootstrapMedians[b] = stat.Quantile(0.5, stat.Empirical, thisBootstrapSample, nil)
	}

	seMeanCLT := popSD / math.Sqrt(float64(sampleSize))
	seMean := stat.StdDev(bootstrapMeans, nil)
	seMedian := stat.StdDev(bootstrapMedians, nil)

	fmt.Printf("SE Mean from Central Limit Theorem: %.2f\n", seMeanCLT)
	fmt.Printf("SE Mean from Bootstrap Samples: %.2f\n", seMean)
	fmt.Printf("SE Median Bootstrap Samples: %.2f\n", seMedian)

}
