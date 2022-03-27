package main

func main() {

	a := []int32{73, 67, 38, 33}
	gradingStudents(a)

}

func gradingStudents(grades []int32) []int32 {
	output := make([]int32, 0)

	for _, grade := range grades {

		roundValue := ((grade / 5) + 1) * 5

		roundDifference := roundValue - grade

		if grade < 38 {
			output = append(output, grade)
		} else if roundDifference < 3 {
			output = append(output, roundValue)
		} else {
			output = append(output, grade)
		}

	}

	return output
}
