package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func processa(vet []int) {
	_ = vet;
	// 1. defina o ponto de parada
	if len(vet) < 1{
		return 
	}

	// 2. monte o vetor auxiliar com os resultados das somas
	superior := make([]int, 0)
	for i := 0; i < len(vet)-1; i++{
		soma := vet[i] + vet[i+1]
		superior = append(superior, soma)
	}



	// 3. chame recursivamente a função processa para o vetor auxiliar
	processa(superior)
	// 4. imprima o vetor original
	fmt.Print("[ ")
	for _, v := range vet{
		fmt.Printf("%d ", v)
	}
	fmt.Printf("]\n")

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return
	}
	line := scanner.Text()
	parts := strings.Fields(line)
	vet := []int{}
	for _, part := range parts {
		if value, err := strconv.Atoi(part); err == nil {
			vet = append(vet, value)
		}
	}
	processa(vet)
}

func Join[T any](v []T, sep string) string {
	if len(v) == 0 {
		return ""
	}
	s := ""
	for i, x := range v {
		if i > 0 {
			s += sep
		}
		s += fmt.Sprintf("%v", x)
	}
	return s
}
