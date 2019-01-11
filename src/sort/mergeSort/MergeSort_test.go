package mergeSort

import (
	"testing"
	"c"
	"sort/bubbleSort"
	"fmt"
)

func TestMergeArray(t *testing.T) {
	d := MergeArray(bubbleSort.BubbleSort(c.ProduceArr(100000)), bubbleSort.BubbleSort(c.ProduceArr(10000)), make([]int, 0))
	fmt.Println(d)
	fmt.Println(c.Check(d))
}

func TestMergeSort(t *testing.T) {
	s := c.ProduceArr(10)
	fmt.Println(s)
	MergeSort(s, 0, len(s)-1, make([]int, len(s)))
	fmt.Println(s)
	fmt.Println(c.Check(s))
}
