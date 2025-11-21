package main

import (
	"fmt"
	"math"
)

// Encontrar si un numero es primo o no
// Primo es divisible solamente por 1 y por si mismo
func EsPrimoV1(numero int) bool {
	iteraciones := 0
	for i := 2; i < numero; i++ {
		iteraciones++;
		if numero % i == 0 {
			fmt.Println("Iteraciones:", iteraciones)
			return false
		}
	}

	fmt.Println("Iteraciones:", iteraciones)
	return true
}

func EsPrimoV2(numero int) bool {
	iteraciones := 0
	if numero % 2 == 0 {
		return true
	}

	for i := 3; i < numero / 3; i+=2 {
		iteraciones++
		if numero % i == 0 {
			fmt.Println("Iteraciones:", iteraciones)
			return false
		}
	}
	fmt.Println("Iteraciones:", iteraciones)
	return true
}

func EsPrimoV3(numero int) bool {
	iteraciones := 0
	limite := int(math.Sqrt(float64(numero)))
	for i := 2; i < limite; i++ {
		iteraciones++;
		if numero % i == 0 {
			fmt.Println("Iteraciones:", iteraciones)
			return false
		}
	}

	fmt.Println("Iteraciones:", iteraciones)
	return true
}


func main() {
	// fmt.Println("Version 1 O(n)...")
	// fmt.Println(EsPrimoV1(101))
	// fmt.Println(EsPrimoV1(683021))
	//fmt.Println(EsPrimoV1(9480049021))

	// fmt.Println("Version 2 O(n)...")
	// fmt.Println(EsPrimoV2(101))
	// fmt.Println(EsPrimoV2(683021))
	// fmt.Println(EsPrimoV2(596697466927))

	fmt.Println("Version 3 O(√n)...")
	fmt.Println(EsPrimoV3(101))
	fmt.Println(EsPrimoV3(683021))
	fmt.Println(EsPrimoV3(46451526632369))
}