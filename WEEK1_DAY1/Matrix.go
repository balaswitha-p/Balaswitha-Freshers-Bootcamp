package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Matrix struct {
	Rows     int
	Cols     int
	Elements [][]float64
}

func NewMatrix(rows, cols int) (*Matrix, error) {
	if rows <= 0 || cols <= 0 {
		return nil, errors.New("matrix dimensions must be positive")
	}
	elements := make([][]float64, rows)
	for i := range elements {
		elements[i] = make([]float64, cols)
	}
	return &Matrix{Rows: rows, Cols: cols, Elements: elements}, nil
}

func (m *Matrix) GetRows() int {
	return m.Rows
}

func (m *Matrix) GetCols() int {
	return m.Cols
}

func (m *Matrix) Set(i, j int, value float64) error {
	if i < 0 || i >= m.Rows || j < 0 || j >= m.Cols {
		return errors.New("index out of bounds")
	}
	m.Elements[i][j] = value
	return nil
}

func (m *Matrix) Add(other *Matrix) (*Matrix, error) {
	if m.Rows != other.Rows || m.Cols != other.Cols {
		return nil, errors.New("matrices must have the same dimensions for addition")
	}

	result, err := NewMatrix(m.Rows, m.Cols)
	if err != nil {
		return nil, err
	}

	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Cols; j++ {
			result.Elements[i][j] = m.Elements[i][j] + other.Elements[i][j]
		}
	}
	return result, nil
}

func (m *Matrix) PrintJSON() error {
	jsonOutput, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshalling matrix to JSON: %w", err)
	}
	fmt.Println(string(jsonOutput))
	return nil
}

func (m *Matrix) PrintMatrix() {
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Cols; j++ {
			fmt.Printf("%8.2f ", m.Elements[i][j])
		}
		fmt.Println()
	}
}

func main() {
	matrixA, err := NewMatrix(2, 3)
	if err != nil {
		fmt.Println("Error creating matrix A:", err)
		return
	}
	matrixA.Set(0, 0, 1.0)
	matrixA.Set(0, 1, 2.0)
	matrixA.Set(0, 2, 3.0)
	matrixA.Set(1, 0, 4.0)
	matrixA.Set(1, 1, 5.0)
	matrixA.Set(1, 2, 6.0)

	fmt.Println("Matrix A:")
	matrixA.PrintMatrix()
	fmt.Println("Rows:", matrixA.GetRows(), "Cols:", matrixA.GetCols())

	matrixB, err := NewMatrix(2, 3)
	if err != nil {
		fmt.Println("Error creating matrix B:", err)
		return
	}
	matrixB.Set(0, 0, 7.0)
	matrixB.Set(0, 1, 8.0)
	matrixB.Set(0, 2, 9.0)
	matrixB.Set(1, 0, 10.0)
	matrixB.Set(1, 1, 11.0)
	matrixB.Set(1, 2, 12.0)

	fmt.Println("Matrix B:")
	matrixB.PrintMatrix()

	matrixC, err := matrixA.Add(matrixB)
	if err != nil {
		fmt.Println("Error adding matrices:", err)
	} else {
		fmt.Println("Matrix C (A + B):")
		matrixC.PrintMatrix()
	}

	matrixD, _ := NewMatrix(2, 2)
	_, err = matrixA.Add(matrixD)
	if err != nil {
		fmt.Println("Attempt to add matrices with different dimensions (expected error):", err)
	}

	fmt.Println("Matrix A as JSON:")
	matrixA.PrintJSON()
}
