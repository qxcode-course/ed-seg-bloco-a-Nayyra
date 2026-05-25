package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)


type Set struct{
	data []int
	size int
	capacity int
}

func NewSet(value int) *Set{ //construtor
	return &Set{
		data:     make([]int, value), // nunca use len(data) ou cap(data) ou qq método do go de manipulação de array
		size:     0,
		capacity: value,
	}
}

func (s *Set) insert(index int, value int) error {
	if index < 0 || index > s.size {
		return fmt.Errorf("invalid index")
	}

	if s.size == s.capacity {
		if s.capacity == 0 {
			s.capacity = 1
		} else {
			s.capacity *= 2
		}

		newData := make([]int, s.capacity)

		for i := 0; i < s.size; i++ {
			newData[i] = s.data[i]
		}

		s.data = newData
	}

	for i := s.size; i > index; i-- {
		s.data[i] = s.data[i-1]
	}

	s.data[index] = value

	s.size++

	return nil
}

func (s *Set) Insert(value int) error {

	// impede duplicatas
	for i := 0; i < s.size; i++ {
		if s.data[i] == value {
			return nil
		}
	}

	index := 0

	for index < s.size && s.data[index] < value {
		index++
	}

	return s.insert(index, value)
}

func (s *Set) String() string{
	str := "["
	for i := 0; i < s.size; i++{
		str += fmt.Sprint(s.data[i])
		if i < s.size-1 {
        str += ", "
    	}
	}
	
	str += "]"
	return str
}

func (s *Set) Contains(value int) bool{
	for i := 0; i < s.size; i++{
		if value == s.data[i]{
			return true
		}
	}
	return false
}

func (s *Set) Erase(value int) error{
	if !s.Contains(value) {
		return fmt.Errorf("value not found")
	}

	index := -1

	for i := 0; i < s.size; i++ {
		if s.data[i] == value {
			index = i
			break
		}
	}

	newData := make([]int, s.capacity)

	for i := 0; i < index; i++ {
		newData[i] = s.data[i]
	}

	for i := index + 1; i < s.size; i++ {
		newData[i-1] = s.data[i]
	}

	s.data = newData
	s.size--

	return nil
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	s := NewSet(0)
	for scanner.Scan() {
		fmt.Print("$")
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
			s = NewSet(value)
		case "insert":
			for _, part := range parts[1:] {
				value, err := strconv.Atoi(part)

				if err != nil {
					fmt.Println("invalid number")
					continue
				}

				err = s.Insert(value)

				if err != nil {
					fmt.Println(err)
				}
			}
		case "show":
			fmt.Println(s.String())
		case "erase":
			index, _ := strconv.Atoi(parts[1])
			err := s.Erase(index)
			if err != nil {
				fmt.Println(err)
			}
		case "contains":
			value, _ := strconv.Atoi(parts[1])
			bo := s.Contains(value)
			fmt.Println(bo)
		case "clear":
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
