// Indico que este archivo pertenece al paquete "constantes"
package constantes

/*
##################
#CONSTANTES EN GO#
##################

Las constantes almacenan valores que NO pueden cambiar durante la ejecución del programa. Se definen utilizando la palabra reservada [const] seguida del nombre de la constante, eltipo de dato (opcional) y el valor asignado.

Para trabajar con constantes, es recomendable seguir estas prácticas.

1. Nomrbes descriptivos:
	→ Utilizar nombres claros y descriptivos para las constantes, preferiblemente en mayúsculas y con guines bajos para separa palabras (MAX_VALUE).

2. Agrupación:
	→ Si tenemos varias constantes relacionadas, es buena práctica agurparalas utilizando paréntesis para mejorar su legibilidad.

3. Tipos datos:
	→ Aunque el tipo de dato es opcional, es una buena práctica especificarlo para evitar confusiones.

4. Uso adecuado:
	→ Utilizar constantes para valores que no deben cambiar, como configuraciones, límites o valores predeterminados.

5. Declaración global:
	→ Declarar las constantes a "nivel de paquete" para que sean accesibles desde cualquier función dentro del mismo paquete.

6. No uso:
	→ Al contrario que las variables, las constantes no está obligadas a ser usadas después de su declaración; si embargo, es una buena práctica utilizarlas para evitar advertencias del compilador.
*/

// Importación paquetes.
import (
	// Importo el paquete para tratar con la consola.
	"fmt"
	"os"

	// Importo el paquete de colores.
	"github.com/fatih/color"
	// Importo el paquete para crear tablas.
	"github.com/aquasecurity/table"
)

/*
####################################
#Zona de declaración de constantes.#
####################################
*/
// Declaro y defino la constante [PI]
const PI float32 = 3.14

// Como vemos, si intento modificar el valor de [PI], obtendremos un error de compilación.
// PI = 3.1416

// Igual que con las variables, puedo declarar múltiples constantes en un solo bloque y también pueden almacenar distintos tipos de datos.
const (
	// Constante con valor entero.
	X = 100
	// Constante con valor binario.
	Y = 0b1010
	// Constante con valor octal.
	Z = 0o12
	// Constante con valor hexadecimal.
	W = 0xff
)

/*
Cuando trabajamos con constantes, lo hacemos con valores que no van a variar, como la numeración de días de la semana.
*/
// Declaro una lista de días de la semana como constantes.
const (
	LUNES     = 1
	MARTES    = 2
	MIERCOLES = 3
	JUEVES    = 4
	VIERNES   = 5
	SABADO    = 6
	DOMINGO   = 7
)

// Defino una función pública llamada [ImprimirConstantes] que no recibe ningún parámetro ni devuelve ningún valor.
func ImprimirConstantes() {
	// Imprimo el valor de la constante [PI]
	fmt.Println("El valor de PI es: ", PI)
	// Imprimo las constantes definidas en bloque en una tabla.
	fmt.Println("+---------------------------------+")
	color.Red("+---------------------------------+")
	fmt.Println("| Constante | Valor | Tipo Dato   |")
	fmt.Println("+---------------------------------+")
	fmt.Println("|    X      | ", X, " | Decimal     |")
	fmt.Println("|    Y      | ", Y, "  | Binario     |")
	fmt.Println("|    Z      | ", Z, "  | Octal       |")
	fmt.Println("|    W      | ", W, " | Hexadecimal |")
	fmt.Println("+---------------------------------+")

	// Imprimo los días de la semana en una tabla.
	fmt.Println("+-------------------+")
	fmt.Println("| Días de la Semana |")
	fmt.Println("+-------------------+")
	fmt.Println("| Día       | Valor |")
	fmt.Println("+-------------------+")
	color.Red("| Lunes     |  ", LUNES, "  |")
	fmt.Println("| Martes    |  ", MARTES, "  |")
	fmt.Println("| Miércoles |  ", MIERCOLES, "  |")
	fmt.Println("| Jueves    |  ", JUEVES, "  |")
	fmt.Println("| Viernes   |  ", VIERNES, "  |")
	fmt.Println("| Sábado    |  ", SABADO, "  |")
	color.Blue("| Domingo   |  ", DOMINGO, "  |")
	fmt.Println("+-------------------+")

	// Instancio una tabla.
	t := table.New(os.Stdout)

	// Indico las cabeceras de la tabla.
	t.SetHeaders("Dia", "Valor")

	t.AddRow("Lunes", "1")
	t.AddRow("Martes", "2")
	t.AddRow("Miércoles", "3")
	t.AddRow("Jueves", "4")
	t.AddRow("Viernes", "5")
	t.AddRow("Sábado", "6")
	t.AddRow("Domingo", "7")

	t.Render()
}
