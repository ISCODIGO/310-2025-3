// funcion recursiva es una funcion que se llama a si misma.
//

package main

import "fmt"

func f(n int) {
	fmt.Println("Iniciando f...", n)
	if n < 5 {
		f(n + 1)  // llamado recursivo
	}
	fmt.Println("Finaliza f...", n)
}


// O(n)
func sumatoria(n int) int {
	if n == 1 {
		return 1
	}

	return sumatoria(n - 1) + n
}

var recursiones int
var recursiones_m int

// O (2^n)
func fibonacci(n int) int {
	// caso base
	recursiones++
	if n == 0 || n == 1 {
		return n
	}

	return fibonacci(n-1) + fibonacci(n-2)
}

// O(n)
func fibonacci_m(n int, m []int) int {
	recursiones_m++

	// crear una memoria
	if m == nil {
		m = make([]int, n + 1)
		m[0] = 0
		m[1] = 1
	}

	// caso base
	if n == 0 || n == 1 {
		return n
	}

	// caso base nuevo: buscar en la memoria
	if m[n] > 0 {
		return m[n]
	}

	m[n] = fibonacci_m(n-1, m) + fibonacci_m(n-2, m)
	return m[n]
}


func main() {
	f(0)
	fmt.Println("sumatoria: ", sumatoria(5))
	fmt.Println("Fibonacci mejorado: ", fibonacci_m(20, nil), "recursiones ", recursiones_m)
	fmt.Println("Fibonacci: ", fibonacci(20), "recursiones ", recursiones)
}