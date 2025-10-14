package main

import "fmt"


type Persona struct {
	nombre string
	edad int
	altura float32
	ciudad string
}

func (p Persona) Saludar(mensaje string) string {
	return mensaje + ", " + p.nombre
}

func main() {
	var jose Persona = Persona{
		"Jose", 28, 1.60, "Tegucigalpa",
	}

	fmt.Println(jose)

	var maria Persona = Persona{
		ciudad: "San Pedro Sula",
		altura: 1.76,
		nombre: "Maria",
		edad: 27,
	}

	maria.ciudad = "La Ceiba"

	fmt.Println(maria)

	fmt.Println(maria.Saludar("Buenas tardes"))

}