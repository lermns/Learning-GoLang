package main

import "fmt"

func main() {
	x := []string{"a", "b", "c", "d"}
	y := x[:2] //a,b
	z := x[1:] //b,c,d
	x[1] = "y" //a,y,c,d
	y[0] = "x" //x,y - x=x,y,c,d
	z[1] = "z" //y,z,d - x=x,y,z,d
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
}
