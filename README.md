# Linear Regression in Go

This repository contains an implementation of a linear regression model using gradient descent in Go. The model is designed to predict a target variable based on one or more input features.

## Overview

Linear regression is a fundamental machine learning algorithm used for predicting a continuous target variable based on one or more input features. This implementation includes:

- A `LinearRegression` struct to encapsulate the model's weights and bias.
- Methods to fit the model to training data using gradient descent.
- Methods to make predictions on new data.

## Features

- **Gradient Descent Optimization**: The model is trained using the gradient descent algorithm, an iterative optimization technique to minimize the mean squared error.
- **From Scratch Implementation**: All matrix operations and gradient calculations are implemented from scratch, providing a clear understanding of the underlying mathematics.

## Project Structure


- `main.go`: Demonstrates how to use the `LinearRegression` model with example data.
- `ml/linear_regression.go`: Contains the implementation of the `LinearRegression` model.

## Usage

### 1. Clone the Repository

```bash
git clone https://github.com/your-username/ml_project.git
cd ml_project
ml_project/
│
├── main.go
└── ml/
    └── linear_regression.go
package main

import (
    "fmt"
    "ml"
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
    fmt.Println("Weights:", lr.Weights)
    fmt.Println("Bias:", lr.Bias)
}
