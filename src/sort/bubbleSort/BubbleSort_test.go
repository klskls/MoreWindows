package bubbleSort

import (
	"testing"
	"c"
)



func TestBubbleSort(t *testing.T) {
	n:=50000
	c.Check(BubbleSort(c.ProduceArr(n)))
	c.Show(BubbleSort(c.ProduceArr(n)))
}

func TestBubbleSort2(t *testing.T) {
	n:=50000
	c.Check(BubbleSort3(c.ProduceArr(n)))
	c.Show(BubbleSort3(c.ProduceArr(n)))
}

func TestBubbleSort3(t *testing.T) {
	n:=50000
	c.Check(BubbleSort2(c.ProduceArr(n)))
	c.Show(BubbleSort2(c.ProduceArr(n)))
}
