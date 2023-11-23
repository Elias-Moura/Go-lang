package main

import "fmt"



var nome = "Variável de escopo de pacote não podemos usar o := operador marmota"


func main() {
	fmt.Print(nome, " Temos acesso a ela de qualquer lugar do pacote após a sua declaração.")
	resultado := soma(10, 10)
	fmt.Print(resultado)
}

func soma(x int64, y int64) int64{
	return x + y
}