package paquete2

type Persona struct {
	Nombre string
	Edad int
}

func New(nombre string, edad int) Persona {
	return Persona{
		Nombre: nombre,
		Edad: edad,
	}
}