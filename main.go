package main

import (
	"fmt"
)

type LinkedList struct {
	head *Node
}

type Node struct {
	value int
	next  *Node
}

func (l *LinkedList) addFirst(value int) {
	l.head = &Node{value: value, next: l.head}
}

func (l *LinkedList) addLast(value int) {
	node := l.head
	for node.next != nil {
		node = node.next
	}
	node.next = &Node{value: value, next: nil}
}

func (l *LinkedList) delValue(value int) {
	node := l.head
	for i := 0; node != nil; i++ {
		if node.value == value {
			switch {
			case l.previus(i) != nil && l.next(i) != nil:
				l.previus(i).next = l.next(i)
			case l.previus(i) == nil && l.next(i) != nil:
				l.head = l.next(i)
			case l.previus(i) != nil && l.next(i) == nil:
				l.previus(i).next = nil
			case l.previus(i) == nil && l.next(i) == nil:
				l.head = nil
			}
			return
		}
		node = node.next
	}
}

func (l *LinkedList) delIndex(index int) {
	node := l.head
	for i := 0; node != nil; i++ {
		if i == index {
			switch {
			case l.previus(i) != nil && l.next(i) != nil:
				l.previus(i).next = l.next(i)
			case l.previus(i) == nil && l.next(i) != nil:
				l.head = l.next(i)
			case l.previus(i) != nil && l.next(i) == nil:
				l.previus(i).next = nil
			case l.previus(i) == nil && l.next(i) == nil:
				l.head = nil
			}
			return
		}
		node = node.next
	}
}

func (l *LinkedList) node(index int) *Node {
	node := l.head
	for i := 0; node != nil; i++ {
		if i == index {
			return node
		}
		node = node.next
	}
	return nil
}

func (l *LinkedList) index(value int) int {
	node := l.head
	for i := 0; node != nil; i++ {
		if node.value == value {
			return i
		}
		node = node.next
	}
	return -1
}

func (l *LinkedList) value(index int) int {
	node := l.head
	for i := 0; node != nil; i++ {
		if i == index {
			return node.value
		}
		node = node.next
	}
	return -1
}

func (l *LinkedList) previus(index int) *Node {
	node := l.head
	for i := 0; node != nil; i++ {
		if i == index-1 {
			return node
		}
		node = node.next
	}
	return nil
}

func (l *LinkedList) next(index int) *Node {
	node := l.head
	for i := 0; node != nil; i++ {
		if i == index+1 {
			return node
		}
		node = node.next
	}
	return nil
}

func (l *LinkedList) len() int {
	node := l.head
	len := 0
	for ; node != nil; len++ {
		node = node.next
	}
	return len
}

func (l *LinkedList) insert(index, value int) {
	node := l.head
	for i := 0; node != nil; i++ {
		if i == index {
			node = &Node{value: value, next: node}
			if index == 0 {
				l.head = node
			}
			if l.previus(i) != nil {
				l.previus(i).next = node
			}
			return
		}
		node = node.next
	}
}

func (l *LinkedList) swap(left, right int) {
	leftPrev := l.previus(left)
	leftNode := l.node(left)
	leftNext := l.next(left)

	rightPrev := l.previus(right)
	rightNode := l.node(right)
	rightNext := l.next(right)

	leftNode.next = rightNext
	if rightPrev == leftNode {
		rightNode.next = leftNode
	} else {
		rightPrev.next = leftNode
		rightNode.next = leftNext
	}

	if leftPrev == nil {
		l.head = rightNode
	} else {
		leftPrev.next = rightNode
	}
}

func (l *LinkedList) printList() {
	node := l.head
	for i := 0; node != nil; i++ {
		fmt.Printf("%d: %d\n", i, node.value)
		node = node.next
	}
	fmt.Println("--------")
}

func (l *LinkedList) sort() {
	quickSortByCenter(l, 0, l.len()-1)
}

func quickSortByCenter(l *LinkedList, start, end int) {
	left := start
	right := end
	center := l.value((left + right) / 2)
	for left <= right {
		for l.value(right) > center {
			right--
		}
		for l.value(left) < center {
			left++
		}
		if left <= right {
			l.swap(left, right)
			left++
			right--
		}
	}
	if right > start {
		quickSortByCenter(l, start, right)
	}
	if left < end {
		quickSortByCenter(l, left, end)
	}
}

func main() {
	fmt.Println("The linked list is initialized with 1 element with the value 5")
	list := LinkedList{head: &Node{value: 5, next: nil}}
	fmt.Printf("Lenght of list: %d\n", list.len())
	list.printList()

	fmt.Println("Insert an element at the beginning of a linked list: 3, 2, 8")
	list.addFirst(3)
	list.addFirst(2)
	list.addFirst(8)
	list.printList()

	fmt.Println("Insert an element at the end of a linked list: 4, 7, 6")
	list.addLast(4)
	list.addLast(7)
	list.addLast(6)
	list.printList()

	fmt.Println("index of value 3: ", list.index(3))
	fmt.Println("index of value 7: ", list.index(7))
	fmt.Println("index of value 8: ", list.index(8))
	fmt.Println("index of value 6: ", list.index(6))
	fmt.Println("--------")

	fmt.Println("value of index 0: ", list.value(0))
	fmt.Println("value of index 3: ", list.value(3))
	fmt.Println("value of index 5: ", list.value(5))
	fmt.Println("value of index 6: ", list.value(6))
	fmt.Println("--------")

	fmt.Println("Inset element at 1")
	list.insert(1, 9)
	list.printList()

	fmt.Printf("Lenght of list: %d\n", list.len())
	fmt.Println("--------")

	fmt.Println("Delete element with value of 5")
	list.delValue(5)
	list.printList()

	fmt.Println("Delete element with value of 8")
	list.delValue(8)
	list.printList()

	fmt.Println("Delete element with value of 6")
	list.delValue(6)
	list.printList()

	fmt.Println("Delete element at index 2")
	list.delIndex(2)
	list.printList()

	fmt.Println("Delete element at index 0")
	list.delIndex(0)
	list.printList()

	fmt.Println("Delete element at index 2 - last element")
	list.delIndex(2)
	list.printList()

	fmt.Println("new elements added")
	list.insert(1, 6)
	list.insert(1, 1)
	list.insert(2, 3)
	list.insert(0, 9)
	list.printList()

	fmt.Println("Sorted linked list")
	list.sort()
	list.printList()
}
