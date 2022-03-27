func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
	// Write your code here
	appleCount := 0
	orangesCount := 0
	for _, apple := range apples {

		location := a + apple

		fmt.Printf("lengh: %v\n", location)

		if location >= s && location <= t {
			appleCount = appleCount + 1

		}
	}
	for _, orange := range oranges {

		lengh := b + orange

		if lengh >= s && lengh <= t {
			orangesCount = orangesCount + 1

		}
	}

	fmt.Println(appleCount)
	fmt.Println(orangesCount)

}

func main() {

	var apple []int32 = []int32{-2}
	var orange []int32 = []int32{-1}

	countApplesAndOranges(2, 3, 1, 5, apple, orange)

}
