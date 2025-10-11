package main

import (
	"fmt"
	"slices"
)

// Funcion para eliminar un elemento del slice por posicion
func RemoverPosicion(slice []int, pos int) []int {
	s1 := slice[:pos]
	s2 := slice[pos+1:]
	return append(s1, s2...)
}

func RemoverValor(slice []int, valor int) []int {
	pos := -1

	// busqueda lineal
	for p, v := range slice {
		if valor == v {
			pos = p
			break
		}
	}

	if pos >= 0 {
		return RemoverPosicion(slice, pos)
	}
	return slice
}

func RemoverValorTodos(slice []int, valor int) []int {
	var nuevo []int

	for _, v := range slice {
		if v != valor {
			nuevo = append(nuevo, v)
		}
	}
	return nuevo
}

func main() {
	var m int
	var v int
	fmt.Println("Leer mes y el valor")
	fmt.Scan(&m, &v)
	fmt.Println("Mes:", m, "Facturacion:", v)


	var s []int
	s = append(s, 10, 20, 30)
	s = append(s, 100)
	fmt.Println(s)

	s2 := make([]string, 5) // slice de tamaño 5
	s2[0] = "hola"
	s2[1] = "10"
	s2[2] = "mitad"
	s2[3] = "2"
	s2 = append(s2, "sabado")
	fmt.Println(s2, "Tamaño:", len(s2))
	s3 := s2[:3]
	fmt.Println(s3)

	arreglo := [7]int{100, 200, 300, 400, 500, 300, 600}
	s4 := arreglo[:] // convertir un array completo en slice
	fmt.Printf("Tipado: %T, Valor: %v\n", s4, s4)

	/*s4 = append(s4, 100)
	s4 = append(s4, 100)
	s4 = append(s4, 100)
	s4 = append(s4, 100)
	s4 = append(s4, 100)
	s4 = append(s4, 100)*/

	fmt.Print("La estructura se creo con make...")
	fmt.Println("Capacity:", cap(s2), " Tamaño:", len(s2))

	fmt.Print("La estructura se creo con slicing...")
	fmt.Println("Capacity:", cap(s4), " Tamaño:", len(s4))

	s4 = RemoverPosicion(s4, 0)
	fmt.Println(s4)

	s4 = RemoverValorTodos(s4, 300)
	fmt.Println(s4)
	
	s5 := []int {10, 20, 30, 40, 50}
	s5 = slices.Delete(s5, 1, 3)  // Eliminar 20 y 30
	fmt.Println(s5)
	
	x := 40
	posicion := slices.Index(s5, x)
	esta := slices.Contains(s5, x)
	fmt.Printf("El valor %v esta en el slice? %v\n", x, esta)
	fmt.Printf("La posicion de %v es %v\n", x, posicion)
	
	fmt.Println("Antes del ordenamiento...")
	fmt.Println(s2)
	fmt.Println("Despues del ordenamiento...")
	slices.Sort(s2)
	fmt.Println(s2)
	
	s2 = slices.Insert(s2, 4, "xyz")
	fmt.Print(">>>")
	fmt.Println(s2)
}
