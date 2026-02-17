package ejerciciotrece

/*
Se pide desarrollar una aplicación que permita a los usuarios guardar y administrar contactos en un archivo.
Hay que recordar manejar los posibles errores que puedan ocurrir.
*/

import (
	"bufio"
	"ejerciciotrece/internal"
	"fmt"
	"log"
	"os"
	"strings"
)

func RegisterValues() {
	users := []internal.User{
		{Name: "Juan", Age: 25, Phone: 123456789, Dni: "12345678A"},
		{Name: "Ana", Age: 30, Phone: 987654321, Dni: "87654321B"},
		{Name: "Pedro", Age: 28, Phone: 111222333, Dni: "12345678A"},
	}

	filename := "../data/data.txt"

	for _, user := range users {
		err := appendFile(user, filename)
		if err != nil {
			log.Printf("NO -> Error: %v\n", err)
		} else {
			log.Printf("SI -> Usuario %s agregado\n", user.Name)
		}
	}

}

func isRepeat(dni string, filename string) (bool, error) {

	file, err := os.Open(filename)

	if os.IsNotExist(err) {
		return false, nil
	}

	if err != nil {
		return false, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")

		if len(parts) >= 4 && parts[3] == dni {
			return true, nil
		}
	}

	if err := scanner.Err(); err != nil {
		return false, err
	}

	return false, nil
}

func appendFile(user internal.User, filename string) error {
	duplicate, err := isRepeat(user.Dni, filename)

	if err != nil {
		return err
	}
	if duplicate {
		return fmt.Errorf("usuario con DNI %s ya existe", user.Dni)
	}

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer file.Close()

	line := fmt.Sprintf("%s,%d,%d,%s\n", user.Name, user.Age, user.Phone, user.Dni)
	_, err = file.WriteString(line)

	return err
}
