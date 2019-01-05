package mergeSort

import (
	"testing"
	"c"
	"sort/bubbleSort"
	"fmt"
)

func TestMergeArray(t *testing.T) {
	d:=MergeArray(bubbleSort.BubbleSort(c.ProduceArr(100)), bubbleSort.BubbleSort(c.ProduceArr(100)))
	fmt.Println(d)
	fmt.Println(c.Check(d))
}
