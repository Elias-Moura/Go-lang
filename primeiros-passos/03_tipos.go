package main

import "fmt"

// Go é uma linguagem de tipagem estática.
// Os tipos primitivos são string int e bool
// Os tipos de dados compostos são slice, array, struct e map
// Valor zero de alguns tipos:

var a int // 0
var b float64 // 0.0
var c string // ""
var d bool // false
// pointers, functions, interfaces, slices, channels, maps: nil

func main(){
		fmt.Printf("%d, %T\n", a, a)
		fmt.Printf("%f, %T\n", b, b)
		fmt.Printf("%s, %T\n", c, c)
		fmt.Printf("%v, %T\n", d, d)
}