package main

import "fmt"

func main() {
	x := []string{"a", "b", "c", "d"}
	y := x[:2] //a,b
	// cuando asignas un slice a otro, este tambien le pasa su capacidad lo cual muchas veces no es recomendable
	fmt.Println(cap(x), cap(y)) //4 4
	// tambien cuando hacemos el append, las copias se ven afectadas ya que comparten la referencia en memoria
	y = append(y, "z")   //a,b,z
	fmt.Println("x:", x) //x: a,b,z,d
	fmt.Println("y:", y) //y: a,b,z
}
