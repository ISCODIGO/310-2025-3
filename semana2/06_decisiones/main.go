package main

import "fmt"

func esPar(num int) bool {
	return num%2 == 0
}

func calificacion(nota int, completoCurso bool) (string, error) {
	if nota < 0 || nota > 100 {
		return "", fmt.Errorf("Valores fuera de indice")
	} else {
		if nota == 0 {
			return "NSP", nil
		} else {
			if nota <= 64 {
				if completoCurso {
					return "RPB", nil
				} else {
					return "ABD", nil
				}
			} else {
				return "APB", nil
			}
		}
	}
}

func calificacion2(nota int, completoCurso bool) (string, error) {
	if nota < 0 || nota > 100 {
		return "", fmt.Errorf("Valores fuera de indice")
	} else if nota == 0 {
		return "NSP", nil
	} else if nota <= 64 && completoCurso {
		return "RPB", nil
	} else if nota <= 64 && !completoCurso {
		return "ABD", nil
	} else {
		return "APB", nil
	}
}

func diaSemana(posicion int) (string, error) {
	switch posicion {
	case 0:
		return "Domingo", nil
	case 1:
		return "Lunes", nil
	case 6:
		return "Sabado", nil
	default:
		return "", fmt.Errorf("El rango de dias es [0, 6]")
	}
}

func main() {
	x := 23

	if esPar(x) {
		fmt.Println(x, "es un numero par")
	} else {
		fmt.Println(x, "es impar")
	}

	etiqueta, err := calificacion(x, true)

	if err != nil {
		fmt.Println("Hubo un error...", err)
	} else {
		fmt.Println("Nota:", x, etiqueta)
	}

	dia, err := diaSemana(1)

	if err != nil {
		fmt.Println("Hubo un error...", err)
	} else {
		fmt.Println(dia)
	}

}
