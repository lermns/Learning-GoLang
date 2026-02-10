package ejerciciosiete

// 2.     Solicita un número e imprime todos los números pares e impares desde 1 hasta ese número con el  mensaje
// "es par" o "es impar" si el número es 5 el resultado será: 1 - es impar 2 - es par 3 – es impar 4 - es par 5 - es impar.

import (
	"fmt"
	"math"
)

func ParImpar() {
	var numero int8

	fmt.Print("Introduce un número entero:\n→")
	fmt.Scanln(&numero)
	numero = int8(math.Abs(float64(numero)))
	//fmt.Println(numero)

	// parseamos i a int8 para su comparación con numero
	for i := int8(1); i <= numero; i++ {
		if i%2 == 0 {
			fmt.Printf("%d - es par\n", i)
			continue
		}
		fmt.Printf("%d - es impar\n", i)
	}
}
