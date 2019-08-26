package jobutil

import (
	"reflect"
	"sync"
)

// Parallel executes fnc on each element of the provided slice with go routine
// max defines the max number of workers to launch, when it's set to 0, the len of the slice is used
func Parallel(slice interface{}, max int, fnc func(i int)) {
	var w sync.WaitGroup
	jobs := reflect.ValueOf(slice)
	l := jobs.Len()

	if max <= 0 {
		max = l
	}

	for i := 0; i < l; i = i + max {
		x := i + max
		if x >= l {
			x = l
		}
		w.Add(x - i)
		for j := i; j < x; j++ {
			go func(k int) {
				defer w.Done()
				fnc(k)
			}(j)
		}
		w.Wait()
	}
}
