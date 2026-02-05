package variables

import (
	"fmt"
)

// Importación de paquetes.

// Función en la que trataré con ditintos tipos de variables y datos e imprmiré sus valores.
func TratoVariables() {
	// Declaro una variable y le asigno un tipo de dato.
	var nombreUno string

	// Decalro varias variables DEL MISMO TIPO en una sola línea.
	var nombreDos, apellidoUno string

	// Declaro y defino una variable para almacenar la edad.
	var edad int = 29

	// Declaro 3 variables de tipos distintos.
	var (
		pais, ciudad string
		cPostal      int
	)

	// Declaro y defino 3 variables al mismo tiempo.
	var (
		dni         string = "15988765a"
		tfno        int    = 689452566
		estadoCivil string = "Soltero"
	)

	// Declaro y defio las variables [altura] y [unidadMediad], sin especificar el tipo de dato.
	var (
		altura       = 1.60
		unidadMedida = "metros"
	)

	// Defino la variable [nombreUno], asignándole un valor de tipo [string]
	nombreUno = "Víctor"

	// Defino las variables [nombreDos] y [apellidoUno]
	nombreDos = "Luis"
	apellidoUno = "Glaría"

	// Asigno valores a las variables [pais], [ciudad] y [cPostal]
	pais = "España"
	ciudad = "Madriz"
	cPostal = 28956

	// Modifico los valores directamente, en una sola línea.
	nombreUno, edad, dni, tfno, estadoCivil, altura = "Juan Ignacio", 56, "15684635491g", 689515678, "Es Complicado", 1.93

	// Imprimo el valor de las variables.
	fmt.Println(nombreUno)
	fmt.Println(nombreDos, apellidoUno)
	// Imprimo el valor de la variable [edad]
	fmt.Println(edad)

	// Imprimo los valores del lugar de residencia.
	fmt.Println(pais, ciudad, cPostal)

	// Imprimo los valores "sobre mi"
	fmt.Println("DNI: ", dni, "Nº Tfno: ", tfno, "Estado Civil: ", estadoCivil)

	// Imprimio el valor de las variables [altura] y [unidadMedida]
	println("Altura → ", altura, unidadMedida)

	/*
		Otra manera de declarar variables, es sin utilizar la palabra [var], utilizando el operador [:=]

		IMPORTANTE: Esta manera de declarar variables, solo puede utilizarse dentro de las funciones.
	*/
	// Defino varias variables en una sola línea, sin untilizar la palabra reservada [var]
	estudios, trabajo, experiencia := "Informática", "Desarrollador Go", 5

	// Imprimo todos los valores sobre mí.
	fmt.Println("\nDatos Personales:")
	fmt.Println("Nombre Completo: ", nombreUno)
	fmt.Println("Edad: ", edad)
	fmt.Println("DNI: ", dni)
	fmt.Println("Nº Tfno: ", tfno)
	fmt.Println("Estado Civil: ", estadoCivil)
	// Utilizo el formateo de texto, para imprimir una frase junto con las variables.
	fmt.Printf("Altura: %.2f %s\n", altura, unidadMedida)
	// Imprimo las variables, mostrando también el tipo de dato que almacena una de ellas.
	fmt.Printf("Estudios %s | Trabajo %s | Años Experiencia %d [La variable experiencia es de tipo → %T ]", estudios, trabajo, experiencia, altura)
}
