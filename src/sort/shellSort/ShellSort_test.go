package shellSort

import (
	"testing"
	"c"
	"fmt"
)

func TestShellSort1(t *testing.T) {
	fmt.Print(ShellSort1(c.ProduceArr(1000000)))
	fmt.Print(c.Check(ShellSort1(c.ProduceArr(1000000))))
}

func TestShellSort2(t *testing.T) {
	fmt.Print(ShellSort2(c.ProduceArr(1000000)))
	fmt.Print(c.Check(ShellSort2(c.ProduceArr(1000000))))
}

func TestShellSort3(t *testing.T) {
	fmt.Print(ShellSort3(c.ProduceArr(100000)))
	fmt.Print(c.Check(ShellSort3(c.ProduceArr(100000))))
}