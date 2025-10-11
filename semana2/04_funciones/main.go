package main

import "fmt"

func sumar(a int, b int) int {
	return a + b
}

func sumar2(a int, b int) (t int) {
	t = a + b
	return
}

func dividir(a int, b int) (int, int) {
	cociente := a / b
	residuo := a % b
	return cociente, residuo
}

func dividir2(a int, b int) (float64) {
	cociente := float64(a) / float64(b)
	return cociente
}


func main() {
	x := 10
	y := 20
	fmt.Println("La suma de", x, "y", y, "es", sumar2(x, y))
	co, re := dividir(10, 5)
	fmt.Println("10/5 =","cociente:", co, "residuo:", re)
	fmt.Println(dividir2(1, 2))
}