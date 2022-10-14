package main

import (
	"fmt"
	"strings"
)

func getInput(height, width int) (int, int) {
	var x, y int
	for {
		fmt.Println("Отправьте строку")
		_, err := fmt.Scan(&x)
		if err != nil {
			panic(err)
		}

		fmt.Println("Отправьте столбец")
		_, err = fmt.Scan(&y)
		if err != nil {
			panic(err)
		}

		if x > height || x < 0 || y > width || y < 0 {
			fmt.Println("Вы ввели что-то не то")
		} else {
			break
		}
	}
	return x, y
}

func printMap(playMap [3][3]int) {
	m := map[int]string{
		0: "-",
		1: "x",
		2: "o",
	}

	for i := 0; i < len(playMap); i++ {
		fmt.Print("\t ", i)
	}
	fmt.Println()

	for i, x := range playMap {
		fmt.Println(strings.Repeat("\t___ ", len(x)))
		fmt.Print(i)
		for _, y := range x {
			fmt.Print("\t|", m[y], "| ")
		}
		fmt.Println("\n\t" + strings.Repeat("--- \t", len(x)))
	}
}

func checkWin(playMap [3][3]int, checkInt int) bool {
	for n := range playMap {
		if playMap[n][0] == checkInt && playMap[n][1] == checkInt && playMap[n][2] == checkInt {
			return true
		}
	}

	/*
		1
		 1
		  1
	*/
	if playMap[0][0] == checkInt && playMap[1][1] == checkInt && playMap[2][2] == checkInt {
		return true
	}

	/*
		  1
		 1
		1
	*/
	if playMap[0][2] == checkInt && playMap[1][1] == checkInt && playMap[0][0] == checkInt {
		return true
	}

	return false
}

func main() {
	var x1, x2, y1, y2 int

	playMap := [3][3]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}
	height := len(playMap)
	width := len(playMap[0])
	printMap(playMap)

	for {
		fmt.Println("Первый игрк!")
		for {
			x1, y1 = getInput(height, width)

			if playMap[x1][y1] != 0 {
				fmt.Println("Так сходить нельзя")
			} else {
				playMap[x1][y1] = 1
				break
			}

		}

		printMap(playMap)
		isWin := checkWin(playMap, 1)
		if isWin {
			fmt.Println("Первый игрок победил!")
			break
		}

		fmt.Println("Второй игрок!")
		for {
			x2, y2 = getInput(height, width)

			if playMap[x2][y2] != 0 {
				fmt.Println("Так сходить нельзя")
			} else {
				playMap[x2][y2] = 2
				break
			}

		}
		printMap(playMap)
		isWin = checkWin(playMap, 2)
		if isWin {
			fmt.Println("Второй игрок победил!")
			break
		}
	}
}
