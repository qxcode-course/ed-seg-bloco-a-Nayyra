package main
import "fmt"
func main() {
	var n int //número inicial de pessoas na fila
	fmt.Scan(&n)
	filaIni := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&filaIni[i])
	}
	var m int //número de pessoas que deixaram a fila
	fmt.Scan(&m)

	fila := make(map[int]bool)
	for i := 0; i < m; i++ {
		var x int
		fmt.Scan(&x)
		fila[x] = true
	}
	restos := make([]int, 0)
	for _, pessoa := range filaIni {
		if !fila[pessoa] {
			restos = append(restos, pessoa)
		}
	}
	for _, pessoa := range restos {
		fmt.Print(pessoa)
		fmt.Print(" ")
		if pessoa == restos[len(restos)-1]{
			fmt.Print("\n")
		}
	}	
	
}