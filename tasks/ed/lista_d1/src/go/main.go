package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Node struct{
	info int
	next *Node
	prev *Node
}

type LList struct{
	head *Node
}

func NewLList() *LList{ //criando lista vazia
	ll := LList{}
	ll.head = &Node{} 
	ll.head.next = ll.head
	ll.head.prev = ll.head
	return &ll
}

func (list *LList) String() *LList{
	root := head
	
	for root != nil{
		return root
	}
	return eita
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	ll := NewLList()

	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		fmt.Println(line)
		args := strings.Fields(line)

		if len(args) == 0 {
			continue
		}

		cmd := args[0]

		switch cmd {
		case "show":
			fmt.Println(ll.String())
		// case "size":
		// 	fmt.Println(ll.Size())
		// case "push_back":
		// 	for _, v := range args[1:] {
		// 		num, _ := strconv.Atoi(v)
		// 		ll.PushBack(num)
		// 	}
		// case "push_front":
		// 	for _, v := range args[1:] {
		// 		num, _ := strconv.Atoi(v)
		// 		ll.PushFront(num)
		// 	}
		// case "pop_back":
		// 	ll.PopBack()
		// case "pop_front":
		// 	ll.PopFront()
		// case "clear":
		// 	ll.Clear()
		// case "end":
		// 	return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
