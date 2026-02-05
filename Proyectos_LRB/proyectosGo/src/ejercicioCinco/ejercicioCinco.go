package ejerciciocinco

import "fmt"

func CalcularPares() {
	var numero int

	fmt.Println("Dame un número entero")
	fmt.Scanln(&numero)

	if numero == 0 {
		fmt.Printf("El número %d es cero\n", numero)
	} else if numero%2 == 0 {
		fmt.Printf("El número %d es par\n", numero)
	} else {
		fmt.Printf("El número %d es impar\n", numero)
	}
}
