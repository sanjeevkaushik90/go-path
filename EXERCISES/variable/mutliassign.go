// EXERCISE: Multi Assign
//
//  1. Assign "go" to `lang` variable
//     and assign 2 to `version` variable
//     using a multiple assignment statement
//
//  2. Print the variables
//
// EXPECTED OUTPUT
//  go version 2
// ---------------






// package main


// import "fmt"

// func main() {
	
// 	var (
// 		lang    string
// 		version int
// 	)

// 	lang="go"
// 	version =2

	

	
// 	fmt.Println(lang, "version", version)
// }




// EXERCISE: Multi Assign #2
//
//  1. Assign the correct values to the variables
//     to match to the EXPECTED OUTPUT below
//
//  2. Print the variables
//
// HINT
//  Use multiple Println calls to print each sentence.
//
// EXPECTED OUTPUT
//  Air is good on Mars
//  It's true
//  It is 19.5 degrees



// package main

// import "fmt"

// func main() {
	

// 	var (
// 		planet string
// 		isTrue bool
// 		temp   float64
// 	)


// 	planet ="Air is good on Mars"
// 	isTrue= true
// 	temp=19.5

// 	fmt.Println(planet)
// 	fmt.Println(isTrue)
// 	fmt.Println(temp)

	
// }









// EXERCISE: Multi Short Func
//
// 	1. Declare two variables using multiple short declaration syntax
//
//  2. Initialize the variables using `multi` function below
//
//  3. Discard the 1st variable's value in the declaration
//
//  4. Print only the 2nd variable
//
// NOTE
//  You should use `multi` function
//  while initializing the variables
//
// EXPECTED OUTPUT
//  4



package main

import "fmt"


func main() {
	
	_,b:=multi()
	fmt.Println(b)
}


func multi() (int, int) {

	
	return 5,4
}