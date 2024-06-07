package ml

func Transpose(matrix [][]float64) [][]float64 {
	rows := len(matrix)
	cols := len(matrix[0])
	transposedMatrix := make([][]float64, cols)
	for i := range transposedMatrix {
		transposedMatrix[i] = make([]float64, rows)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			transposedMatrix[j][i] = matrix[i][j]
		}
	}
	return transposedMatrix
}

func Multiply(matrix1, matrix2 [][]float64) [][]float64 {
	rows1 := len(matrix1)
	cols1 := len(matrix1[0])
	cols2 := len(matrix2[0])
	result := make([][]float64, rows1)
	for i := range result {
		result[i] = make([]float64, cols2)
	}

	for i := 0; i < rows1; i++ {
		for j := 0; j < cols2; j++ {
			for k := 0; k < cols1; k++ {
				result[i][j] += matrix1[i][k] * matrix2[k][j]
			}
		}
	}
	return result
}
