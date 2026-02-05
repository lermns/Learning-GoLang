package ejerciciotres

import (
	"fmt"
	"os"

	// "table" library import
	"github.com/aquasecurity/table"
)

// type struct definido para los productos
type Producto struct {
	Name  string
	Price float32
}

const IVA float32 = 0.21

func MostrarProductos() {
	var total float32

	// creo slice de productos
	productos := []Producto{
		{Name: "Caja", Price: 25.4},
		{Name: "Cama", Price: 120.5},
		{Name: "Silla", Price: 47.7},
		{Name: "Ropero", Price: 117.2},
		{Name: "Cocina", Price: 347.9},
		{Name: "Sofá", Price: 271.1},
		{Name: "TV", Price: 242.3},
	}

	// creo la tabla
	t := table.New(os.Stdout)

	// cabeceras.
	t.SetHeaders("ID", "Nombre", "Precio (€)", "Precio con IVA (€)")

	// relleno la tabla con los productos
	for i, producto := range productos {
		t.AddRow(fmt.Sprintf("%d", i+1), producto.Name, fmt.Sprintf("%.2f", producto.Price), calcularIVA(producto.Price))
		total += producto.Price
	}

	// fila total
	t.AddRow(fmt.Sprintf("%d", len(productos)+1), "TOTAL", fmt.Sprintf("%.2f", total), calcularIVA(total))
	// renderizo la tabla
	t.Render()

}

// función para calcular el precio con IVA
func calcularIVA(precio float32) string {
	return fmt.Sprintf("%.2f", (precio*IVA)+precio)
}
