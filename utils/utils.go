package utils

import (
	"fmt"
	"os"
	"strings"
)

func ReadInput(fileName string, separator string) ([]string, error) {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения файла %s: %v", fileName, err)
	}
	// Сплитим строку по символу переноса
	return strings.Split(string(bytes), separator), nil
}
