package main

import (
	"fmt"
)

// data types, inputs

func main() {
	var age int8 = 23
	fmt.Println(age)

	var uAge uint = 2 // cannot be less than 0 'u'int
	fmt.Println(uAge)

	var number float64 = 275.672
	fmt.Println(number)

	age1 := 16
	fmt.Println(age1)

	var name string = "denis"
	name1 := "denis" //best Practice
	fmt.Println(name)
	fmt.Println(len(name1)) // string length len(String)

	//user inputs

	var nameI string
	var age2 uint8
	fmt.Println("What is your name?")
	fmt.Scan(&nameI) // symbol & - index of var
	fmt.Println("Hello " + nameI + "!")
	fmt.Println("How old are you?")
	fmt.Scan(&age2)
	fmt.Println("Your are age " + fmt.Sprint(age2) + " years!") // Number -> String fmt.Sprint(Numeric) supported all types of numebers

	var h int8 = 2
	var g int64 = int64(h) // int8 and int64 is a different data types
	fmt.Println(fmt.Sprint(g))
}
