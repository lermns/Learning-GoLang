package ejerciciosiete

//4. Escribe un programa que permita ir introduciendo una serie indeterminada de números mientras su suma no supere 50.
// Cuando esto ocurra, se debe mostrar el total acumulado y el contador de cuantos números se han introducido

import (
	"fmt"
)

const LIMIT = 50

func SumaIndeterminada() {
	var num, sum uint32

	for sum < LIMIT { //16-48-49-50(sale del for)
		fmt.Print("Introduce un número entero:\n→")
		fmt.Scanln(&num) //16-32-1-1
		sum += num       //16-48-49-50
		if sum > LIMIT {
			sum -= num
			fmt.Println("fuera de rango")
		}
	}

	fmt.Printf("La suma total es %d.\n", sum)
}
