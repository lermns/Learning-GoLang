package ejercicioocho

// Crea 1 array números y dos Slicen. El primero tendrá 5 números, el segundo se llamará pares y el tercero impares,
// ambos estarán vacíos. Después multiplica cada uno de los números del primer array por un número introducido por el usuario
// entre 1 y 10, si el resultado es par guarda ese número en el Slicen de pares y si es impar en el Slicen de impares.
// Muestra por consola la multiplicación que se produce junto con su resultado, con el formato:
// 2 x 3 = 6 -el Slicen de pares e impares.

import "fmt"

func Arrays() {
	var mult int
	x := [5]int{1, 2, 3, 4, 5}
	prueba(x)
	pares := []int{}
	impares := []int{}

	// recorremos el array
	for i := 0; i < len(x); {
		fmt.Println("dame un número entre 1-10: ")
		fmt.Scanln(&mult)

		// si esta fuera de rango volvemos iterar
		if mult < 1 || mult > 10 {
			fmt.Println("Número fuera de rango. Inténtalo de nuevo.")
			continue
		}

		result := x[i] * mult

		// si es par lo agrega e incrementa el valor i del for
		if result%2 == 0 {
			pares = append(pares, result)
			fmt.Printf("%d x %d = %d PAR\n", x[i], mult, result)
			i++
			continue
		}
		// caso impar lo mismo
		impares = append(impares, result)
		fmt.Printf("%d x %d = %d IMPAR\n", x[i], mult, result)
		i++

	}
	// termina a quita iteración y muestra ambos slices
	fmt.Println("Pares:", pares)
	fmt.Println("Impares:", impares)
}

func prueba(arry [5]int) {

}
