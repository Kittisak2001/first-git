package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("Hello, KTS2001")
	fmt.Println("Fix 2")
	fmt.Println(Add(5,5))
}

func Add(x,y int) int{
	return x+y
}