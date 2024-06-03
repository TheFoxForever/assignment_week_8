package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"

	"gonum.org/v1/gonum/stat"
)

func generateSample(sampleSize int, mean, sd float64, randNumGen *rand.Rand) []float64 {
	sample := make([]float64, sampleSize)
	for i := 0; i < sampleSize; i++ {
		sample[i] = randNumGen.NormFloat64()*sd + mean
	}
	return sample
}

func bootstrapSample(data []float64, rng *rand.Rand) []float64 {
	data_len := len(data)
	resample := make([]float64, data_len)
	for i := range resample {
		resample[i] = data[rng.Intn(data_len)]
	}
	return resample
}

func main() {
	B := 100
	sampleSize := 100
	popMean := 100.0
	popSD := 10.0

	randNumGen := rand.New(rand.NewSource(9999))

	thisSample := generateSample(sampleSize, popMean, popSD, randNumGen)

	bootstrapMeans := make([]float64, B)
	bootstrapMedians := make([]float64, B)

	for b := 0; b < B; b++ {
		BootSample := bootstrapSample(thisSample, randNumGen)
		sort.Float64s(BootSample)
		bootstrapMeans[b] = stat.Mean(BootSample, nil)
		bootstrapMedians[b] = stat.Quantile(0.5, stat.Empirical, BootSample, nil)
	}

	seMeanCLT := popSD / math.Sqrt(float64(sampleSize))
	seMean := stat.StdDev(bootstrapMeans, nil)
	seMedian := stat.StdDev(bootstrapMedians, nil)

	fmt.Printf("SE Mean from Central Limit Theorem: %.2f\n", seMeanCLT)
	fmt.Printf("SE Mean from Bootstrap Samples: %.2f\n", seMean)
	fmt.Printf("SE Median Bootstrap Samples: %.2f\n", seMedian)

}
