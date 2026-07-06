package main

import (
	"bufio"
	"fmt"
	"os"
)

// Não mude a assinatura desta função, ela é a função chamada pelo LeetCode
func exist(grid [][]byte, word string) bool {
	if len(grid) == 0{
		return false
	}

	m := len(grid)
	n := len(grid[0])

	for i:= 0; i < m; i++{
		for j :=0; j < n; j++{
			if backtrack(grid, word, i, j, 0){
				return true
			}
		}
	}
	return false
}

func backtrack(grid [][]byte, word string, l int, c int, i int) bool{
	if i == len(word){
		return true
	}

	if l < 0 || l >= len(grid) || c < 0 || c >= len(grid[0]) || grid[l][c] != word[i]{
		return false
	}

	letraO := grid[l][c]
	grid[l][c] = '*'

	if backtrack(grid, word, l-1, c, i+1)|| 
			backtrack(grid, word, l+1, c, i+1) ||
			backtrack(grid, word, l, c+1, i+1) ||
			backtrack(grid, word, l, c-1, i+1){
				return true
	}

	grid[l][c] = letraO

	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var word string
	fmt.Sscanf(scanner.Text(), "%s", &word)
	grid := make([][]byte, 0)
	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
	}
	if exist(grid, word) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
