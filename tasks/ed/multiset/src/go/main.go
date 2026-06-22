package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type MultiSet struct {
	data     []int //memória dinâmica
	size int
	capacity int
}

func NewMultiSet(capacity int) *MultiSet{
	return &MultiSet{
		data: make([]int, capacity),
		size: 0,
		capacity: capacity,
	}
}

func (ms *MultiSet) Show() string{
	str := "["
	for i := 0; i < ms.size; i++{
		str += fmt.Sprint(ms.data[i])

		if i < ms.size-1 {
        str += ", "
    	}
	}
	
	str += "]"
	return str
}

func (ms *MultiSet) Insert(value int){
	posicao := ms.size
	for i:=0; i < ms.size; i++{
		if ms.data[i] >= value{
			posicao = i
			break
		}
	}

	ms.insert(posicao, value)
}

func (ms *MultiSet) insert(index int, value int ) error{
	if index < 0 || index > ms.size {
		return fmt.Errorf("vector is empty")
	}

	if ms.size == ms.capacity {
        ms.capacity *= 2
    }

	newData := make([]int, ms.capacity)
	for i := 0; i < index; i++{
			newData[i] = ms.data[i]
		}
	newData[index] = value

	for i := index; i < ms.size; i++{
			newData[i+1] = ms.data[i]
		}
	
	ms.data = newData
	ms.size++
	
	return nil
}

func (ms *MultiSet) Contains(value int) bool{
	for i := 0; i < ms.size; i++{
		if ms.data[i] == value{
			return true
		}
	}
	return false
}

func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	result := fmt.Sprintf("%d", slice[0])
	for _, value := range slice[1:] {
		result += sep + fmt.Sprintf("%d", value)
	}
	return result
}

func (ms *MultiSet) Erase(value int) error{
	for i:=0; i < ms.size; i++{
		if ms.data[i] == value{
			for j := i; j < ms.size-1; j++ {
				ms.data[j] = ms.data[j+1]
			}
			ms.size--
			return nil
		}
	}
	return fmt.Errorf("value not found")

}

func (ms *MultiSet) Count(value int) int{
	contador := 0
	for i:=0; i < ms.size; i++{
		if ms.data[i] == value{
			contador++
		}
	}

	return contador
}

func (ms *MultiSet) Unique() int{
	if ms.size == 0{
		return 0
	}
	contador := 1
	for i:=1; i < ms.size; i++{
		if ms.data[i] != ms.data[i-1]{
			contador++
		}
	}

	return contador
}

func (ms *MultiSet) Clear(){
	newData := make([]int, 0)
	ms.data = newData
	ms.size = 0
}


func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)
	ms := NewMultiSet(0)

	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		args := strings.Fields(line)
		fmt.Println(line)
		if len(args) == 0 {
			continue
		}
		cmd = args[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(args[1])
			ms = NewMultiSet(value)
		case "insert":
			for _, part := range args[1:] {
				value, _ := strconv.Atoi(part)
				ms.Insert(value)
			}
			
		case "show":
			fmt.Println(ms.Show())
		case "erase":
			value, _ := strconv.Atoi(args[1])
			err := ms.Erase(value)
			if err != nil {
				fmt.Println(err)
			}
		case "contains":
			value, _ := strconv.Atoi(args[1])
			//ms.Contains(value)
			if ms.Contains(value){
				fmt.Println("true")
			}else{
				fmt.Println("false")
			}
		case "count":
			value, _ := strconv.Atoi(args[1])
			fmt.Println(ms.Count(value))
		case "unique":
			fmt.Println(ms.Unique())
		case "clear":
			ms.Clear()
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
