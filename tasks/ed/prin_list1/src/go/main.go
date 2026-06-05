	package main

	import (
		"fmt"
		"strings"
	)

	// mostra a lista com o elemento sword destacado
	func ToStr(l *DList[int], sword *DNode[int]) string {
		result := "[ "
		for node := l	.Front(); node != l.End(); node = node.Next() {
			if node == sword{
				result += fmt.Sprintf("%d> ", node.Value)
			}else{
				result += fmt.Sprintf("%d ", node.Value)

			}

		} 
		result += "]"
		return result
	}

	// move para frente na lista circular
	func Next(l *DList[int], it *DNode[int]) *DNode[int] {
		if it == nil || it == l.End() {
      	return l.Front()
   	}

   	next := it.Next()

		if next == l.End() {
      	return l.Front()
   	}

   	return next

	}


	func main() {
		var qtd, chosen int
		fmt.Scan(&qtd, &chosen)
		fmt.Println(qtd, chosen)
		l := NewDList[int]()
		for i := 1; i <= qtd; i++ {
			l.PushBack(i)
		}
		sword := l.Front()
		for range chosen - 1 {
			sword = Next(l, sword)
		}
		for range qtd - 1 {
			fmt.Println(ToStr(l, sword))
			vitima := Next(l, sword)
			nextEspada := Next(l, vitima)

			if nextEspada == vitima || nextEspada == l.End() {
				nextEspada = sword
			}

			l.Erase(vitima)
			sword = nextEspada
		}
		fmt.Println(ToStr(l, sword))
	}
