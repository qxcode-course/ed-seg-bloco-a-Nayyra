package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BetterSearch(slice []int, value int) (bool, int) {
	if len(slice) == 0{
		return false, 0
	}
	inicio := 0
	fim := len(slice)-1

	for inicio <= fim{
		mei := (fim+inicio)/2
		if slice[mei] == value{
			return true, mei
		}

		if slice[mei] < value{
			inicio = mei +1
		}

		if slice[mei] > value{
			fim = mei - 1
		}
	}

	return false, inicio

}

func main() {
	scanner := bufio.NewScanner(os.Stdin) 
	scanner.Scan()
	parts := strings.Split(scanner.Text(), " ")
	slice := []int{}
	for _, elem := range parts[1 : len(parts)-1] {
		value, _ := strconv.Atoi(elem)
		slice = append(slice, value)
	}
	scanner.Scan()
	value, _ := strconv.Atoi(scanner.Text())
	found, result := BetterSearch(slice, value)
	if found {
		fmt.Println("V", result)
	} else {
		fmt.Println("F", result)
	}
}
