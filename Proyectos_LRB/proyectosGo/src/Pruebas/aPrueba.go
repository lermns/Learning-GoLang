package pruebas

import "fmt"

type Persona struct {
	Name     string
	Edad     int
	LastName string
}

func Pruebas() {

	fmt.Println("-----------------------------------------")

	p := Persona{}
	var p1 Persona
	fmt.Println(p)
	fmt.Println(p1)

	fmt.Println("-----------------------------------------")

	p2 := Persona{"leo", 25, "ramos"}
	p3 := Persona{Name: "leo", Edad: 25, LastName: "ramos"}
	fmt.Println(p2)
	fmt.Println(p3)

	fmt.Println("-----------------------------------------")

	p4 := struct {
		name string
		ape  string
		edad int
	}{
		name: "leo",
		ape:  "ramos",
		edad: 12,
	}

	fmt.Println(p4)

	fmt.Println("-----------------------------------------")

	fmt.Println(p4.name)
	fmt.Println(p4.ape)

	fmt.Println("-----------------------------------------")

	p4.name = "hernan"
	p4.edad = 34
	fmt.Println(p4.name)
	fmt.Println(p4.edad)

	fmt.Println("-----------------------------------------")

}
