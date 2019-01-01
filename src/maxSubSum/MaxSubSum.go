package maxSubSum

import (
	"math/rand"
)

/**
问题描述

　一个整数数组中的元素有正有负，在该数组中找出一 个连续子数组，
要求该连续子数组中各元素的和最大，这个连续子数组便被称作最大连续子数组。
比如数组{2,4,-7,5,2,-1,2,-4,3}的最大连续子数组为{5,2,-1,2}，
最大连续子数组的和为5+2-1+2=8。问题输入就是一个数组，输出该数组的“连续子数组的最大和”。
 */

/**
一、暴力解法

	 对数组内每一个数A[i]进行遍历，然后遍历以它们为起点的子数组，比较各个子数组的大小，
找到最大连续子数组
MaxSubSum1
MaxSubSum2

 */

type result struct {
	sum   int
	low   int
	hight int
}

func MaxSubSum1(a [] int) result {

	b := make([] result, 0)
	for i := 0; i < len(a); i++ {
		b = append(b, result{a[i], i, i})
	}
	var temp int
	for i := 0; i < len(a); i++ {
		temp = a[i]
		for j := i + 1; j < len(a); j++ {
			temp = a[j] + temp
			b = append(b, result{temp, i, j})
		}
	}
	return MaxValue(b)
}

func MaxSubSum2(a [] int) result {
	var r result
	max:=a[0]
	for i:=0;i<len(a);i++{
		for j:=i;j<len(a);j++{
			temp:=0
			for k:=i;k<=j;k++{
				temp=temp+a[k]
			}
			if temp>max{
				max=temp
				r=result{max,i,j}
			}
		}
	}
	return r

}

func MaxValue(a [] result) result {
	max := a[0]
	for i := 1; i < len(a); i++ {
		if a[i].sum > max.sum {
			max = a[i]
		}
	}
	return max
}

func Produce(n int) [] int {

	s := make([]int, 0)
	for i := 0; i < n; i++ {
		flag := 1
		if i%3 == 0 {
			flag = -1
		}
		s = append(s, flag*rand.Intn(n*100))
	}
	return s
}

func Produce1()[]int{
	 a := [10]int{1,2,3,4,-5,-6,5,-9,0,1}
	 return a[0:9]
}
