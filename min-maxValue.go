package main

import "fmt"

func miniMaxSum(arr []int64) {
	// Write your code here
	max := arr[0]
	min := arr[0]

	maxVal := int64(0)
	minVal := int64(0)
	totalVal := int64(0)

	for _, v := range arr {

		if max < v {
			max = v
		}
		if min > v {
			min = v
		}
		totalVal = totalVal + v
	}
	fmt.Println(totalVal)

	minVal = totalVal - max
	maxVal = totalVal - min

	fmt.Print(minVal, maxVal)
	fmt.Println()

}

func main() {
	arr := []int64{793810624, 895642170, 685903712, 623789054, 468592370}
	miniMaxSum(arr)
}
