package ml

import (
	"math"
	"sort"
)

func Mean(data []float64) float64 {
	sum := 0.0
	for _, num := range data {
		sum += num
	}
	return sum / float64(len(data))
}

func Median(data []float64) float64 {
	sortedData := make([]float64, len(data))
	copy(sortedData, data)
	sort.Float64s(sortedData)
	middle := len(sortedData) / 2
	if len(sortedData)%2 == 0 {
		return (sortedData[middle-1] + sortedData[middle]) / 2.0
	}
	return sortedData[middle]
}

func Variance(data []float64) float64 {
	mean := Mean(data)
	temp := 0.0
	for _, num := range data {
		temp += (num - mean) * (num - mean)
	}
	return temp / float64(len(data))
}

func StandardDeviation(data []float64) float64 {
	return math.Sqrt(Variance(data))
}
