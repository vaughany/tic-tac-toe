package main

import "fmt"

func drawGrid(choices [9]player) string {
	// Turn the data into an interface for use with Sprintf.
	data := [][]interface{}{
		{choices[0].mark, choices[1].mark, choices[2].mark},
		{choices[3].mark, choices[4].mark, choices[5].mark},
		{choices[6].mark, choices[7].mark, choices[8].mark},
	}

	// gridBorder = "+---+---+---+\n"    // Grid borders
	// gridRow    = "| %s | %s | %s |\n" // Grid cells
	// out := gridBorder + fmt.Sprintf(gridRow, data[0]...) + gridBorder + fmt.Sprintf(gridRow, data[1]...) + gridBorder + fmt.Sprintf(gridRow, data[2]...) + gridBorder

	gridBorder := "---+---+---\n" // Grid borders
	gridRow := " %s | %s | %s\n"  // Grid cells
	out := fmt.Sprintf(gridRow, data[0]...) + gridBorder + fmt.Sprintf(gridRow, data[1]...) + gridBorder + fmt.Sprintf(gridRow, data[2]...)

	return out
}

func drawWelcome() string {
	t := player{mark: "T"}
	i := player{mark: "i"}
	c := player{mark: "c"}
	a := player{mark: "a"}
	o := player{mark: "o"}
	e := player{mark: "e"}

	return fmt.Sprintln("\nWelcome to:\n" + drawGrid([9]player{t, i, c, t, a, c, t, o, e}))
}

func drawGridNumbers() string {
	return fmt.Sprintln("When choosing a cell, use the following numbers:\n" + drawGrid([9]player{
		{mark: "1"},
		{mark: "2"},
		{mark: "3"},
		{mark: "4"},
		{mark: "5"},
		{mark: "6"},
		{mark: "7"},
		{mark: "8"},
		{mark: "9"},
	}))
}

func drawScoreboard(p0, p1 player) string {
	// Slightly verbose way of getting singular of 'wins' for just one win.
	p0wins, p1wins := "wins", "wins"
	if p0.wins == 1 {
		p0wins = "win"
	}
	if p1.wins == 1 {
		p1wins = "win"
	}

	return fmt.Sprintf("\n%s: %d %s.\n%s: %d %s.\n", p0.name, p0.wins, p0wins, p1.name, p1.wins, p1wins)
}
