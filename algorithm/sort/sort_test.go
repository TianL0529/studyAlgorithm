package sort

import (
	"fmt"
	"testing"
)

func TestSort(T *testing.T) {
	numSlice := []int{1, 7, 9, 5, 4, 9, 2}
	quickSort(numSlice, 0, len(numSlice)-1)
	fmt.Printf("numSoreSlice: %+v", numSlice)
}
