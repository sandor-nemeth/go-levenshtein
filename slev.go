package slev

import (
	"math"
)

type Row []int32

func SLev(w1 string, w2 string) int32 {
	cols := len(w1) + 1
	rows := len(w2) + 1

	var currentRow Row
	currentRow = Row(make([]int32, cols))
	currentRow[0] = 0
	for i := 1; i < cols; i++ {
		currentRow[i] = currentRow[i-1] + 1
	}

	for r := 1; r < rows; r++ {
		previousRow := currentRow
		currentRow = Row(make([]int32, cols))
		currentRow[0] = previousRow[0] + 1
		var rc int32
		for c := 1; c < cols; c++ {
			ic := currentRow[c-1] + 1
			dc := previousRow[c] + 1
			if w1[c-1] != w2[r-1] {
				rc = previousRow[c-1] + 1
			} else {
				rc = previousRow[c-1]
			}
			currentRow[c] = int32(math.Min(float64(ic), math.Min(float64(dc), float64(rc))))
		}
	}
	return currentRow[cols-1]
}
