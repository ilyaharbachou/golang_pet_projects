package main

import (
	"fmt"
	"math"
)

func main() {
	num := 3

	if num > 0 {
		fmt.Println("Number grate then 0")
	} else if num < 0 {
		fmt.Println("Number less than 0")
	} else if num == 3 {
		fmt.Println("Number equal 3")
	}

	var a float64
	var b float64
	var c float64

	fmt.Println("Введите а: ")
	fmt.Scan(&a)
	fmt.Println("Введи b:")
	fmt.Scan(&b)
	fmt.Println("Введи c:")
	fmt.Scan(&c)

	d := (b * b) - 4*(a*c)

	if d > 0 {
		var x1 float64
		var x2 float64

		x1 = (-b + math.Sqrt(d)) / (2 * a)
		x2 = (-b - math.Sqrt(d)) / (2 * a)

		fmt.Println(
			"Ваше уравнение имеет 2 корня\nD= " +
				fmt.Sprint(d) + "\nx1 = " +
				fmt.Sprint(x1) + "\nx2 = " +
				fmt.Sprint(x2))
	} else if d == 0 {
		var x float64

		x = (-b) / (2 * a)

		fmt.Println(
			"Ваше уравнение имеет один корень\nd= " +
				fmt.Sprint(d) + "\nx = " + fmt.Sprint(x))
	} else {
		fmt.Println("Ваше уравнение не имеет корней\nd = " + fmt.Sprint(d))
	}

	fmt.Scan(&a)

	//  go build -ldflags "-s -w" - optimized build without debug sign
}
