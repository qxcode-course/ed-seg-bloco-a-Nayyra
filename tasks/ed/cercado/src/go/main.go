package main

import (
	"bufio"
	"fmt"
	"os"
)

type Position struct{
	c int
	l int
}

func dfs(board [][]byte, mapa map[Position]bool, l, c int){
	if l < 0 || l >= len(board) || c < 0 || c >= len(board[0]){
		return
	}

	if board[l][c] == 'X'{
		return
	}

	if mapa[Position{l, c}]{
		return
	}

	mapa[Position{l, c}] = true

	dfs(board, mapa, l-1, c)
	dfs(board, mapa, l+1, c)
	dfs(board, mapa, l, c-1)
	dfs(board, mapa, l, c+1)
}

// NÃO ALTERE A ASSINATURA DA FUNÇÃO solve
func solve(board [][]byte) {
	mapa := make(map[Position]bool)
	//linha //registra no mapa a posição 
	// for i := 0; i < len(board); i++{
	// 	if i == 'O'{
	// 		mapa[i] = true
	// 	}
	// }
	
	for i := 0; i < len(board[0]); i++{
		if board[0][i] == 'O'{
			dfs(board, mapa, 0, i)
		}
	}

	//coluna
		for j := 0; j < len(board); j++{
		if board[j][0] == 'O'{
			dfs(board, mapa, j, 0)
		}
	}


	//linha inversa
	for j := 0; j < len(board); j++{
		if board[len(board)-1][j] == 'O'{
			dfs(board, mapa, len(board)-1, j)
		}
	}

	//coluna inversa
	for i := 0; i < len(board[0]); i++{
		if board[i][len(board)-1] == 'O'{
			dfs(board, mapa, i, len(board[0])-1)
		}
	}

	for i := 0; i < len(board); i++{
		for j := 0; j < len(board[0]); j++{
			if  board[i][j] == 'O' && !mapa[Position{i, j}] {
				board[i][j] = 'X'
			}
		}
	} 
}


// NÃO ALTERE A MAIN
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var nrows, ncols int
	fmt.Sscanf(scanner.Text(), "%d %d", &nrows, &ncols)
	board := make([][]byte, nrows)
	for i := 0; i < nrows; i++ {
		scanner.Scan()
		board[i] = []byte(scanner.Text())
	}
	solve(board)
	for _, row := range board {
		fmt.Println(string(row))
	}
}
