package internal

import (
	"fmt"
)

type User struct {
	Name  string
	Age   uint8
	Phone uint32
	Dni   string
}

func (user User) ShowValues() {
	fmt.Println("Name:", user.Name)
	fmt.Println("Age:", user.Age)
	fmt.Println("Phone:", user.Phone)
	fmt.Println("Dni:", user.Dni)
}
