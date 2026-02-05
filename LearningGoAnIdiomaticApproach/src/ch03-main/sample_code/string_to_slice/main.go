package main

import "fmt"

func main() {
	var s string = "Hello, 🌞"
	// Convertir la cadena a un slice de bytes y a un slice de runas
	var bs []byte = []byte(s)
	var rs []rune = []rune(s)
	fmt.Println(bs)
	fmt.Println(string(bs))
	fmt.Println(rs)
	fmt.Println(string(rs))
}
