package main

import (
	"fmt"
	"log"
	"ru/adventofcode2025/utils"
	"strconv"
	"strings"
)

func main() {
	result := 0
	lines, err := utils.ReadInput("input.txt", ",")
	if err != nil {
		log.Fatal(err)
	}

	for _, value := range lines {
		ranges := strings.Split(value, "-")

		fromInt, _ := strconv.Atoi(ranges[0])
		toInt, _ := strconv.Atoi(ranges[1])

		for fromInt <= toInt {
			fromStr := strconv.Itoa(fromInt)
			if isInvalidIds(fromStr) {
				fmt.Printf("Invalid ids: %s \n", fromStr)
				result = result + fromInt
			}

			fromInt++
		}
	}

	fmt.Printf("\n Result %d", result)
}

func isInvalidIds(s string) bool {
	var tmpVar = ""
	for i := 0; i < len(s); i++ {
		tmpVar += string(s[i])
		//Число вхождений в строку
		countEnth := strings.Count(s, tmpVar)
		//Рассчитываем максимульную длинну всех вхождений
		allLength := len(tmpVar) * countEnth
		//Сравниваем, если длинна равна всем вхождениями то true
		if len(s) == allLength {
			return true
		}
	}
	return false
}

//part1 solved
// func isInvalidIds(s string) bool {
// 	firstSubstring := s[:len(s)/2]
// 	lastSubstring := s[len(s)/2:]

// 	return firstSubstring == lastSubstring
// }
