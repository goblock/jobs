package main

import (
	"fmt"

	"github.com/goblock/jobs"
)

func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7}

	// run with 3 go routines
	jobs.Parallel(data, 3, func(i int) {
		data[i]++
	})
	fmt.Println(data)

	// using jobs.Run
	f1 := func() {
		for i := range data {
			data[i] *= 2
		}
	}
	data2 := []int{1, 2, 3, 4, 5, 6, 7}
	f2 := func() {
		for i := range data2 {
			data2[i] *= 2
		}
	}
	jobs.Run(f1, f2)
	fmt.Println(data, data2)
}
