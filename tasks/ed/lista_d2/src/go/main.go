package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type Node struct{
	value int
	next *Node
	prev *Node
	root *Node
}

type LList struct{
	root *Node
	size int
}

func NewLList() *LList{
	ll := LList{}
	ll.root = &Node{}
	ll.root.prev = ll.root
	ll.root.next = ll.root
	ll.root.root = ll.root
	ll.size = 0
	return &ll
}

func (node *Node) Next() *Node{
	if node.next == node.root{
		return nil
	}
	return node.next
}

func (node *Node) Prev() *Node{
	if node.prev == node.root{
		return nil
	}
	return node.prev
}

func (list *LList) Back() *Node{
	if list.size == 0{
		return nil
	}

	return list.root.prev
}

func (list *LList) Front() *Node{
	if list.size == 0{
		return nil
	}

	return list.root.next
}


func (list *LList) String() string{
	result := "["
	for node := list.Front(); node != nil; node = node.Next() {
   	result += fmt.Sprintf("%d", node.value)

		if node.Next() != nil {
			result += ", "
		}
	}

	result += "]"
	return result
}

func (list *LList) PushBack(eita int){
	node := &Node{
		value: eita,
		root:  list.root,
	}

	last := list.root.prev

	node.prev = last
	node.next = list.root

	last.next = node
	list.root.prev = node

	list.size++
}

func (list *LList) PushFront(eita int){
	node := &Node{
		value: eita,
		root:  list.root,
	}

	first := list.root.next

	node.next = first
	node.prev = list.root

	first.prev = node
	list.root.next = node

	list.size++
}

func (list *LList) Size() int{
	return list.size
}

func (list *LList) Clear(){
	list.root.prev = list.root
	list.root.next = list.root
	list.size = 0
}

func (list *LList) Remove(node *Node) *Node{
	
			antes := node.prev
			depois := node.next

			antes.next = depois
			depois.prev = antes

			list.size--
			return node.next
}

func (list *LList) Insert(node *Node, valuee int) {
	newNode := &Node{
		value: valuee,
		root: list.root,
	}
		anterior := node.prev

		node.prev = newNode
		newNode.next = node
		anterior.next = newNode
		newNode.prev = anterior
}

func (list *LList) Search(Value int) *Node{
	for node := list.Front(); node != nil; node = node.Next() {
   	if node.value == Value{
			return node
		}
	}
	return nil

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
			// ll.PopBack()
		case "pop_front":
			// ll.PopFront()
		case "clear":
			ll.Clear()
		case "walk":
			fmt.Print("[ ")
			for node := ll.Front(); node != nil; node = node.Next() {
				fmt.Printf("%v ", node.value)
			}
			fmt.Print("]\n[ ")
			for node := ll.Back(); node != nil; node = node.Prev() {
				fmt.Printf("%v ", node.value)
			}
			fmt.Println("]")
		case "replace":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
				node.value = newvalue
			} else {
				fmt.Println("fail: not found")
			}
		case "insert":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
				ll.Insert(node, newvalue)
			} else {
				fmt.Println("fail: not found")
			}
		case "remove":
			oldvalue, _ := strconv.Atoi(args[1])
			node := ll.Search(oldvalue)
			if node != nil {
				ll.Remove(node)
			} else {
				fmt.Println("fail: not found")
			}
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
