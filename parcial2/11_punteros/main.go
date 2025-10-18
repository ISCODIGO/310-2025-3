package main

import "fmt"

func main() {
	x := 10
	// px es un puntero de x (se le asigna la direccion de memoria)
	px := &x

	// modificar x desde el puntero
	*px = 20

	fmt.Println("Valor de x:",x)
	fmt.Println("Valor del puntero x (direccion de mem.):", px)
	fmt.Println("Valor del puntero desreferencia:", *px)
}