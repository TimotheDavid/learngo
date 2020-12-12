package main

import (
	"fmt"
	"math/rand"
)

func split(sum int) (x , y int) {
	x = sum * 3 / 10
	y = sum - x
	// return var nominal : sum 
	return 
}


// function who go two parameter of int and return a int 
func add(x, y int) int {
	return x + y
}

// main function mandatory
func main() {
	fmt.Println("Hello world", rand.Intn(30))
	fmt.Println(add(20,50))


}