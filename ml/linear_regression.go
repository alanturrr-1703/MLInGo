package ml

type LinearRegression struct {
	Weights []float64
	Bias    float64
}

func (lr *LinearRegression) Fit(X [][]float64, y []float64, iterations int, learningRate float64) {
	n := len(X)
	m := len(X[0])
	lr.Weights = make([]float64, m)
	lr.Bias = 0

	for iter := 0; iter < iterations; iter++ {
		predictions := lr.Predict(X)
		errors := make([]float64, n)
		for i := range errors {
			errors[i] = y[i] - predictions[i]
		}

		for j := 0; j < m; j++ {
			gradient := 0.0
			for i := 0; i < n; i++ {
				gradient += errors[i] * X[i][j]
			}
			lr.Weights[j] += learningRate * gradient / float64(n)
		}

		biasGradient := 0.0
		for i := 0; i < n; i++ {
			biasGradient += errors[i]
		}
		lr.Bias += learningRate * biasGradient / float64(n)
	}
}

func (lr *LinearRegression) Predict(X [][]float64) []float64 {
	n := len(X)
	predictions := make([]float64, n)
	for i := 0; i < n; i++ {
		predictions[i] = lr.Bias
		for j := 0; j < len(X[0]); j++ {
			predictions[i] += lr.Weights[j] * X[i][j]
		}
	}
	return predictions
}
