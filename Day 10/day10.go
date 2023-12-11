package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type grid [][]rune

func (g *grid) Get(c coordinate) rune {
	return (*g)[c.Row()][c.Col()]
}

type coordinate [2]int

func (c *coordinate) Row() int {
	return c[0]
}

func (c *coordinate) Col() int {
	return c[1]
}

func (c *coordinate) String() string {
	return fmt.Sprintf("(%v, %v)", c.Row(), c.Col())
}

func (c *coordinate) Equal(o *coordinate) bool {
	return c.Col() == o.Col() && c.Row() == o.Row()
}

const (
	PIPE_N_S  rune = '|'
	PIPE_E_W       = '-'
	PIPE_N_E       = 'L'
	PIPE_N_W       = 'J'
	PIPE_S_W       = '7'
	PIPE_S_E       = 'F'
	PIPE_NONE      = '.'
)

type direction uint

const (
	NORTH direction = iota
	EAST
	SOUTH
	WEST
)

type pipe struct {
	pos    coordinate
	symbol rune
	north  *pipe
	east   *pipe
	south  *pipe
	west   *pipe
}

func (p *pipe) NextDirection(dir direction) direction {
	if p.symbol == PIPE_N_S {
		if dir == NORTH {
			return NORTH
		} else {
			return SOUTH
		}
	}
	if p.symbol == PIPE_N_E {
		if dir == SOUTH {
			return EAST
		} else {
			return NORTH
		}
	}
	if p.symbol == PIPE_N_W {
		if dir == SOUTH {
			return WEST
		} else {
			return NORTH
		}
	}
	if p.symbol == PIPE_S_E {
		if dir == NORTH {
			return EAST
		} else {
			return SOUTH
		}
	}
	if p.symbol == PIPE_S_W {
		if dir == NORTH {
			return WEST
		} else {
			return SOUTH
		}
	}
	if p.symbol == PIPE_E_W {
		if dir == EAST {
			return EAST
		} else {
			return WEST
		}
	}

	panic(fmt.Sprintf("unknown symbol %s", string(p.symbol)))
}

func findStartPos(grid grid) coordinate {
	for r, row := range grid {
		for c, col := range row {
			if col == 'S' {
				return coordinate{r, c}
			}
		}
	}
	panic("Start not found")
}

func determineStartType(grid grid, startPos coordinate) rune {
	var connectsNorth, connectsEast, connectsSouth, connectsWest bool

	northPipe := grid.Get(getPos(startPos, NORTH))
	eastPipe := grid.Get(getPos(startPos, EAST))
	southPipe := grid.Get(getPos(startPos, SOUTH))
	westPipe := grid.Get(getPos(startPos, WEST))

	connectsNorth = northPipe == PIPE_N_S || northPipe == PIPE_S_E || northPipe == PIPE_S_W
	connectsEast = eastPipe == PIPE_E_W || eastPipe == PIPE_S_W || eastPipe == PIPE_N_W
	connectsSouth = southPipe == PIPE_N_S || southPipe == PIPE_N_E || southPipe == PIPE_N_W
	connectsWest = westPipe == PIPE_E_W || westPipe == PIPE_N_E || westPipe == PIPE_S_E

	switch {
	case connectsNorth && connectsSouth:
		return PIPE_N_S
	case connectsNorth && connectsEast:
		return PIPE_N_E
	case connectsNorth && connectsWest:
		return PIPE_N_W
	case connectsSouth && connectsEast:
		return PIPE_S_E
	case connectsSouth && connectsWest:
		return PIPE_S_W
	case connectsEast && connectsWest:
		return PIPE_E_W
	}
	panic(fmt.Sprintf("unable to determine start type: %t %t %t %t", connectsEast, connectsEast, connectsSouth, connectsWest))
}

func getStartDirection(startSymbol rune, forward bool) direction {
	switch startSymbol {
	case PIPE_N_S:
		if forward {
			return NORTH
		} else {
			return SOUTH
		}
	case PIPE_N_E:
		if forward {
			return NORTH
		} else {
			return EAST
		}
	case PIPE_N_W:
		if forward {
			return NORTH
		} else {
			return WEST
		}
	case PIPE_S_E:
		if forward {
			return SOUTH
		} else {
			return EAST
		}
	case PIPE_S_W:
		if forward {
			return SOUTH
		} else {
			return WEST
		}
	case PIPE_E_W:
		if forward {
			return EAST
		} else {
			return WEST
		}
	}
	panic(fmt.Sprintf("unknown symbol %s", string(startSymbol)))
}

func getPos(currentPos coordinate, dir direction) coordinate {
	switch dir {
	case NORTH:
		return coordinate{currentPos.Row() - 1, currentPos.Col()}
	case EAST:
		return coordinate{currentPos.Row(), currentPos.Col() + 1}
	case SOUTH:
		return coordinate{currentPos.Row() + 1, currentPos.Col()}
	case WEST:
		return coordinate{currentPos.Row(), currentPos.Col() - 1}
	default:
		panic(fmt.Sprintf("invalid direction: %v", dir))
	}
}

func makeNodes(startNode *pipe, grid grid) int {
	node := startNode
	dir := getStartDirection(startNode.symbol, false)
	length := 0
	for {
		pos := getPos(node.pos, dir)
		fmt.Printf("Walk Direction: %v Pos: %v\n", dir, pos)
		symbol := grid.Get(pos)
		if symbol == 'S' {
			break
		}
		nextNode := pipe{pos: pos, symbol: symbol}
		switch dir {
		case NORTH:
			node.north = &nextNode
			nextNode.south = node
		case EAST:
			node.east = &nextNode
			nextNode.west = node
		case SOUTH:
			node.south = &nextNode
			nextNode.north = node
		case WEST:
			node.west = &nextNode
			nextNode.east = node
		}

		node = &nextNode
		dir = node.NextDirection(dir)
		length++
	}
	return length
}

func main() {
	// Open the file
	file, err := os.Open("input")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	grid := make(grid, 0)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}

		row := make([]rune, 0)
		for _, r := range line {
			row = append(row, r)
		}
		grid = append(grid, row)
	}

	fmt.Println(grid)
	startPos := findStartPos(grid)
	fmt.Printf("Start Pos: %v\n", startPos)
	startType := determineStartType(grid, startPos)
	fmt.Printf("Start Type: %v\n", string(startType))

	startNode := pipe{pos: startPos, symbol: startType}
	length := makeNodes(&startNode, grid)

	fmt.Println((length + 1) / 2)
}
