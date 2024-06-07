package main

import (
	"MLFromScratch/ml"
	"fmt"
)

func main() {
	// Example data
	X := [][]float64{
		{1, 1},
		{2, 2},
		{3, 3},
	}
	y := []float64{2, 4, 6}

	// Create and fit the model
	lr := &ml.LinearRegression{}
	lr.Fit(X, y, 1000, 0.01)

	// Make predictions
	predictions := lr.Predict(X)
	fmt.Println("Predictions:", predictions)
	fmt.Println("Coefficients:", lr.Weights)
	fmt.Println("Bias:", lr.Bias)
}
