package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getMen(vet []int) []int {
	newVet := make([]int, 0)
	for _, v := range vet{
		if v > 0 {
			newVet = append(newVet, v)
		}
	}

	return newVet
}

func getCalmWomen(vet []int) []int {
	newVet := make([]int, 0)
	for _, v := range vet{
		if v < 0 && v > -10 {
			newVet = append(newVet, v)
		}
	}

	return newVet
}

func sortVet(vet []int) []int {

	newVet := make([]int, 0)
	
	for _, v := range vet{
		newVet = append(newVet, v)
		
		j := len(newVet) - 1

		for j > 0 && newVet[j] < newVet[j-1]{
			newVet[j], newVet[j-1] = newVet[j-1], newVet[j]
			j--
		}
	}

	return newVet
}

func sortStress(vet []int) []int {

	newVet := make([]int, 0)
	negaP := make([]int, 0)
	merje := make([]int, 0)
	
	for _, v := range vet{
		if v < -49{
			negaP = append(negaP, v)
		}else{
				newVet = append(newVet, v)
				
				j := len(newVet) - 1

				for j > 0 && newVet[j] < newVet[j-1]{
					newVet[j], newVet[j-1] = newVet[j-1], newVet[j]
					j--
				}
			
		}
	}

	for _, v := range newVet{
		merje = append(merje, v)
	}

	for _, v := range negaP{
		merje = append(merje, v)
	}

	return merje
}

func reverse(vet []int) []int {
	vetReverse := make([]int, 0)
	for i := len(vet)-1; i >= 0; i--{
		vetReverse = append(vetReverse, vet[i])

	} 

	return vetReverse
}

func unique(vet []int) []int {
	mape := make(map[int]bool)
	resul := make([]int, 0)

	for _, v := range vet{
		if _, e := mape[v]; !e{
			mape[v] = true
			resul = append(resul, v)
		}

	}
	return resul
}

func repeated(vet []int) []int {
		repetidas := make([]int, 0)
		unicas := make(map[int]bool)

	   for _, v := range vet{
			_, existe := unicas[v]

			if existe{
					repetidas = append(repetidas, v)
			}else{
					unicas[v] = true
			}
   	}

		return repetidas
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if !scanner.Scan() {
			break
		}
		fmt.Print("$")
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "get_men":
			printVec(getMen(str2vet(args[1])))
		case "get_calm_women":
			printVec(getCalmWomen(str2vet(args[1])))
		case "sort":
			printVec(sortVet(str2vet(args[1])))
		case "sort_stress":
			printVec(sortStress(str2vet(args[1])))
		case "reverse":
			array := str2vet(args[1])
			other := reverse(array)
			printVec(array)
			printVec(other)
		case "unique":
			printVec(unique(str2vet(args[1])))
		case "repeated":
			printVec(repeated(str2vet(args[1])))
		case "end":
			return
		}
	}
}

func printVec(vet []int) {
	fmt.Print("[")
	for i, val := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(val)
	}
	fmt.Println("]")
}

func str2vet(s string) []int {
	if s == "[]" {
		return nil
	}
	s = s[1 : len(s)-1]
	parts := strings.Split(s, ",")
	var vet []int
	for _, part := range parts {
		n, _ := strconv.Atoi(part)
		vet = append(vet, n)
	}
	return vet
}

