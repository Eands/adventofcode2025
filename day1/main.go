package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	var count = 0
	var position = 50
	//Читаем файлик
	fileLines, err := readInput("input.txt")
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	//Парсим
	directionMap, err := parseDirections(fileLines)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	for _, value := range directionMap {
		position = int(math.Mod(float64(position-value), 100))
		if position == 0 {
			count++
		}
	}

	fmt.Println(count)
}

func absInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func readInput(fileName string) ([]string, error) {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения файла %s: %v", fileName, err)
	}
	// Сплитим строку по символу переноса
	return strings.Split(string(bytes), "\n"), nil
}

func parseDirections(input []string) ([]int, error) {
	result := make([]int, len(input))

	for index, s := range input {
		if s != "" {
			key := string(s[0])
			symbol := ""
			if key == "L" {
				symbol = "-"
			}
			numStr := symbol + (s[1:])

			value, err := strconv.Atoi(numStr)
			if err != nil {
				return nil, fmt.Errorf("ошибка парсинга %s: %v", s, err)
			}
			result[index] = value
		} else {
			fmt.Println("На индексе %d пустая строка", index)
		}
	}

	return result, nil
}
