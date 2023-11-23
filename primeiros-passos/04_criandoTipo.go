package main

import "fmt"

// Tipo base do hotdog é int - tipo subjacente 
// Entretanto, hotdog não pode ser atribuído a uma variável do tipo int
// Não há nenhum tipo de equivalência.
type hotdog int 
var b hotdog    

func main() {
	x := "Nome: "
	y := "Elias"
	fmt.Println(x + y)
	fmt.Printf("%T", b)
}

func soma(x int64, y int64) int64{
	return x + y
}