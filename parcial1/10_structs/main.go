package main

import "fmt"

type Persona struct {
	nombre string
	edad int
	altura float32
	ciudad string
}

// Constructor
func New(nombre string, altura float32, edad int) (Persona, error) {
	if nombre == "" {
		return Persona{}, fmt.Errorf("Colocar un nombre valido")
	}

	if altura <= 0 {
		return Persona{}, fmt.Errorf("Altura: Colocar un flotante positivo")
	}

	if edad <= 0 {
		return Persona{}, fmt.Errorf("Edad: Colocar un entero positivo")
	}

	p := Persona {
		nombre: nombre,
		altura: altura,
		edad: edad,
		ciudad: "Tegucigalpa",
	}
	return p, nil
}

// Metodo donde el emisor se pasa por valor (copia)
func (p Persona) Saludar(mensaje string) string {
	return mensaje + ", " + p.nombre
}

// Metodo donde el emisor se pasa por referencia (el objeto)
func (p *Persona) CumplirEdad() {
	p.edad += 1
}



func main() {
	// Creacion de una variable de tipo struct Persona
	var jose Persona = Persona{
		"Jose", 28, 1.60, "Tegucigalpa",
	}

	fmt.Println(jose)

	// Inicializar una persona nombrando los atributos
	var maria Persona = Persona{
		altura: 1.76,
		nombre: "Maria",
		edad: 27,
	}

	//maria.ciudad = "La Ceiba"

	fmt.Println(maria)

	fmt.Println(maria.Saludar("Buenas tardes"))
	fmt.Println(jose.Saludar("Buenos dias"))

	maria.CumplirEdad()
	fmt.Println("Edad de Maria:", maria.edad)

	p, err := New("Alberto", 1.5, 14)

	if err != nil {
		fmt.Println("Hubo un error...", err)
	} else {
		fmt.Println("Se ha creado la persona:", p)
	}

}