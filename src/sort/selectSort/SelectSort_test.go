package selectSort

import (
	"testing"
	"c"
	"fmt"
)

func TestSelectSort(t *testing.T) {
	fmt.Println(SelectSort(c.ProduceArr(10000)))
	fmt.Println(c.Check(SelectSort(c.ProduceArr(10000))))
}
