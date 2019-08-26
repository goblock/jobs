package jobs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJobutil(t *testing.T) {
	d1 := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	d2 := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	for j := 1; j < len(d1)*2; j++ {
		Parallel(d1, j, func(i int) {
			d1[i] = d1[i] + 10
		})
		Parallel(d2, j+1, func(i int) {
			d2[i] = d2[i] + 10
		})
		assert.ElementsMatch(t, d1, d2)
	}
}
