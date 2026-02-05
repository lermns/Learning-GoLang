package ejerciciocuatro

import (
	"fmt"
	"math"
)

const DECIMAL = 2

var lado1, lado2, lado3 float64

func CalcularTriangulo() {
	var opcion int

	// Solicitar los lados al usuario
	guardarLados()

	// Menú de opciones en bucle
	for i := 0; i < 1; {
		// Mostrar menú
		menu()
		fmt.Scanln(&opcion)

		// Incrementar i solo si la opción es salir
		if opcion == 6 {
			i++
		}

		// Procesar opción seleccionada con switch
		switch opcion {

		case 1:
			// Calcular hipotenusa con método
			lado3 = hipotenusa()
			fmt.Println("La hipotenusa es: ", lado3)

		case 2:
			// Calcular área con método
			fmt.Println("El área es: ", area())

		case 3:
			// Calcular perímetro con método
			fmt.Println("El perímetro es: ", perimetro())

		case 4:
			// Calcular área y perímetro con formato de x decimales
			fmt.Printf("El área es: %.*f \n", DECIMAL, area())
			fmt.Printf("El perímetro es: %.*f \n", DECIMAL, perimetro())

		case 5:
			// Ingresar nuevos valores
			guardarLados()

		case 6:
			fmt.Println("Saliendo...")
			return

		default:
			fmt.Println("Opción no válida, intente de nuevo.")
		}

	}
}

// Función para mostrar el menú de opciones
func menu() {
	fmt.Printf("\n\tSeleccione una opción:\n1. Calcular hipotenusa.\n2. Calcular área.\n3. Calcular perímetro.\n4. Área y perímetro con %d decimales.\n5. Ingresar nuevos valores.\n6. Salir.\n→ ", DECIMAL)
}

// Función para calcular la hipotenusa
func hipotenusa() float64 {
	hipotenusa := math.Pow(lado1, 2) + math.Pow(lado2, 2)
	hipotenusa = math.Sqrt(hipotenusa)
	return hipotenusa
}

// Función para calcular el área
func area() float64 {
	area := (lado1 * lado2) / 2
	return area
}

// Función para calcular el perímetro
func perimetro() float64 {
	lado3 = hipotenusa()
	perimetro := lado1 + lado2 + lado3
	return perimetro
}

// Función para solicitar y guardar los lados del triángulo
func guardarLados() {
	fmt.Print("Ingrese lado 1: ")
	fmt.Scanln(&lado1)

	fmt.Print("Ingrese lado 2: ")
	fmt.Scanln(&lado2)
}
