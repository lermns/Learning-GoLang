// Indico el paquete al que pertenece el archivo principal.
package main

// Importación de paquetes.
import (
	"fmt"
	// Importo el paquete que me permite imprimir el mensaje "hola mundo".
	"rsc.io/quote"
	// Importo un paquete externo, que me permite escribir texto con colores.
	"github.com/fatih/color"
)

// Punto de entrada del programa.
func main() {
	// Imprimo el mensaje "hola, Mundo!", utilizando el paquete importado [fmt]
	fmt.Println("hola, Mundo!")
	// Imprimo el mensaje de "hola mundo", mediante la función [Hello()], del paquete [quote]
	fmt.Println(quote.Hello())
	// Imprimo un mensaje famoso de Go, con la función [Go()], de paquete [quote].
	fmt.Println(quote.Go())
	// Imprimo un mensaje en color azul.
	color.Blue("Soy un texto en rojo")
	// Imprimo un mensaje en color azul.
	color.Red("Soy un texto en amarillo")
	color.Yellow("Soy un texto en verde")
}
