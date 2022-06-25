package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type player struct {
	name, mark string
	wins       uint
}

var (
	player0   = player{}
	player1   = player{}
	playerNil = player{mark: " "}
)

func main() {
	p0Name, p0Mark := "Player 1", "o"
	p1Name, p1Mark := "Player 2", "x"

	playerToggle := 1

	flag.StringVar(&p0Name, "p1name", p0Name, "Name of player 1")
	flag.StringVar(&p1Name, "p2name", p1Name, "Name of player 2")
	flag.StringVar(&p0Mark, "p1mark", p0Mark, "Mark of player 1")
	flag.StringVar(&p1Mark, "p2mark", p1Mark, "Mark of player 2")
	flag.Parse()

	if p0Name != p1Name {
		player0.name = p0Name
		player1.name = p1Name
	}

	if p0Mark != p1Mark {
		player0.mark = p0Mark
		player1.mark = p1Mark
	}

	fmt.Print(drawWelcome())
	time.Sleep(time.Second)

	fmt.Print(drawGridNumbers())
	time.Sleep(time.Second * 1)

	for {
		// Some initialisation.
		choices := initChoices()
		turns := 0

		for {
			playerToggle = 1 - playerToggle
			currentPlayer := getCurrentPlayer(playerToggle)

			fmt.Print(drawGrid(choices))

			if turns == 9 {
				fmt.Println("Sorry, no winners.")

				input := getInput("Another game? (Y/n)")
				if input == "n" || input == "q" {
					doExit()
				}
				break
			}

			input := getInput(fmt.Sprintf("%s, choose a square from 1-9, or q to quit:", currentPlayer.name))
			if input == "q" {
				doExit()
			}

			cell, err := strconv.Atoi(input)
			if err != nil {
				continue
			}
			if cell < 1 || cell > 9 {
				continue
			}

			choices, err = setCell(choices, cell, currentPlayer)
			if err != nil {
				fmt.Println(err)
				continue
			}

			choices, win := checkWin(choices, currentPlayer)
			if win {
				fmt.Print(drawGrid(choices))
				doWin(playerToggle)

				fmt.Print(drawScoreboard(player0, player1))

				input := getInput("Another game? (Y/n)")
				if input == "n" || input == "q" {
					doExit()
				}

				break
			}

			turns++
		}
	}
}

func getInput(prompt string) string {
	fmt.Printf("\n%s ", prompt)
	reader := bufio.NewReaderSize(os.Stdin, 1)
	input, _ := reader.ReadByte()
	return strings.ToLower(string(input))
}

func getCurrentPlayer(cp int) player {
	if cp == 0 {
		return player0
	}

	return player1
}

func checkWin(choices [9]player, p player) ([9]player, bool) {
	// Green and 'bold'.
	newMark := fmt.Sprintf("\033[1;32m%s\033[0m", p.mark)

	switch {
	// Rows.
	case (choices[0] == p) && (choices[1] == p) && (choices[2] == p):
		choices[0].mark, choices[1].mark, choices[2].mark = newMark, newMark, newMark
		return choices, true
	case (choices[3] == p) && (choices[4] == p) && (choices[5] == p):
		choices[3].mark, choices[4].mark, choices[5].mark = newMark, newMark, newMark
		return choices, true
	case (choices[6] == p) && (choices[7] == p) && (choices[8] == p):
		choices[6].mark, choices[7].mark, choices[8].mark = newMark, newMark, newMark
		return choices, true

	// Cols.
	case (choices[0] == p) && (choices[3] == p) && (choices[6] == p):
		choices[0].mark, choices[3].mark, choices[6].mark = newMark, newMark, newMark
		return choices, true
	case (choices[1] == p) && (choices[4] == p) && (choices[7] == p):
		choices[1].mark, choices[4].mark, choices[7].mark = newMark, newMark, newMark
		return choices, true
	case (choices[2] == p) && (choices[5] == p) && (choices[8] == p):
		choices[2].mark, choices[5].mark, choices[8].mark = newMark, newMark, newMark
		return choices, true

	// Diagonals.
	case (choices[0] == p) && (choices[4] == p) && (choices[8] == p):
		choices[0].mark, choices[4].mark, choices[8].mark = newMark, newMark, newMark
		return choices, true
	case (choices[2] == p) && (choices[4] == p) && (choices[6] == p):
		choices[2].mark, choices[4].mark, choices[6].mark = newMark, newMark, newMark
		return choices, true
	}

	return choices, false
}

func doWin(playerToggle int) {
	var p player

	if playerToggle == 0 {
		player0.wins++
		p = player0
	} else {
		player1.wins++
		p = player1
	}

	fmt.Printf("Well done, %s, you have won!\n", p.name)
}

func setCell(choices [9]player, cell int, currentPlayer player) ([9]player, error) {
	if choices[cell-1] == (playerNil) {
		choices[cell-1] = currentPlayer
		return choices, nil
	}

	return choices, fmt.Errorf("cell %d already taken", cell)
}

func initChoices() [9]player {
	return [9]player{playerNil, playerNil, playerNil, playerNil, playerNil, playerNil, playerNil, playerNil, playerNil}
}

func doExit() {
	fmt.Println("\nBye! Thanks for playing!")
	os.Exit(0)
}
