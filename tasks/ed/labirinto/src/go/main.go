package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	l, c int
}

func getNeig(p Pos) []Pos {
	return []Pos{{p.l, p.c - 1}, {p.l - 1, p.c}, {p.l, p.c + 1}, {p.l + 1, p.c}}
}

func inside(grid [][]rune, p Pos) bool {
	return !(p.l < 0 || p.l >= len(grid) || p.c < 0 || p.c >= len(grid[0]))
}

func match(grid [][]rune, p Pos, value rune) bool {
	return inside(grid, p) && grid[p.l][p.c] == value
}

// Função recursiva que tenta encontrar o caminho do início ao fim
func search(grid [][]rune, l, c Pos) bool {
	// _, _, _ = grid, startPos, endPos
	if l < 0 || l >= len(grid) || c < 0 || c >= len(grid[0]){
		return false
	}
	if grid[l][c] == '#'{
		return false 
	}

	if grid[l][c] == '.'{
		return false
	}

	if grid[l][c] == 'F'{
		return
	}

	grid[l][c] = '.'

	if search(grid, l-1, c) ||
	 search(grid, l-1, c) || 
	 search(grid, l, c-1)  ||
	  search(grid, l, c+1){
		return true 
	}

	grid[l][c] = ' '
	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nl_nc := scanner.Text()
	var nl, nc int
	fmt.Sscanf(nl_nc, "%d %d", &nl, &nc)
	grid := make([][]rune, nl)

	// Lê a gridriz
	for i := range nl {
		scanner.Scan()
		grid[i] = []rune(scanner.Text())
	}

	// Procura posições de início e endPos e conserta para _
	var startPos, endPos Pos
	for l := range nl {
		for c := range nc {
			if grid[l][c] == 'I' {
				grid[l][c] = ' '
				startPos = Pos{l, c}
			}
			if grid[l][c] == 'F' {
				grid[l][c] = ' '
				endPos = Pos{l, c}
			}
		}
	}

	search(grid, startPos, endPos)

	// Imprime o labirinto final
	for _, line := range grid {
		fmt.Println(string(line))
	}
}
