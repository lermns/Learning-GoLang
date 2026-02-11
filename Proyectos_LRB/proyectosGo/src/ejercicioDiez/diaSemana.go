package ejerciciodiez

/*
Crear un mapa que tenga como claves los números del día de la semana y como valores el nombre del día de la semana correspondiente:

1 → Lunes
2 → Martes
3 → Miércoles
4 → Jueves
5 → Viernes
6 → Sábado
7 → Domingo

Luego, crear un bucle que pregunte al usuario si quiere modificar algún día de la semana y cuántos. Si es así, le preguntará el número del día y el nuevo nombre del día. Finalmente, imprimirá el diccionario modificado.

*/

import (
	"fmt"
	"slices"
	"strconv"
)

func DiasSemana() {
	var cant uint8
	var day string
	var newValue string

	orderDays := []string{}

	mapWeek := map[uint8]string{
		1: "Lunes",
		2: "Martes",
		3: "Miercoles",
		4: "Jueves",
		5: "Viernes",
		6: "Sábado",
		7: "Domingo",
	}

	fmt.Println("Cuantos dias de la semana quieres cambiar?")
	fmt.Scanln(&cant)

	if cant < 1 || cant > 7 {
		fmt.Println("Valor no permitido")
		return
	}

	for i := 0; i < int(cant); {
		fmt.Println("Que día de la semana quieres cambiar?")
		fmt.Scanln(&day)

		// lo convertimos a entero para comprobar su valor
		num, err := strconv.Atoi(day)

		if err != nil || num < 1 || num > 7 {
			fmt.Println("Dia fuera de rango o no válido")
			continue
		}

		day := uint8(num)

		fmt.Printf("Vas a cambiar %s por\n->", mapWeek[day])
		fmt.Scanln(&newValue)

		mapWeek[day] = newValue
		i++
	}

	// pasamos la clave y el valor del mapa a un Slice
	for key := range mapWeek {
		fullValue := strconv.Itoa(int(key)) + " -> " + mapWeek[key]
		orderDays = append(orderDays, fullValue)
	}

	// lo ordenamos
	slices.Sort(orderDays)

	for _, v := range orderDays {
		fmt.Println(v)
	}
}
