package arrays

import "fmt"

func PruebaArrays() {
	array := []string{"rojo", "azul", "verde", "amarillo", "dorado", "celeste", "cyan", "magenta"}
	fmt.Println(array)

	array2 := array[1:4]
	fmt.Println(array2)

	array = append(array[:4], array[5:]...)
	fmt.Println(array)
}
