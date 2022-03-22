package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	// create matrix A
	mat_a_rows := 2000
	mat_a_cols := 2000
	mat_a := createMatrix(mat_a_rows, mat_a_cols, true)
	fmt.Printf("matrix a dimensions: %d x %d\n", mat_a_rows, mat_a_cols)
	//fmt.Printf("%v\n\n", mat_a)

	// create matrix B
	mat_b_rows := 2000
	mat_b_cols := 2000
	mat_b := createMatrix(mat_b_rows, mat_b_cols, true)
	fmt.Printf("matrix b dimensions: %d x %d\n", mat_b_rows, mat_b_cols)
	//fmt.Printf("%v\n\n", mat_b)

	//=========================================================================
	// single-threaded matrix multiplication
	//=========================================================================

	// start timer
	start := time.Now()

	// perform calculation
	mat_c_st, err := multiplySt(mat_a, mat_b)

	// stop timer
	duration := time.Since(start)

	checkError(err)

	fmt.Printf("single-threaded matrix multiplication time: %v\n", duration)
	//fmt.Printf("single-threaded matrix result:\n%v\n\n", mat_c_st)

	//=========================================================================
	// multi-threaded matrix multiplication
	//=========================================================================

	// start timer
	start = time.Now()

	// perform calculation
	mat_c_mt, err := multiplyMt(mat_a, mat_b)

	// stop timer
	duration = time.Since(start)

	checkError(err)

	fmt.Printf("multi-threaded matrix multiplication time: %v\n", duration)
	//fmt.Printf("multi-threaded matrix result:\n%v\n\n", mat_c_mt)

	// check if single and multi-threaded results were identical
	if !reflect.DeepEqual(mat_c_st, mat_c_mt) {
		fmt.Println("[ ERROR ] single and multi-threaded results are not identical")
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
