package main

import "fmt"

func calcular(entrada [3]int) int {
	// array de 3 elementos:
	// el primer elemento si es positivo sumamos los siguientes
	// el primer elemento si es negativo restamos los siguientes
	if entrada[0] < 0 {
		return entrada[1] - entrada[2]
	} else {
		return entrada[1] + entrada[2]
	}
}

// funcion multivariadica (permite multiples argumentos)
func sumar(entradas ...int) int {
	total := 0
	for _, valor := range entradas {
		total += valor
	}
	return total
}

func main() {
	var arr [10]int
	arr[0] = 10
	arr[9] = 90

	// contenido
	fmt.Println(arr)

	// largo del array
	fmt.Println(len(arr))

	// recorrido
	for i, v := range arr {
		fmt.Println("Indice:", i, "Valor:", v)
	}

	var otroArr = [...]int {10, 20, 30}
	fmt.Println(otroArr)

	var matriz [2][3]int
	matriz[1][0] = 99
	matriz[0][1] = 88
	fmt.Println(matriz)

	matriz2 := [2][3]int {
		{1, 2},
		{4, 5, 6},
	}

	fmt.Println(matriz2)

	var arr2 = [5]int {10, 20, 30, 40, 50}
	fmt.Println(">>>")
	fmt.Printf("Tipo: %T\n", arr2)
	fmt.Printf("Tipo: %T\n", arr2[:3])
	fmt.Println(arr2[2:])
	fmt.Println(arr2[1:4])

	var inputs = [3]int {1, 2, 3}
	fmt.Println(calcular(inputs))
	fmt.Println(sumar(10, 20, 30, 40, 50, 60))
}