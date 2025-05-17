package internal

import "fmt"

func ArraysRun() {
	var arrays [10]int
	numberEven := [6]int{2, 4, 6, 8, 10, 12}

	fmt.Printf("the variable is of type %T\n", arrays)
	fmt.Println("value of an array", numberEven)

	// Slices

	var s []int = numberEven[1:4]

	fmt.Println("value of an slices", s)

}
