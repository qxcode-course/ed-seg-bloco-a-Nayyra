package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type Node struct{
	info int
	next *Node
	prev *Node
}

type LList struct{
	head *Node
	size int
}

func NewLList() *LList{ //criando lista vazia
	ll := LList{}
	ll.head = &Node{} 
	ll.head.next = ll.head
	ll.head.prev = ll.head
	ll.size = 0
	return &ll
}

func (list *LList) String() string{
	result := "["
	root := 	list.head.prev
	for root != list.head{
		result += fmt.Sprintf("%d", root.info)
		if root.prev != list.head {
			result += ", "
		}
		root = root.prev
	} 
		
	result += "]"
	return result
}

func (list *LList) Size() int{
	return list.size
}

func (list *LList) PushFront(value int) {
	node := &Node{
		info: value,
	}

	ultimo := list.head.prev
	node.next = list.head
	node.prev = ultimo

	ultimo.next = node
	list.head.prev = node

	list.size++
}

func (list *LList) PushBack(value int) {
	node := &Node{
		info: value,
	}

	ultimo := list.head.next
	node.prev = list.head
	node.next = ultimo

	ultimo.prev = node
	list.head.next = node

	list.size++
}

func (list *LList) PopBack(){
	if list.size == 0{
		return
	}

	prim := list.head.next
	segun := prim.next

	list.head.next = segun
	segun.prev = list.head

	list.size--
}

func (list *LList) PopFront(){
	if list.size == 0{
		return
	}

	prim := list.head.prev
	segun := prim.prev

	list.head.prev = segun
	segun.next = list.head

	list.size--
}

func (list *LList) Clear(){
	list.head.prev = list.head
	list.head.next = list.head
	list.size = 0
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
		case "size":
			fmt.Println(ll.Size())
		case "push_back":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushBack(num)
			}
		case "push_front":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushFront(num)
			}
		case "pop_back":
			ll.PopBack()
		case "pop_front":
			ll.PopFront()
		case "clear":
			ll.Clear()
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
