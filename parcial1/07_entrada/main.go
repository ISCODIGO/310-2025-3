package main

import "fmt"

func bisiesto(a int) bool {
	if a % 400 == 0 {
		return true
	} else if a % 4 == 0 && a % 100 != 0 {
		return true
	} else {
		return false
	}
}

func main() {
	// var numero int
	// var numero2 int
	// var operador string

	// fmt.Print("Escriba un numero: ")
	// fmt.Scan(&numero)

	// fmt.Print("Escriba un segundo numero: ")
	// fmt.Scan(&numero2)

	// fmt.Print("Escriba un operador: ")
	// fmt.Scan(&operador)

	// switch operador {
	// case "+":
	// 	fmt.Println(numero + numero2)
	// case "*":
	// 	fmt.Println(numero * numero2)
	// default:
	// 	fmt.Println("Operacion invalida")
	// }

	var anio int
	fmt.Scan(&anio)
	fmt.Println(bisiesto(anio))

}
