package main

import "fmt"

// Go tem apenas a conversão de tipos

func main() {
	x := 10
	y := 90.5
	z := int(y)
	fmt.Printf("x: %v %T\n", x, x)
	fmt.Printf("y: %v %T\n", y, y)
	fmt.Printf("z: %v %T\n", z, z)	
}