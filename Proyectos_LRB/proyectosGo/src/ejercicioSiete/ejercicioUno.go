package ejerciciosiete

// 1.     Crea un sistema que pueda adivinar un carácter insertado es una vocal o no. Para eso tiene que ser insertado un carácter por consola.

import (
	"fmt"
	"strings"
)

func EsVocal() {
	var caracter byte

	fmt.Print("Introduce un carácter:\n→")
	fmt.Scanf("%c", &caracter)
	// fmt.Println(caracter)

	// salto de linea o espacio
	if caracter == 13 || caracter == 32 {
		fmt.Println("No se ha introducido ningún carácter")
		return
	}

	switch strings.ToLower(string(caracter)) {
	case "a", "e", "i", "o", "u":
		fmt.Printf("El carácter %c es una vocal\n", caracter)
	default:
		fmt.Printf("El carácter %c no es una vocal\n", caracter)
	}
}
