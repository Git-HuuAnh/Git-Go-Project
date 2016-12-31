package main 

import (
	"fmt"
)

func swap (a, b int) (x int, y int) {

	temp := a
	a = b
	b = temp
	return 
}

func main() {
	fmt.Printf("Hello world 123 ! \n")
	swap(5, 9)
//	fmt.Printf("a = %d, b = %d \n", a , b)
}