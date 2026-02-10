package ejerciciosiete

//5.Escribe un programa con un bucle infinito con opciones para elegir, que pueda calcular el área de 2 figuras geométricas, triángulo y rectángulo.
// En primer lugar, pregunta de qué figura se quiere calcular el área, después solicita los datos que necesites para calcularlo.
// 1) triángulo = b * h/2
// 2) rectángulo = b * h
// 3) Salir: Solo cuando escojas esta opción el bucle se detendrá

import (
	"fmt"
)

var lado1, lado2 float64

func CalcularAreaFiguras() {
	var opcion int

	// Menú de opciones en bucle
	for {
		// Mostrar menú
		menu()
		fmt.Scanln(&opcion)

		// Procesar opción seleccionada con switch
		switch opcion {

		case 1:
			// Solicitar los lados al usuario
			guardarLados()
			// Calcular Rectángulo con método
			fmt.Println("\nEl área del rectángulo es: ", areaRect())
		case 2:
			// Solicitar los lados al usuario
			guardarLados()
			// Calcular Triangulo con método
			fmt.Println("\nEl área del triángulo es: ", areaTri())
		case 3:
			fmt.Println("Saliendo...")
			return
		default:
			fmt.Println("\nOpción no válida, intente de nuevo.")
		}
	}
}

// Función para mostrar el menú de opciones
func menu() {
	fmt.Printf("\n\tSeleccione una opción:\n1. Calcular Rectángulo.\n2. Calcular Triangulo.\n3. Salir.\n→")
}

// Función para calcular el área
func areaTri() float64 {
	area := (lado1 * lado2) / 2
	return area
}

func areaRect() float64 {
	area := lado1 * lado2
	return area
}

// Función para solicitar y guardar los lados del triángulo
func guardarLados() {
	fmt.Print("Ingrese lado 1: ")
	fmt.Scanln(&lado1)

	fmt.Print("Ingrese lado 2: ")
	fmt.Scanln(&lado2)
}
