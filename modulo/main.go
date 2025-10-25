package main


import (
	"fmt"
	p1 "modulo/paquete1"
	"modulo/paquete2"
)

func main() {
	fmt.Println("Hola")
	suma := p1.Sumar(1, 2)
	fmt.Println(suma)

	fmt.Println("Igual:", p1.EsIgual(1, 2))

	// falla porque comienza con minuscula (no exportada)
	// p1.esDistinto(1, 2)

	persona := paquete2.Persona {
		Nombre: "Maria",
		Edad: 20,
	}

	fmt.Println(persona)

	p1 := paquete2.New("Jose", 20)
	fmt.Println(p1)
}