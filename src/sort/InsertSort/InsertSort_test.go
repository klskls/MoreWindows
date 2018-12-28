package InsertSort

import (
	"testing"
	"c"
)

func TestInsertSort(t *testing.T){
	n:=50000
	c.Check(InsertSort(c.ProduceArr(n)))
	c.Show(InsertSort(c.ProduceArr(n)))
}

func TestInsertSort2(t *testing.T){
	n:=50000
	c.Check(InsertSort2(c.ProduceArr(n)))
	c.Show(InsertSort2(c.ProduceArr(n)))
}

func TestInsertSort3(t *testing.T){
	n:=50000
	c.Check(InsertSort3(c.ProduceArr(n)))
	c.Show(InsertSort3(c.ProduceArr(n)))
}