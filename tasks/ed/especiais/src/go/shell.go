package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sort"
)

type Pair struct {
	One int
	Two int
}

func occurr(vet []int) []Pair {
	f := make(map[int]int)
	
	for _, v := range vet{
		if v <0 {
			v = -v
		}
		f[v]++
	}

	ordenado := make([]int, 0, len(f))
	for v := range f{
		ordenado = append(ordenado, v)
	}
	sort.Ints(ordenado)

	r := []Pair{}
	for _, v := range ordenado{
		r = append(r, Pair{v, f[v]})
	}

	return r
}

func teams(vet []int) []Pair {

	if len(vet) == 0 {
		return []Pair{}
	}

	res := []Pair{}

	a := vet[0]
	c := 1

	for i := 1; i < len(vet); i++{
		if vet[i] == a{
			c++
		}else{
			res = append(res, Pair{a, c})
			a = vet[i]
			c = 1
		}
	}

	res = append(res, Pair{a, c})

	return res
}

func mnext(vet []int) []int {
	res := make([]int, len(vet))

	//a := vet[0]
	for i := 0; i < len(vet); i++{
			if vet[i] > 0{
				if i > 0 && vet[i-1] < 0{
					res[i] = 1
				}
				if i < len(vet)-1 && vet[i+1] < 0{
					res[i] = 1
				}
			}
	} 

	return res
}

func alone(vet []int) []int {
	res := make([]int, len(vet))


	//a := vet[0]
	for i := 0; i < len(vet); i++ {

		if vet[i] > 0 {
			mulher := false
			if i > 0 && vet[i-1] < 0 {
				mulher = true
			}

			if i < len(vet)-1 && vet[i+1] < 0 {
				mulher = true
			}

			if !mulher{
				res[i] = 1
			}
		}
	} 

	return res
}

func couple(vet []int) int {

	homens := make(map[int]int)
	mulheres := make(map[int]int)

	for _, v := range vet {
		if v > 0 {
			homens[v]++
		} else {
			mulheres[-v]++
		}
	}

	total := 0

	for k, qtdHomens := range homens {
		qtdMulheres := mulheres[k]

		if qtdHomens < qtdMulheres {
			total += qtdHomens
		} else {
			total += qtdMulheres
		}
	}

	return total
}

func hasSubseq(vet []int, seq []int, pos int) bool {
	_ = vet
	_ = seq
	_ = pos
	return false
}

func subseq(vet []int, seq []int) int {
	_ = vet
	_ = seq
	return -1
}

func erase(vet []int, posList []int) []int {
	_ = vet
	_ = posList
	return nil
}

func clear(vet []int, value int) []int {
	_ = vet
	_ = value
	return nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "occurr":
			printSlice(occurr(str2vet(args[1])))
		case "teams":
			printSlice(teams(str2vet(args[1])))
		case "mnext":
			printSlice(mnext(str2vet(args[1])))
		case "alone":
			printSlice(alone(str2vet(args[1])))
		case "erase":
			printSlice(erase(str2vet(args[1]), str2vet(args[2])))
		case "clear":
			val, _ := strconv.Atoi(args[2])
			printSlice(clear(str2vet(args[1]), val))
		case "subseq":
			fmt.Println(subseq(str2vet(args[1]), str2vet(args[2])))
		case "couple":
			fmt.Println(couple(str2vet(args[1])))
		case "end":
			return
		default:
			fmt.Println("Invalid command")
		}
	}
}

// Funções auxiliares

func str2vet(str string) []int {
	if str == "[]" {
		return nil
	}
	str = str[1 : len(str)-1]
	parts := strings.Split(str, ",")
	var vet []int
	for _, part := range parts {
		num, _ := strconv.Atoi(strings.TrimSpace(part))
		vet = append(vet, num)
	}
	return vet
}

func printSlice[T any](vet []T) {
	fmt.Print("[")
	for i, x := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(x)
	}
	fmt.Println("]")
}

func (p Pair) String() string {
	return fmt.Sprintf("(%v, %v)", p.One, p.Two)
}
