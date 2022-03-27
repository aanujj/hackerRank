package main

import "fmt"

func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	// Write your code here
	output := ""
	if v1 <= v2 {
		output := "NO"
		return output
	}

	// s := d/t
	// d = s * t
	//     (v1 * j)+x1
	//     (v2 * j)+x2

	if (x2-x1)%(v2-v1) == 0 {
		output = "YES"
		return output

	} else {
		output = "NO"
		return output
	}

}

func main() {
	strOutput := kangaroo(2, 1, 2, 2)
	fmt.Printf("strOutput: %v\n", strOutput)
}
