package main

import "fmt"

func hourglassSum(arr [][]int32) (max int32) {
	// Write your code here
	var sumArr []int32
	var sum int32
	for i := 0; i < len(arr)-2; i++ {
		for x := 0; x < len(arr)-2; x++ {
			//temp := arr[i][j] + arr[i][j+1] + arr[i][j+2] + arr[i+1][j+1] + arr[i+2][j] + arr[i+2][j+1] + arr[i+2][j+2]
			top := arr[i][x] + arr[i][x+1] + arr[i][x+2]
			middle := arr[i+1][x+1]
			bottom := arr[i+2][x] + arr[i+2][x+1] + arr[i+2][x+2]

			sum = top + middle + bottom

			sumArr = append(sumArr, sum)
		}
	}
	fmt.Printf("sumArr: %v\n", sumArr)
	max = sumArr[0]
	for i := 0; i < len(sumArr); i++ {

		if max < sumArr[i] {
			max = sumArr[i]
		}

	}
	fmt.Printf("max: %v\n", max)
	return max

}

func main() {
	a1 := []int32{-1, -1, 0, -9, -2, -2}
	a2 := []int32{-2, -1, -6, -8, -2, -5}
	a3 := []int32{-1, -1, -1, -2, -3, -4}
	a4 := []int32{-1, -9, -2, -4, -4, -5}
	a5 := []int32{-7, -3, -3, -2, -9, -9}
	a6 := []int32{-1, -3, -1, -2, -4, -5}

	var twoD = [][]int32{a1, a2, a3, a4, a5, a6}
	hourglassSum(twoD)

}
