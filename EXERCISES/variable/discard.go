package main

import ("fmt";"path")


func main(){

	var dir,file string

	dir,file =path.Split("css/main.go")

	fmt.Println("dir: ",dir ,"File: ",file)
}