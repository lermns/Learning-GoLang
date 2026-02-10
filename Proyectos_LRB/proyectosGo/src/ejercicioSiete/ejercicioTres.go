package ejerciciosiete

//3.     Escriba un programa que pida un número entero mayor que cero y calcule su factorial.
// La factorial es el resultado de multiplicar ese número por sus anteriores hasta la unidad.

import (
	"fmt"
)

func Factorial() {
	var numero, factorial uint64

	fmt.Print("Introduce un número entero:\n→")
	fmt.Scanln(&numero)
	factorial = numero

	if numero < 2 {
		fmt.Println("Valor no válido")
		return
	}

	for i := numero; i > 2; { //i = 5-4-3-2(sale del for)
		i-- //i=4-3-2
		fmt.Println(factorial, " x ", i)
		factorial = factorial * i // fac=20-60-120
	}
	fmt.Println(factorial)
	fmt.Printf("El factorial de %d es %d\n", numero, factorial)
}
