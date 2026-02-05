package ejercicioseis

import "fmt"

func NumPal() {
	var numero uint8

	fmt.Print("Dame un número entre 1 y 5\n→")
	fmt.Scanln(&numero)

	switch numero {
	case 1:
		fmt.Println("Uno")
	case 2:
		fmt.Println("Dos")
	case 3:
		fmt.Println("Tres")
	case 4:
		fmt.Println("Cuatro")
	case 5:
		fmt.Println("Cinco")
	default:
		fmt.Println("Número fuera de rango o no válido")
	}
}
