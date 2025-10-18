package main

import "fmt"

type Nodo struct {
	dato int
	enlace *Nodo  // un puntero de tipo Nodo
}

func New(dato int) (Nodo) {
	return Nodo {
		dato: dato,
		enlace: nil,
	}
}

func main() {
	n1 := New(10)
	n2 := New(20)
	n3 := New(30)

	fmt.Println("Antes del enlace")
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)

	n1.enlace = &n2
	n2.enlace = &n3
	
	fmt.Println("Despues del enlace")
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)

	// Modificar desde n1 a n2
	n1.enlace.dato = 21
	n1.enlace.enlace.dato = 31
	fmt.Println("Modificando desde el puntero")
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)

	var n4 Nodo = New(40)
	// n3.enlace = &n4
	n1.enlace.enlace.enlace = &n4
	n3.enlace.dato = 41
	fmt.Println("Se puede modificar desde n1 y n3 a n4?", n4)


	// Recorrido de todos los elementos a partir de n1
	fmt.Println("Recorrido...")
	var temp *Nodo = &n1
	for temp != nil {
		fmt.Println(temp.dato)
		temp = temp.enlace
	}
}