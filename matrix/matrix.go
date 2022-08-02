package matrix

import (
	"bytes"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

//Matrix type
type Matrix [][]float64

//NewMatrix constructor for a new Matrix
func NewMatrix(rows, columns int) Matrix {
	m := make([][]float64, rows)
	for i := 0; i < rows; i++ {
		m[i] = make([]float64, columns)
	}
	return m
}

//NewRandomizeMatrix constructor for a new Matrix with random values
func NewRandomizeMatrix(rows, columns int, start, end float64) Matrix {
	m := make([][]float64, rows)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < rows; i++ {
		m[i] = make([]float64, columns)
		for j := range m[i] {
			m[i][j] = rand.Float64()*(end-start) + start
		}
	}
	return m
}

//MatrixAdition add two materices
func MatrixAdition(a, b Matrix) (Matrix, error) {
	if (len(a) != len(b)) || (len(a[0]) != len(b[0])) {
		return nil, errors.New("No se pueden sumar matrices de tamaños diferentes")
	}
	m := make([][]float64, len(a))
	for i := range a {
		m[i] = make([]float64, len(b))
		for j := range a[i] {
			m[i][j] = a[i][j] + b[i][j]
		}
	}
	return Matrix(m), nil
}

//MatrixMultiplication multiply two matrices
func MatrixMultiplication(a, b Matrix) (Matrix, error) {
	if len(a[0]) != len(b) {

		return nil, errors.New("Los tamaños entre columnas y filas no coinciden")
	}
	m := make([][]float64, len(a))
	for i := range a {
		m[i] = make([]float64, len(b[0]))
		for j := range b[0] {
			for k := range b {
				m[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return Matrix(m), nil
}

//MatrixWiseMultiplication mutliply two matrices
func MatrixWiseMultiplication(a, b Matrix) (Matrix, error) {
	if len(a[0]) != len(b[0]) || len(a) != len(b) {
		return nil, errors.New("Los tamaños de las matrices no coinciden")
	}
	m := make([][]float64, len(a))
	for i := range a {
		m[i] = make([]float64, len(a[0]))
		for j := range a[0] {
			m[i][j] = a[i][j] * b[i][j]
		}
	}
	return Matrix(m), nil
}

//Map Apply a function to every element in the matrix given
func Map(a Matrix, f func(float64) float64) Matrix {
	m := make([][]float64, len(a))
	for i := range a {
		m[i] = make([]float64, len(a[i]))
		for j := range a[i] {
			m[i][j] = f(a[i][j])
		}
	}
	return m
}

//MatrixScalar scalar multiplication
func MatrixScalar(a Matrix, b float64) Matrix {
	m := make([][]float64, len(a))
	for i := range a {
		m[i] = make([]float64, len(a[0]))
		for j := range a[i] {
			m[i][j] = a[i][j] * b
		}
	}
	return Matrix(m)
}

//Transpose transpose operation
func Transpose(a Matrix) Matrix {
	if len(a) == 0 {
		return nil
	}
	m := NewMatrix(len(a[0]), len(a))
	for i := range a {
		for j := range a[0] {
			m[j][i] = a[i][j]
		}
	}
	return m
}

//Add add a funciton over themself
func (m *Matrix) Add(b Matrix) error {
	if (len(*m) != len(b)) || (len((*m)[0]) != len(b[0])) {
		return errors.New("No se pueden sumar matrices de tamaños diferentes")
	}
	for i := range b {
		for j := range b[i] {
			(*m)[i][j] += b[i][j]
		}
	}
	return nil
}

func (m Matrix) String() string {
	var buffer bytes.Buffer

	for i := range m {
		buffer.WriteString("\033[4m")
		for j := range m[0] {
			buffer.WriteString(fmt.Sprintf("%2.5f|", m[i][j]))
		}
		buffer.WriteString("\n")
	}
	buffer.WriteString("\033[0m")
	return buffer.String()
}
