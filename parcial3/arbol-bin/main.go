package main

import "fmt"

type Nodo struct {
	dato string
	izq *Nodo
	der *Nodo
}

func New(dato string, izquierda *Nodo, derecha *Nodo) (Nodo) {
	return Nodo {
		dato: dato,
		izq: izquierda,
		der: derecha,
	}
}

func Preorden(n Nodo) {
	fmt.Print(n.dato, "-")  // visitar la raiz del arbol o subarbol
	if n.izq != nil {
		Preorden(*n.izq)
	}

	if n.der != nil {
 		Preorden(*n.der)
	}
}

func Enorden(n Nodo) {
	if n.izq != nil {
		Enorden(*n.izq)
	}

	fmt.Print(n.dato, "-")  // visitar la raiz del arbol o subarbol

	if n.der != nil {
 		Enorden(*n.der)
	}
}

func PosOrden(n Nodo) {
	if n.izq != nil {
		PosOrden(*n.izq)
	}

	if n.der != nil {
		PosOrden(*n.der)
	}

	fmt.Print(n.dato, "-")
}


func main() {
	g := New("G", nil, nil)
	h := New("H", nil, nil)
	e := New("E", nil, nil)
	f := New("F", nil, nil)
	d := New("D", &g, &h)
	c := New("C", &f, nil)
	b := New("B", &d, &e)
	a := New("A", &b, &c)

	fmt.Print("\nPre-orden... ")
	Preorden(a)

	fmt.Print("\nEn-orden... ")
	Enorden(a)

	fmt.Print("\nPos-orden... ")
	PosOrden(a)

}