package ejerciciosiete

//4. Escribe un programa que permita ir introduciendo una serie indeterminada de números mientras su suma no supere 50.
// Cuando esto ocurra, se debe mostrar el total acumulado y el contador de cuantos números se han introducido

import (
	"fmt"
)

func SumaIndeterminada() {
	var num, sum uint32

	for sum < 50 {
		fmt.Print("Introduce un número entero:\n→")
		fmt.Scanln(&num)
		sum += num
	}

	fmt.Printf("La suma total es %d.\n", sum)
}
