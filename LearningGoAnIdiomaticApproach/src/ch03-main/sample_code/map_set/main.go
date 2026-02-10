package main

import "fmt"

func main() {
	// Crear un conjunto de enteros que son la clave y usando un mapa con valores booleanos
	intSet := map[int]bool{}
	// 	Agregar algunos valores de este slice al conjunto
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
	for _, v := range vals {
		// cuando la clave coincide con el valor, se establece en true, lo que indica que el valor está presente en el conjunto
		intSet[v] = true
	}
	fmt.Println(len(vals), len(intSet))
	fmt.Println(intSet[5])
	fmt.Println(intSet[500])

	if intSet[100] {
		fmt.Println("100 is in the set")
	}
}
