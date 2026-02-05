/*
Para poder ejecutar correctamente el proyecto, nos moveremos con el cmd a la carpeta donde se encuentra nuestro proyecto y ejecutaremos el comando [go mod init nombreCarpetaTrabajo].

Para ejecutar el programa, utilizaremos el comando [go run nombrePrograma.go]
*/
// Indico el paquete al que pertenece el programa. Si no se encuentra dentro de ninguna carpeta, se suele dejar por defecto, en este caso, el nombre main.
package main

// Importo el paquete ["fmt"] (format).
// Contiene las funciones necesarias para imprimir text y leer entradas por consola.
import (
	"fmt"
	// "variable importado desde otra carpeta y haciendo el replace en el go.mod"
	"variables"

	// Importo que paquete que me permite imprimir "frases célebres". En este caso, "hola, mundo". Esto lo hago utilizando la ruta que podemos ver en el paquete [go.mod].
	"rsc.io/quote"
	// Importo el paquete de Github que me permite utilizar colores en la consola.

	"github.com/fatih/color"
)

// La función [main()] es el punto de entrada obligatorio de la aplicación.
// No recibe parámetros ni devuelve valores.
func main() {
	// Imprimo la mítica frase "hola, Mundo!" en la consola.
	// Esto lo haremos con la función [Println()] del paquete [fmt]. Esta función nos permite imprimir mensajes en la consola.
	// fmt.Print("hola Mundo! 🐱‍👤")
	// Imprimo la frase de "hola mundo", mediante la función [Hello()], del paquete [quote].
	fmt.Println(quote.Hello())
	// Ahora, utilizo otra de sus funciones para imprimir un mensaje distinto, famoso del lenguaje de Go.
	fmt.Println(quote.Go())
	// Utilizo la librería [github.com/fatih/color], para imprimir un mensaje en un color.
	color.Red("Hola, soy Íñigo Montoya.")
	// Llamo a la función que imprime lo requerido en el ejercicio 1.
	//ejercicios.MostrarNombre()
	// Llamo a la función para imprimir los tipos de datos.
	variables.MostrarDatos()

}

// Si ejecutamos el comando [go build main.go], se creará un [main.exe] que nos permitirá ejecutar el programa en cualquier parte.
