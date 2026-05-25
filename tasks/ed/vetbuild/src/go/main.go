package main

import (
	"bufio"
	"fmt"
	"strings"
	"os"
	"strconv"
)

type Vector struct {
	data     []int //memória dinâmica
	size     int
	capacity int
}

func NewVector(capacity int) *Vector { //Construtor 
	return &Vector{
		data:     make([]int, capacity), // nunca use len(data) ou cap(data) ou qq método do go de manipulação de array
		size:     0,
		capacity: capacity,
	}
}


func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	var result strings.Builder
	fmt.Fprintf(&result, "%d", slice[0])
	for _, value := range slice[1:] {
		fmt.Fprintf(&result, "%s%d", sep, value)
	}
	return result.String()
}

func (v *Vector) Show() string{
	str := "["
	for i := 0; i < v.size; i++{
		str += fmt.Sprint(v.data[i])
		if i < v.size-1 {
        str += ", "
    	}
	}
	
	str += "]"
	return str
}

func (v *Vector) Status() string{
	size := v.size
	capacity := v.capacity

	str := "size:" + fmt.Sprint(size) + " capacity:" + fmt.Sprint(capacity)
	return str

	
}

func (v *Vector) Get(value int) int{
	return v.data[value]
}

func (v *Vector) At(value int) (int, error){
	if value < 0 || value >= v.size{
		return 0, fmt.Errorf("index out of range")
	}
	return v.data[value], nil
}

func (v *Vector) PushBack(value int) {
	if v.size == v.capacity{
		newCapacity := v.capacity*2
		newData := make([]int, newCapacity) 
		
		for i := 0; i < v.size; i++{
			newData[i] = v.data[i]
		}

		v.data = newData
		v.capacity = newCapacity
	}

	v.data[v.size] = value
	v.size++
}

func (v *Vector) Set(index int, value int) error{
	if index < 0 || index >= v.size{
		return fmt.Errorf("index out of range")
	}

	v.data[index] = value
	return nil
}

func (v *Vector) Clear(){
	newData := make([]int, 0)
	v.data = newData
	v.size = 0
}

func (v *Vector) Reserve(newCapacity int){
	v.capacity = newCapacity
}

func (v *Vector) PopBack() error{
	if v.size == 0{
		return fmt.Errorf("vector is empty")
	}

	newSize := v.size-1 
	newData := make([]int, v.capacity)
	for i := 0; i < v.size-1; i++{
			newData[i] = v.data[i]
		}

	v.data = newData
	v.size = newSize
	return nil
}

func (v *Vector) Insert(index int, value int ) error{
	if index < 0 || index > v.size {
		return fmt.Errorf("vector is empty")
	}
	//talvez eu precise aumentar a capacidade

	if v.size == v.capacity {
        v.capacity *= 2
    }

	newData := make([]int, v.capacity)
	for i := 0; i < index; i++{
			newData[i] = v.data[i]
		}
	newData[index] = value

	for i := index; i < v.size; i++{
			newData[i+1] = v.data[i]
		}
	
	v.data = newData
	v.size++
	
	return nil
}

func (v *Vector) Erase(index int ) error{
	if index < 0 || index > v.size {
		return fmt.Errorf("index out of range")
	}

	newData := make([]int, v.capacity)
	for i := 0; i < index; i++{
			newData[i] = v.data[i]
		}

	for i := index; i < v.size-1; i++{
			newData[i] = v.data[i+1]
		}
	
	v.data = newData
	v.size--
	
	return nil
}

func (v *Vector) IndexOf(value int) int{
	for i := 0; i < v.size; i++{
		if v.data[i] == value{
			return i
		}
	}
	return -1
}

func (v *Vector) Contains(value int) bool{
	for i := 0; i < v.size; i++{
		if v.data[i] == value{
			return true
		}
	}
	return false
}

func (v *Vector) Slice(start int, end int) string {

    if end == -1 {
        end = v.size - 1
    }

    if start < 0 || end > v.size || start > end {
        return ""
    }

    newData := make([]int, end-start)

    for i := start; i < end; i++ {
        newData[i-start] = v.data[i]
    }

    newVector := Vector{
        data:     newData,
        size:     len(newData),
        capacity: len(newData),
    }

    return newVector.Show()
}


func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	v := NewVector(0)
	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line = scanner.Text()
		fmt.Println(line)
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		cmd = parts[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(parts[1])
			v = NewVector(value)
		case "push":
			for _, part := range parts[1:] {
				value, _ := strconv.Atoi(part)
				v.PushBack(value)
			}
		case "show":
			fmt.Println(v.Show())
		case "status":
			fmt.Println(v.Status())
		case "pop":
			err := v.PopBack()
			if err != nil {
				fmt.Println(err)
			}
		case "insert":
			index, _ := strconv.Atoi(parts[1])
			value, _ := strconv.Atoi(parts[2])
			err := v.Insert(index, value)
			if err != nil {
				fmt.Println(err)
			}
		case "erase":
			index, _ := strconv.Atoi(parts[1])
			err := v.Erase(index)
			if err != nil {
				fmt.Println(err)
			}
		case "indexOf":
			value, _ := strconv.Atoi(parts[1])
			index := v.IndexOf(value)
			fmt.Println(index)
		case "contains":
			value, _ := strconv.Atoi(parts[1])
			if v.Contains(value) {
				fmt.Println("true")
			} else {
				fmt.Println("false")
			}
		case "clear":
			v.Clear()
		case "capacity":
			// fmt.Println(v.Capacity())
		case "get":
			index, _ := strconv.Atoi(parts[1])
			value, err := v.At(index)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(value)
			}
		case "set":
			index, _ := strconv.Atoi(parts[1])
			value, _ := strconv.Atoi(parts[2])
			err := v.Set(index, value)
			if err != nil {
				fmt.Println(err)
			}
			
		case "reserve":
			newCapacity, _ := strconv.Atoi(parts[1])
			v.Reserve(newCapacity)
		case "slice":
			start, _ := strconv.Atoi(parts[1])
			end, _ := strconv.Atoi(parts[2])
			slice := v.Slice(start, end)
			fmt.Println(slice)
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
