package main

import (
	"errors"
	"math/rand"
	"sync"
)

func createMatrix(rows int, cols int, random bool) [][]int {
	matrix := make([][]int, rows)

	for i := range matrix {
		matrix[i] = make([]int, cols)

		for j := range matrix[0] {
			if random {
				matrix[i][j] = rand.Intn(10)
			} else {
				matrix[i][j] = 0
			}
		}
	}

	return matrix
}

// Function to check if matrix multiplication is legal
func checkMultLegal(mat_a [][]int, mat_b [][]int) bool {
	// Ensure rows >= 1
	if len(mat_a) < 1 || len(mat_b) < 1 {
		return false
	}

	// Ensure cols >= 1
	if len(mat_a[0]) < 1 || len(mat_b[0]) < 1 {
		return false
	}

	// Ensure cols A == rows B
	if len(mat_a[0]) != len(mat_b) {
		return false
	}

	return true
}

// Function to multiply two matricies on a single thread
func multiplySt(mat_a [][]int, mat_b [][]int) (mat_c [][]int, err error) {
	mat_c = make([][]int, len(mat_a))
	err = nil

	if !checkMultLegal(mat_a, mat_b) {
		err = errors.New("matrix multiplication illegal")
		return mat_c, err
	}

	for i := 0; i < len(mat_c); i++ {
		mat_c[i] = make([]int, len(mat_b[0]))

		for j := 0; j < len(mat_c[0]); j++ {
			func(row int, col int) {
				dot := 0
				for k := 0; k < len(mat_a[0]); k++ {
					dot += mat_a[row][k] * mat_b[k][col]
				}
				mat_c[row][col] = dot
			}(i, j)
		}
	}

	return mat_c, err
}

// Function to multiply two matricies using multiple threads (Go routines)
func multiplyMt(mat_a [][]int, mat_b [][]int) (mat_c [][]int, err error) {
	mat_c = make([][]int, len(mat_a))
	err = nil

	if !checkMultLegal(mat_a, mat_b) {
		err = errors.New("matrix multiplication illegal")
		return mat_c, err
	}

	var wg sync.WaitGroup

	for i := 0; i < len(mat_c); i++ {
		mat_c[i] = make([]int, len(mat_b[0]))

		for j := 0; j < len(mat_c[0]); j++ {
			wg.Add(1)

			go func(row int, col int) {
				defer wg.Done()

				dot := 0
				for k := 0; k < len(mat_a[0]); k++ {
					dot += mat_a[row][k] * mat_b[k][col]
				}
				mat_c[row][col] = dot
			}(i, j)
		}
	}

	// make sure all calculations are finished before returning
	wg.Wait()

	return mat_c, err
}
