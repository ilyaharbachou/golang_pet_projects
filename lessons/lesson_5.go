package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("Hello \n", i)
	}

	for i := 0; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
			continue
		}

		fmt.Println(i)
		if i == 50 {
			break
		}
		fmt.Println(i)
	}

	//Like while
	f := 0

	for f < 5 {
		fmt.Println(f)
		f++
	}

	// for with array
	nums := []int{1, 2, 3, 4, 5}

	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}

	for index, element := range nums {
		fmt.Printf("Index %d Element %d \n", index, element)
	}

	for _, element := range nums {
		fmt.Printf("Element %d \n", element)
	}

	matrix := [][]int{{14, 25, 34}, {44, 54, 64}, {74, 84, 94}}

	for _, row := range matrix {
		for _, col := range row {
			fmt.Printf("%d ", col)
		}
		fmt.Println()
	}

}
