package ejerciciosiete

// 2.     Solicita un número e imprime todos los números pares e impares desde 1 hasta ese número con el  mensaje
// "es par" o "es impar" si el número es 5 el resultado será: 1 - es impar 2 - es par 3 – es impar 4 - es par 5 - es impar.

import (
	"fmt"
)

func ParImpar() {
	var numero uint16

	fmt.Print("Introduce un número entero:\n→")
	fmt.Scanln(&numero)

	for i := uint16(1); i <= numero; i++ {
		if i%2 == 0 {
			fmt.Printf("%d - es par\n", i)
		} else {
			fmt.Printf("%d - es impar\n", i)
		}
	}
}
