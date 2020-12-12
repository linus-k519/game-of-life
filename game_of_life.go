package main

import (
	"fmt"
	"strings"
)

type Game [][]Cell

type Cell bool

func (c Cell) String() string {
	if c {
		return "X"
	} else {
		return " "
	}
}

func (c Cell) Int() int {
	if c {
		return 1
	} else {
		return 0
	}
}

func (g *Game) LivingNeighbours(x, y int) int {
	livingNeighbours := 0
	for yDelta := -1; yDelta <= 1; yDelta++ {
		for xDelta := -1; xDelta <= 1; xDelta++ {
			if yDelta == 0 && xDelta == 0 {
				continue
			}
			livingNeighbours += g.safeGet(x+xDelta, y+yDelta).Int()
		}
	}
	return livingNeighbours
}

func (g Game) NextGen(cell Cell, livingNeighbours int) Cell {
	if !cell && livingNeighbours == 3 {
		// Gets born
		return true
	}
	if cell && livingNeighbours < 2 {
		// Dead by loneliness
		return false
	}
	if cell && (livingNeighbours == 2 || livingNeighbours == 3) {
		// Stay alive
		return true
	}
	if cell && livingNeighbours > 3 {
		// Overpopulation
		return false
	}
	// Keep status
	return cell
}

func (g Game) safeGet(x, y int) Cell {
	if y < 0 || y >= len(g) {
		return false
	}
	if x < 0 || x >= len(g[y]) {
		return false
	}
	return g[y][x]
}

func (g Game) Round() {
	for y := 0; y < len(g); y++ {
		for x := 0; x < len(g[y]); x++ {
			g[y][x] = g.NextGen(g[y][x], g.LivingNeighbours(x, y))
		}
	}
}

func (g Game) String() string {
	output := make([]string, len(g))
	for y := 0; y < len(g); y++ {
		var line strings.Builder
		for x := 0; x < len(g[y]); x++ {
			line.WriteString(g[y][x].String())
		}
		output[y] = line.String()
	}
	return strings.Join(output, "\n")
}

func main() {
	g := Game{{false, false, false}, {false, false, false}, {true, true, true}}
	fmt.Println("Round 1")
	fmt.Println(g.String())
	g.Round()
	fmt.Println("Round 2")
	fmt.Println(g.String())
}