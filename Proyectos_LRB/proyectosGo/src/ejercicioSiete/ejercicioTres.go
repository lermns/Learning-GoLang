package ejerciciosiete

//3.     Escriba un programa que pida un número entero mayor que cero y calcule su factorial.
// La factorial es el resultado de multiplicar ese número por sus anteriores hasta la unidad.

import (
	"fmt"
)

func Factorial() {
	var numero, factorial uint16

	fmt.Print("Introduce un número entero:\n→")
	fmt.Scanln(&numero)
	factorial = numero

	for i := numero; i > 1; i-- {
		factorial = factorial * (i - 1)
	}
	fmt.Printf("El factorial de %d es %d\n", numero, factorial)
}
