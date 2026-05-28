// EXERCISE: Swapper #2
//
//  1. Swap the values of `red` and `blue` variables
//
//  2. Print them
//
// EXPECTED OUTPUT
//  blue red


package main

import "fmt"

func main() {
	

	color, color2 := "red", "blue"

	var color3 string

	color3=color
	color=color2
	color2=color3

	fmt.Println(color,color2)
	
}