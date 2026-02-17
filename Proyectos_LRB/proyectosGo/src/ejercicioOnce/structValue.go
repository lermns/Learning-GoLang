package ejercicioonce

import (
	"fmt"

	// Proporciona lectura y escritura con buffer (almacenamiento temporal).
	"bufio"

	// Interactúa con el sistema operativo,Entrada estándar (teclado), Salida estándar (pantalla), Salida de errores
	"os"

	// Convierte entre strings y otros tipos (números, booleanos).
	"strconv"

	// Funciones para trabajar con cadenas de texto.
	"strings"
)

type Persona struct {
	Nombre         string
	Apellidos      string
	DNI            string
	Edad           int
	EstadoCivil    string
	AnioNacimiento int
}

func StructValue() {
	mapPersonas := map[string]Persona{}
	// creamos una instancia de un buffer y le pasamos os.Stdin para abrir el puntero de entrada
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Vamos a crear personas.")

	for {
		p := Persona{}

		// Leer nombre
		fmt.Println("Ingrese el nombre:")
		nombre, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error al leer nombre:", err)
			continue
		}
		p.Nombre = strings.TrimSpace(nombre)

		// Leer apellidos
		fmt.Println("Ingrese los apellidos:")
		apellidos, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error al leer apellidos:", err)
			continue
		}
		p.Apellidos = strings.TrimSpace(apellidos)

		// Leer DNI
		fmt.Println("Ingrese el DNI:")
		dni, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error al leer DNI:", err)
			continue
		}
		p.DNI = strings.TrimSpace(dni)

		// Leer edad
		fmt.Println("Ingrese la edad:")
		edadStr, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error al leer edad:", err)
			continue
		}
		edad, err := strconv.Atoi(strings.TrimSpace(edadStr))
		if err != nil {
			fmt.Println("Error: edad debe ser un número")
			continue
		}
		p.Edad = edad

		// Leer estado civil
		fmt.Println("Ingrese el estado civil:")
		estadoCivil, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error al leer estado civil:", err)
			continue
		}
		p.EstadoCivil = strings.TrimSpace(estadoCivil)

		// Leer año de nacimiento
		fmt.Println("Ingrese el año de nacimiento:")
		anioStr, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error al leer año:", err)
			continue
		}
		anio, err := strconv.Atoi(strings.TrimSpace(anioStr))
		if err != nil {
			fmt.Println("Error: año debe ser un número")
			continue
		}
		p.AnioNacimiento = anio

		// Agregar al mapa
		mapPersonas[p.DNI] = p
		fmt.Println("\n¡Agregado con éxito!")
		p.mostrarDatos()

		// Preguntar si continuar
		fmt.Println("--------------------------------------------------")
		fmt.Println("Agregar más personas?\n1. si\n2. no")
		opcionStr, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error al leer opción:", err)
			break
		}
		fmt.Println("--------------------------------------------------")

		opcion, err := strconv.Atoi(strings.TrimSpace(opcionStr))
		if err != nil || opcion == 2 {
			break
		}
	}

	// Mostrar todas las personas
	fmt.Println("\n=== Todas las personas ===")
	fmt.Println("--------------------------------------------------")
	for _, p := range mapPersonas {
		p.mostrarDatos()
		fmt.Println("--------------------------------------------------")
	}
}

func (p Persona) mostrarDatos() {
	fmt.Printf("Nombre: %s\n", p.Nombre)
	fmt.Printf("Apellidos: %s\n", p.Apellidos)
	fmt.Printf("DNI: %s\n", p.DNI)
	fmt.Printf("Edad: %d\n", p.Edad)
	fmt.Printf("Estado Civil: %s\n", p.EstadoCivil)
	fmt.Printf("Año de Nacimiento: %d\n", p.AnioNacimiento)
}
