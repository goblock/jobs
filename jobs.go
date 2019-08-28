package jobs

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
		if x > l {
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

// Run executes each function provided from the arguments concurrently and return when all done
func Run(funcs ...func()) {
	var w sync.WaitGroup
	w.Add(len(funcs))
	for i := range funcs {
		go func(k int) {
			defer w.Done()
			funcs[k]()
		}(i)
	}
	w.Wait()
}

// Do excutes each function in the provided slice concurrently and return when all done
// max defines the max number of workers to launch, when it's set to 0, the len of the slice is used
func Do(funcs []func(), max int) {
	var w sync.WaitGroup
	l := len(funcs)

	if max <= 0 {
		max = l
	}

	for i := 0; i < l; i = i + max {
		x := i + max
		if x > l {
			x = l
		}
		w.Add(x - i)
		for j := i; j < x; j++ {
			go func(k int) {
				defer w.Done()
				funcs[k]()
			}(j)
		}
		w.Wait()
	}
}
