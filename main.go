package main

import (
	"fmt"
	"mymath"
	"mysort"
)

func main() {
	fmt.Printf("Hello, world.  Sqrt(2) = %v\n", mymath.Sqrt(2))
	fmt.Printf("Hello, world.  bubblesort = %d\n", mysort.BubbleSort(1))

	// http.HandleFunc("/", Redirect)
	// http.HandleFunc("/add", Add)
	// http.ListenAndServe(":8080", nil)
}
