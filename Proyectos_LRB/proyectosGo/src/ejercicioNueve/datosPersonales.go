package ejercicionueve

// Crear un programa en go, que solicite distintos datos a un usuario desde la consola (nombre, apellido 1, apellido 2, dni y nº tfno).
// El programa deberá mostrar un menú solicitando si desea imprimir todos los datos o solo alguno en concreto.
// Se deberán utilizar las funciones correspondientes para ello.

import "fmt"

type Person struct {
	name   string
	apeOne string
	apeTwo string
	dni    string
	tef    uint32
}

func DatosPersonales() {
	var bandera bool
	p1 := Person{}

	(&p1).menuValues()

	for !bandera {
		switch menu() {
		case 1:
			fmt.Printf("\n\t %v \n", p1)
		case 2:
			fmt.Printf("\n\t %v \n", p1.name)
		case 3:
			fmt.Printf("\n\t %v \n", p1.apeOne)
		case 4:
			fmt.Printf("\n\t %v \n", p1.apeTwo)
		case 5:
			fmt.Printf("\n\t %v \n", p1.dni)
		case 6:
			fmt.Printf("\n\t %v \n", p1.tef)
		case 7:
			p1.menuValues()
		case 8:
			bandera = true
		default:
			fmt.Printf("\n\t Opción no válida \n")
		}
	}
}

// función con receptor menu valores de la persona pasado por referencia
func (p *Person) menuValues() {
	fmt.Printf("\tDame tu nombre ->")
	fmt.Scanln(&p.name)
	fmt.Printf("\n\tDame tu 1º apellido->")
	fmt.Scanln(&p.apeOne)
	fmt.Printf("\n\tDame tu 2º apellido->")
	fmt.Scanln(&p.apeTwo)
	fmt.Printf("\n\tDame tu DNI->")
	fmt.Scanln(&p.dni)
	fmt.Printf("\n\tDame tu número telefónico->")
	fmt.Scanln(&p.tef)
}

// función menu devuelve la opción seleccionada
func menu() uint8 {
	var opc uint8

	fmt.Printf(`
		OPCIONES
	1. Mostrar todos los datos.
	2. Mostrar Nombre.
	3. Mostrar Apellido 1.
	4. Mostrar Apellido 2.
	5. Mostrar DNI.
	6. Mostrar Telf.
	7. Ingresar nuevos valores.
	8. Salir.
	-> `)

	fmt.Scanln(&opc)

	return opc
}
