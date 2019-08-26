# jobs
Go jobs executes provided job on a slice in parallel with the limit of total number of go routines running at the same time.

## Installation
```bash
go get -u github.com/rafaeljesus/parallel-fn
```

## Usage
### Run
```go
package main

import (
  	"errors"

  	"github.com/goblock/jobs"
)

func main() {
    data := []int{1,2,3,4,5,6,7}

    // run with 3 go routines
    job.Parallel(data, 3, func(i int) {
        data[i]++
    })
    fmt.Println(data)
}
```


## Contributing
- Fork it
- Create your feature branch (`git checkout -b my-branch`)
- Commit your changes (`git commit -am 'Add some feature'`)
- Push to the branch (`git push origin my-branch`)
- Create new Pull Request

