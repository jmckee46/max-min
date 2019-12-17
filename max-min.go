package main

import (
	"fmt"
	"sort"
)

type Arr []int32

func (a Arr) Len() int           { return len(a) }
func (a Arr) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Arr) Less(i, j int) bool { return a[i] < a[j] }

func maxMin(k int32, arr []int32) int32 {
	arrLength := int32(len(arr))
	myArr := Arr(arr)
	var unfairness int32
	var temp int32

	sort.Sort(myArr)
	fmt.Println("")
	fmt.Println("k:", k)
	fmt.Println("myArr:", myArr)

	unfairness = myArr[k-1] - myArr[0]

	for i := int32(0); i <= arrLength-k; i++ {
		fmt.Println(myArr[i], myArr[i+k-1], "=", (myArr[i+k-1] - myArr[i]))
		temp = myArr[i+k-1] - myArr[i]
		if temp < unfairness {
			unfairness = temp
		}
	}

	return unfairness
}

func main() {}
