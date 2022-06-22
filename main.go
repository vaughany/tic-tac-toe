package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

const (
	e = " " // Empty, 'cos it's easier to represent an empty cell with a space than an empty string.
	o = "o" // Player 1.
	x = "x" // Player 2.
)

var (
	gridBorder = "+---+---+---+\n"    // Grid borders
	gridRow    = "| %s | %s | %s |\n" // Grid cells

	currentPlayer = 0
)

func main() {
	fmt.Print(drawWelcome())
	time.Sleep(time.Second)

	fmt.Print(drawGridNumbers())
	time.Sleep(time.Second)

	choices := [9]string{e, e, e, e, e, e, e, e, e}

	for {
		fmt.Print(drawGrid(choices))

		fmt.Printf("\nPlayer %d, choose a square from 1-9, or q to quit: ", currentPlayer+1)

		reader := bufio.NewReaderSize(os.Stdin, 1)
		// input should be 1-9 or q to quit.
		input, _ := reader.ReadByte()

		// Check for 'q'.
		if string(input) == "q" {
			os.Exit(0)
		}

		int, err := strconv.Atoi(string(input))
		if err != nil {
			continue
		}
		if int < 1 || int > 9 {
			continue
		}

		choices, err = setCell(choices, int, currentPlayer)
		if err != nil {
			fmt.Println(err)
		} else {
			choices, win := checkWin(choices, currentPlayer)
			if win {
				fmt.Print(drawGrid(choices))
				doWin(currentPlayer)
				os.Exit(0)
			}
			currentPlayer = 1 - currentPlayer
		}
	}
}

func checkWin(choices [9]string, player int) ([9]string, bool) {
	mark := o
	if player == 1 {
		mark = x
	}

	// Green and 'bold'.
	newMark := fmt.Sprintf("\033[1;32m%s\033[0m", mark)

	switch {
	// Rows.
	case (choices[0] == mark) && (choices[1] == mark) && (choices[2] == mark):
		choices[0], choices[1], choices[2] = newMark, newMark, newMark
		return choices, true
	case (choices[3] == mark) && (choices[4] == mark) && (choices[5] == mark):
		choices[3], choices[4], choices[5] = newMark, newMark, newMark
		return choices, true
	case (choices[6] == mark) && (choices[7] == mark) && (choices[8] == mark):
		choices[6], choices[7], choices[8] = newMark, newMark, newMark
		return choices, true

	// Cols.
	case (choices[0] == mark) && (choices[3] == mark) && (choices[6] == mark):
		choices[0], choices[3], choices[6] = newMark, newMark, newMark
		return choices, true
	case (choices[1] == mark) && (choices[4] == mark) && (choices[7] == mark):
		choices[1], choices[4], choices[7] = newMark, newMark, newMark
		return choices, true
	case (choices[2] == mark) && (choices[5] == mark) && (choices[8] == mark):
		choices[2], choices[5], choices[8] = newMark, newMark, newMark
		return choices, true

	// Diagonals.
	case (choices[0] == mark) && (choices[4] == mark) && (choices[8] == mark):
		choices[0], choices[4], choices[8] = newMark, newMark, newMark
		return choices, true
	case (choices[2] == mark) && (choices[4] == mark) && (choices[6] == mark):
		choices[2], choices[4], choices[6] = newMark, newMark, newMark
		return choices, true
	}

	return choices, false
}

func doWin(player int) {
	fmt.Printf("Well done, player %d, you have won!\n", player+1)
}

func setCell(choices [9]string, cell, player int) ([9]string, error) {
	if choices[cell-1] == e {
		if player == 0 {
			choices[cell-1] = o
		} else {
			choices[cell-1] = x
		}
		return choices, nil
	}

	return choices, fmt.Errorf("cell %d already taken", cell)
}

func drawGrid(choices [9]string) string {
	// Turn the data into an interface for use with Sprintf.
	data := [][]interface{}{
		{choices[0], choices[1], choices[2]},
		{choices[3], choices[4], choices[5]},
		{choices[6], choices[7], choices[8]},
	}

	return gridBorder + fmt.Sprintf(gridRow, data[0]...) + gridBorder + fmt.Sprintf(gridRow, data[1]...) + gridBorder + fmt.Sprintf(gridRow, data[2]...) + gridBorder
}

func drawWelcome() string {
	return fmt.Sprintln("\n Welcome to:\n" + drawGrid([9]string{"T", "i", "c", "T", "a", "c", "T", "o", "e"}))
}

func drawGridNumbers() string {
	return fmt.Sprintln("When choosing a cell, use the following numbers:\n" + drawGrid([9]string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}))
}
