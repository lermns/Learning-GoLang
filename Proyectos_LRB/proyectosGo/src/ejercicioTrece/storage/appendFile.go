package storage

import (
	"bufio"
	"ejerciciotrece/internal"
	"fmt"
	"os"
	"strings"
)

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
