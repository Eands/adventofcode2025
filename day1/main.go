package main

import (
	"fmt"
	"math"
	"ru/adventofcode2025/utils"
	"strconv"
)

func main() {
	var part1Result = 0
	var part2Result = 0

	//Читаем файлик
	fileLines, err := utils.ReadInput("input.txt", "\n")
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

	//Part1
	part1Result = part1(directionMap, part1Result)
	//Part2
	part2Result = part2(fileLines)
	fmt.Printf("Part 1 result %d \n", part1Result)
	fmt.Printf("Part 2 result %d \n", part2Result)
}

func part1(directionMap []int, part1Result int) int {
	var position = 50
	for _, value := range directionMap {
		position = int(math.Mod(float64(position-value), 100))
		if position == 0 {
			part1Result++
		}
	}
	return part1Result
}

func part2(fileLines []string) int {
	var result int
	position := 50
	//создаем мапу с направлением вращения
	directionMap := map[byte]int{'R': 1, 'L': -1}

	for _, line := range fileLines {
		//получаем число
		rot, _ := strconv.Atoi(line[1:])

		for range rot {
			//получаем направление вращения 1 или -1 по первому символу из строки(R или L)
			if position += directionMap[line[0]]; position%100 == 0 {
				result++
			}
		}
	}
	return result
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
			fmt.Printf("На индексе %d пустая строка", index)
		}
	}

	return result, nil
}
