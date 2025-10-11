package main

import "fmt"
import "strconv"

func main() {
	f1 := 0.1
	f2 := 0.2

	i1 := 3456
	i2 := 5_000_000_000

	var letra rune = '{'
	var saludo string = "Hola a todos"

	fmt.Println(f1, f2, i1, i2)
	fmt.Println("Esto es un rune:", letra)
	fmt.Println(saludo)

	// sin signo (solo positivos)
	var byte_sin_signo uint8 = 255
	// con signo (permite negativos y positivos)
	var byte_con_signo int8 = 127

	fmt.Println(byte_sin_signo, byte_con_signo)
	fmt.Println(f1 + f2)

	// conversiones
	// entero -> flotante
	a := 10
	var b float32 = float32(a)

	// flotante -> entero
	c := 2.999999
	d := int(c)  // habra perdida de datos

	// entero -> string
	f := 10
	var e string = strconv.Itoa(f)
	fmt.Println(b, d, e)

	// string -> entero
	fmt.Println("Conversion de string a entero (parseo)")
	g := "4r5"
	h, err := strconv.Atoi(g)

	if err == nil {
		fmt.Println(h)
	} else {
		fmt.Println("No se pudo parsear:", err)
	}

	// flotante -> string
	pi := 3.14159
	var pi_string string = strconv.FormatFloat(pi, 'f', 4, 64)
	fmt.Println("Flotante -> string:", pi_string)

	// string -> flotante
	pi2 := "3.14d159"
	j, err2 := strconv.ParseFloat(pi2, 64)
	if err2 == nil {
		fmt.Println("Parseo de string -> float:", j)
	} else {
		fmt.Println("Fallo el parseo:", err2)
	}

	// convertir un string a []
	letras := []rune("Hola")
	fmt.Println(letras)
	fmt.Println(string(rune(letras[0])))

	cadena := string(letras)
	fmt.Println("Arreglo de bytes en string:", cadena)


}
