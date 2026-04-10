package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func tostr(vet []int) string {
	//newVet := make([]int, 0)
	s := "["
	if len(vet) != 0{
		for i, v := range vet{
			a := fmt.Sprint(v)
			if i == len(vet)-1 {
				s = s + a +"]"
			}else{
				s = s + a + ", "
			}
		}
	}else{
		s = s + "]"
	}
	return s
}

func tostrrev(vet []int) string {
	s := "["
	if len(vet) != 0{
		for i := len(vet)-1; i >= 0; i--{
			a := fmt.Sprint(vet[i])
			if i == 0 {
				s = s + a +"]"
			}else{
				s = s + a + ", "
			}
		}
	}else{
		s = s + "]"
	}
	return s
}

// reverse: inverte os elementos do slice
func reverse(vet []int){
	for i := 0; i < len(vet)-1; i++{
		j := len(vet) - 1 - i
		vet[i], vet[j] = vet[j], vet[i]
	}
}

// sum: soma dos elementos do slice
func sum(vet []int) int {
	//newVet := make([]int, 0)
	contador := 0
	for _, v := range vet{
		contador += v
	}

	return contador
}

// mult: produto dos elementos do slice
func mult(vet []int) int {
	contador := 1
	for _, v := range vet{
		contador = contador * v
	}

	return contador
}


// min: retorna o índice e valor do menor valor
// crie uma função recursiva interna do modelo
// var rec func(v []int) (int, int)
// para fazer uma recursão que retorna valor e índice
func min(vet []int) int {
	if len(vet) == 0{
		return -1
	}
	 
	var rec func(v []int) (int, int)

	rec = func(v []int) (int, int){
		if len(v) == 1{
			return 0, v[0]
		}
	

		i, val := rec(v[1:])
		i += 1
		if v[0] < val{
			return 0, v[0]
		}
		return i, val
	}
	i, _ := rec(vet)
	return i
}

func main() {
	var vet []int
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Fields(line)
		fmt.Println("$" + line)

		switch args[0] {
		case "end":
			return
		case "read":
			vet = nil
			for _, arg := range args[1:] {
				if val, err := strconv.Atoi(arg); err == nil {
					vet = append(vet, val)
				}
			}
		case "tostr":
			fmt.Println(tostr(vet))
		case "torev":
			fmt.Println(tostrrev(vet))
		case "reverse":
			reverse(vet)
		case "sum":
			fmt.Println(sum(vet))
		case "mult":
			fmt.Println(mult(vet))
		case "min":
			fmt.Println(min(vet))
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
