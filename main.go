package main

import (
	"bufio"
	"log"
	"os"
	"time"
	"fmt"
)

const (
	BOARD_WIDTH  = 80
	BOARD_HEIGHT = 24
)

type Board struct {
	array [BOARD_HEIGHT][BOARD_WIDTH]int
}

var (
	board = Board{}
)

func (b *Board) SetXY(row, col, value int) {
	b.array[row][col] = value
}

func (b *Board) GetXY(row, col int) int {
	return b.array[row][col]
}

func initBoard() {
	file, err := os.Open("seed.txt")
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	row := 0

	for scanner.Scan() {
		line := scanner.Text()

		for col, value := range line {
			board.SetXY(row, col, int(value)-48)
		}

		row = row + 1
	}
}

func (b Board) getLiveNeighborCount(x, y int) int {
	count := 0

	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if i != 0 || j != 0 {
				value := b.GetXY((i+x+BOARD_HEIGHT)%BOARD_HEIGHT, (j+y+BOARD_WIDTH)%BOARD_WIDTH)

				if value == 1 {
					count++
				}
			}
		}
	}

	return count
}

func tick() {
	newBoard := Board{}

	for x, row := range board.array {
		for y, value := range row {
			count := board.getLiveNeighborCount(x, y)

			if value == 1 {
				if count < 2 {
					value = 0
				}
				//if count == 2 || count == 3 {
				//	value = 1
				//}
				if count > 3 {
					value = 0
				}
			}
			if value == 0 && count == 3 {
				value = 1
			}

			newBoard.SetXY(x, y, value)
		}
	}

	board = newBoard
}

func main() {
	initBoard()

	iteration := 0

	for {
		iteration++
		tick()

		time.Sleep(time.Duration(time.Millisecond * 10))

		fmt.Printf("\033[2J\033[0;0H")
		fmt.Println("Iteration ", iteration)
		for _, row := range board.array {
			for _, c := range row {
				if c == 1 {
					fmt.Print("X")
				} else {
					fmt.Print(" ")
				}
			}
			fmt.Println()
		}
	}

}
