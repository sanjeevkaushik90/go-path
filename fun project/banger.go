// cool!!!!
//hey!!!

//2. !!!!cool!!!!
//!!!hey!!!

//in upper case

// package main

// import (
// 	"fmt"
// 	"os"
// 	"strings"
// )

// func main() {
// 	str := os.Args[1]

// 	leng := len(str)

// 	banger:=str+strings.Repeat("!",leng)

// 	banger=strings.ToUpper(banger)

// 	fmt.Println(banger)

// }




package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	str := os.Args[1]

	leng := len(str)

	banger:=strings.Repeat("!",leng)+str+strings.Repeat("!",leng)

	banger=strings.ToUpper(banger)

	fmt.Println(banger)

}




