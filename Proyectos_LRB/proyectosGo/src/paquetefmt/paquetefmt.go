package paquetefmt

import (
	"fmt"
)

func PaqueteFMT() {
	fmt.Println("Greetings and Salutations")
	var apellidoUno string
	var apellidoDos string
	// var x int = 10
	// var y float64 = 30.9
	// var sum1 float64 = float64(x) + y
	// var sum2 int = x + int(y)
	fmt.Println("Introduce tus dos apellidos: ")
	fmt.Scanln(&apellidoUno, &apellidoDos)
	fmt.Printf("Tu apellido es: %s %s\n", apellidoUno, apellidoDos)

}
