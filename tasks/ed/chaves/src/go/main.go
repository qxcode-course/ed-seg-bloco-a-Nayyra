package main

import "fmt"

func main() {
	fila := NewQueue[string]()
	for letra := 'A'; letra <= 'P'; letra++{
		fila.Enqueue(string(letra))
	}

	for range 15{
		t1 := fila.Dequeue()
		t2 := fila.Dequeue()

		var gol1, gol2 int
		fmt.Scan(&gol1, &gol2)

		if gol1 > gol2{
			fila.Enqueue(t1)
		}else{
			fila.Enqueue(t2)
		}
	}

	fmt.Println(fila.Dequeue())
}
