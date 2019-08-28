# jobs
Go jobs executes provided job on a slice in parallel with the limit of total number of go routines running at the same time.

## Installation
```bash
go get -u github.com/goblock/jobs
```

## Usage
### Run
```go
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
			data[i] *=2
		}
	}
	data2 := []int{1, 2, 3, 4, 5, 6, 7}
	f2 := func() {
		for i := range data2 {
			data2[i] *=2
		}
	}
	jobs.Run(f1, f2)
	fmt.Println(data, data2)
}
```


## Contributing
- Fork it
- Create your feature branch (`git checkout -b my-branch`)
- Commit your changes (`git commit -am 'Add some feature'`)
- Push to the branch (`git push origin my-branch`)
- Create new Pull Request

---
> GitHub [@goblock](https://github.com/goblock)