package ejercicioonce

/*
Crear una aplicación  en go, que permita crear objetos de tipos [Persona]. Después se mostrará cada uno de los objetos creados.
Los atributos para cada persona, serán:
	Nombre.
	Apellidos.
	DNI.
	Edad.
	Estado civil.
	Año de nacimiento.
Por cada Persona creada, se deberán solicitar los datos pertinentes.

*/

import "fmt"

type Persona struct {
	Nombre        string
	Apellidos     string
	DNI           string
	Edad          int
	EstadoCivil   string
	AñoNacimiento int
}

func StructValue() {
	mapPersonas := map[string]Persona{}
	var cant uint8

	fmt.Println("Vamos a crear personas.")

	for {
		var p Persona

		fmt.Println("Ingrese el nombre:")
		fmt.Scanln(&p.Nombre)
		fmt.Println("Ingrese los apellidos:")
		fmt.Scanln(&p.Apellidos)
		fmt.Println("Ingrese el DNI:")
		fmt.Scanln(&p.DNI)
		fmt.Println("Ingrese la edad:")
		fmt.Scanln(&p.Edad)
		fmt.Println("Ingrese el estado civil:")
		fmt.Scanln(&p.EstadoCivil)
		fmt.Println("Ingrese el año de nacimiento:")
		fmt.Scanln(&p.AñoNacimiento)

		mapPersonas[p.DNI] = p
		fmt.Println("¡Agregado con éxito!")
		p.mostrarDatos()

		fmt.Println("Agregar más personas?\n1. si\n2. no")
		fmt.Scanln(&cant)

		if cant == 2 {
			break
		}
	}

	for _, p := range mapPersonas {
		p.mostrarDatos()
	}

}

func (p Persona) mostrarDatos() {
	fmt.Printf("Nombre: %s\n", p.Nombre)
	fmt.Printf("Apellidos: %s\n", p.Apellidos)
	fmt.Printf("DNI: %s\n", p.DNI)
	fmt.Printf("Edad: %d\n", p.Edad)
	fmt.Printf("Estado Civil: %s\n", p.EstadoCivil)
	fmt.Printf("Año de Nacimiento: %d\n", p.AñoNacimiento)
}
