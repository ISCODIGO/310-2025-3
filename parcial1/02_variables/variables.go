package main

import "fmt"

func main() {
	// declaracion de variable
	var entero int  // int: alias de int32
	entero = 10
	entero++

	// declaracion y asignacion
	var flotante float32 = 3.23

	// version reducida
	valor := 34

	const pi float32 = 3.14159

	//pi = 6.7 --> Error porque no se puede modificar una constante

	fmt.Println("Valor entero:", entero, valor)
	fmt.Println("Valor float32:", flotante)
	fmt.Println("Valor aprox. de PI:", pi)
}