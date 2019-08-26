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
}
