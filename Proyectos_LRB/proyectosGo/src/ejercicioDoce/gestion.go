package gestion

/*
Proyecto gestor de tareas: Crea un programa que permita gestionar una lista de tareas pendientes. Utiliza una estructura para representar
cada tarea (título, descripción, completado.) y un slice para almacenar todas las tareas. Utiliza métodos para agregar, eliminar y actualizar
tareas, y para marcar una tarea como completada. Este proyecto se tiene que hacer con entrada de datos por la consola.


v Crea una estructura para representar cada tarea, con campos para el título, la descripción y el estado de completado.

v Crea un slice para almacenar todas las tareas.

v Implementa un método para agregar una tarea al slice.

v Implementa un método para eliminar una tarea del slice.

v Implementa un método para actualizar una tarea en el slice.

v Implementa un método para marcar una tarea como completada.

v Implementa una función para mostrar la lista de tareas al usuario.

v Implementa una función para recibir la entrada del usuario por la consola,
que permita agregar, eliminar, actualizar y marcar tareas como completadas.

v Implementa un loop que solicite continuamente la entrada del usuario y muestre la lista de
tareas actualizada después de cada operación de gestión de tareas.
*/

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Tarea struct {
	Titulo      string
	Descripcion string
	Completado  bool
}

func Gestion() {
	reader := bufio.NewReader(os.Stdin)
	slcTareas := []Tarea{}
	slcTareasInt := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(slcTareasInt[:3-1])
	fmt.Println(slcTareasInt[3:])
	fmt.Println("Bienvenido al gestor de tareas.")

	for {
		opcion := menuMain(reader)

		switch opcion {
		case 1:
			agregarTarea(reader, &slcTareas)
		case 2:
			eliminarTarea(reader, &slcTareas)
		case 3:
			actualizarTarea(reader, slcTareas)
		case 4:
			mostrarTareas(slcTareas)
		case 5:
			completarTarea(reader, slcTareas)
		case 6:
			fmt.Println("Saliendo del gestor de tareas.")
			return
		default:
			fmt.Println("Opción no válida. Intente nuevamente.")
		}
	}

}

func menuMain(reader *bufio.Reader) int {
	fmt.Printf(`
		Seleccione una opción:
	1. Agregar tarea
	2. Eliminar tarea
	3. Actualizar tarea
	4. Mostrar tareas
	5. Completar tarea
	6. Salir 
	->`)

	opc, err := reader.ReadString('\n')
	num, err := strconv.Atoi(strings.TrimSpace(opc))

	if err != nil {
		fmt.Println("Error al convertir la entrada a número:", err)
		return 0
	}

	return num
}

func agregarTarea(reader *bufio.Reader, slcTareas *[]Tarea) {
	fmt.Print("Ingrese el título de la tarea: ")
	titulo, _ := reader.ReadString('\n')

	titulo = strings.TrimSpace(titulo)

	fmt.Print("Ingrese la descripción de la tarea: ")
	descripcion, _ := reader.ReadString('\n')

	descripcion = strings.TrimSpace(descripcion)

	nuevaTarea := Tarea{
		Titulo:      titulo,
		Descripcion: descripcion,
		Completado:  false,
	}

	*slcTareas = append(*slcTareas, nuevaTarea)
}

func eliminarTarea(reader *bufio.Reader, slcTareas *[]Tarea) {
	num := numTarea(*slcTareas, reader)

	*slcTareas = append((*slcTareas)[:num-1], (*slcTareas)[num:]...)
}

func mostrarTareas(slcTareas []Tarea) {
	if len(slcTareas) == 0 {
		fmt.Println("No hay tareas pendientes.")
		return
	}

	for i, tarea := range slcTareas {
		status := "Pendiente"
		if tarea.Completado {
			status = "Completada"
		}
		fmt.Printf("%d. %s - %s (%s)\n", i+1, tarea.Titulo, tarea.Descripcion, status)
	}
}

func actualizarTarea(reader *bufio.Reader, slcTareas []Tarea) {
	num := numTarea(slcTareas, reader)

	tarea := &slcTareas[num-1]

	fmt.Print("Ingrese el nuevo título de la tarea (deje en blanco para mantener el actual): ")
	titulo, _ := reader.ReadString('\n')

	titulo = strings.TrimSpace(titulo)
	if titulo != "" {
		tarea.Titulo = titulo
	}

	fmt.Print("Ingrese la nueva descripción de la tarea (deje en blanco para mantener la actual): ")
	descripcion, _ := reader.ReadString('\n')

	descripcion = strings.TrimSpace(descripcion)
	if descripcion != "" {
		tarea.Descripcion = descripcion
	}

	fmt.Println("Modificación completada")
}

func completarTarea(reader *bufio.Reader, slcTareas []Tarea) {
	num := numTarea(slcTareas, reader)
	tarea := &slcTareas[num-1]

	fmt.Print("¿La tarea está completada? (s/n): ")
	completadoStr, _ := reader.ReadString('\n')

	completadoStr = strings.TrimSpace(completadoStr)

	switch completadoStr {
	case "s":
		tarea.Completado = true
	case "n":
		tarea.Completado = false
	default:
		fmt.Println("Entrada no válida para el estado de completado. Se mantendrá el estado actual.")
	}

	fmt.Println("Modificación completada")
}

func numTarea(tarea []Tarea, reader *bufio.Reader) int {
	mostrarTareas(tarea)

	fmt.Print("Ingrese el número de la tarea a actualizar: ")
	numStr, _ := reader.ReadString('\n')

	num, err := strconv.Atoi(strings.TrimSpace(numStr))

	if err != nil || num < 1 || num > len(tarea) {
		fmt.Println("Número de tarea no válido.")
		return 0
	}

	return num
}
