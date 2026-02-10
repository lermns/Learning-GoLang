package ejerciciosiete

import "fmt"

func Prueba() {
	for i := 5; i >= 1; i-- {
		for j := i; j <= 5; j++ {
			fmt.Print("*  ")
		}
		fmt.Println()
	}

	fmt.Println()

	for range 5 {
		for range 5 {
			fmt.Print("*  ")
		}
		fmt.Println()
	}
}
