package maxSubSum

import (
	"testing"
	"fmt"
)

func TestMaxSubSum1(t *testing.T) {
    for i:=0;i<100;i++{
		n:=10
		s:=Produce(n)
		fmt.Println(s)
		fmt.Println(MaxSubSum1(s))


		fmt.Println(s)
		fmt.Println(MaxSubSum2(s))
	}


}

func TestMaxSubSum2(t *testing.T) {
	n:=10
	s:=Produce(n)
	fmt.Println(s)
	fmt.Println(MaxSubSum2(s))

	s=Produce1()
	fmt.Println(s)
	fmt.Println(MaxSubSum2(s))
}
