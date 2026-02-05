package main

import "fmt"

func main() {
	// slice con capacidad 5 y longitud 0 usando "make"
	x := make([]string, 0, 5)
	x = append(x, "a", "b", "c", "d")
	y := x[:2:2]  //
	z := x[2:4:4] //
	fmt.Println(cap(x), cap(y), cap(z))
	y = append(y, "i", "j", "k") //
	x = append(x, "x")
	z = append(z, "y")
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
}
