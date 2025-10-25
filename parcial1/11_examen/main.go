package main

import (
	"fmt"
)

func calcularDescuento(productos []float64) (float64, float64) {
	var total float64 = 0

	for _, precio := range productos {
		total += precio
	}

	if total > 5000 {
		descuento := total * 0.1
		return total - descuento, descuento
	}

	return total, 0
}

func main()  {
	// t1, d1 := calcularDescuento([]float64{6000, 4000, 8000, 2000})
	// fmt.Println(t1, d1)

	// t2, d2 := calcularDescuento([]float64{500, 60, 70})
	// fmt.Println(t2, d2)

	var calificaciones []int

	fmt.Println("Escriba calificaciones entre 0 y 100.\nSi coloca fuera de rango termina la lectura")
	for true {
		var entrada int
		fmt.Scan(&entrada)

		if entrada < 0 || entrada > 100 {
			break
		}

		calificaciones = append(calificaciones, entrada)
	}

	mayores_a_70 := 0
	conteo := len(calificaciones)

	for _, valor := range calificaciones {
		if valor >= 70 {
			mayores_a_70 ++
		}
	}

	fmt.Println("Estudiantes aprobados (>=70):", mayores_a_70)
	fmt.Println("Total de estudiantes:", conteo)


	for i := 0; i < 4; i++ {
		fmt.Println("Hola")
	}
}