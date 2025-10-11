package main

import "fmt"

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
}