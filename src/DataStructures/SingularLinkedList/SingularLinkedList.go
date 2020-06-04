package main

import "fmt"

type node struct {
	data int
	next *node
}
type list struct {
	head *node
	tail *node
}

func addNode(value int) node {
	newNode := node{
		data: value,
		next: nil,
	}
	return newNode
}
func isEmptyList(l list) bool {
	if l.head != nil {
		return false
	}
	return true
}

func addFirst(data int, list1 *list) {
	newNode := addNode(data)
	if isEmptyList(*list1) {
		*&list1.head = &newNode
		*&list1.tail = &newNode
	} else {
		newNode.next = *&list1.head
		*&list1.head = &newNode
	}
}
func addLast(data int, list1 *list) {
	newNode := addNode(data)
	if isEmptyList(*list1) {
		*&list1.head = &newNode
		*&list1.tail = &newNode
	} else {
		*&list1.tail.next = &newNode
		*&list1.tail = &newNode
	}
}
func removeFirst(list1 *list) {
	if isEmptyList(*list1) {
		fmt.Println("List is Empty")
	} else if *&list1.head.next == nil {
		*&list1.head = nil
		*&list1.tail = nil
	} else {
		*&list1.head = *&list1.head.next
	}
}
func removeLast(list1 *list) {
	if isEmptyList(*list1) {
		fmt.Println("List is Empty")
	} else if *&list1.head == *&list1.tail {
		*&list1.head = nil
		*&list1.tail = nil
	} else {
		nodePtr := *&list1.head
		for nodePtr.next != *&list1.tail {
			nodePtr = nodePtr.next
		}
		*&list1.tail = nodePtr
		*&list1.tail.next = nil
	}
}
func traverse(list1 *list) {
	if isEmptyList(*list1) {
		fmt.Println("List is Empty")
	} else {
		nodePtr := *&list1.head
		fmt.Printf("%v ", nodePtr.data)
		for nodePtr.next != *&list1.tail {
			fmt.Printf("%v ", nodePtr.next.data)
			nodePtr = nodePtr.next
		}
		fmt.Printf("%v ", nodePtr.next.data)
		fmt.Println()
	}
}
func count(list1 *list) int {
	count := 0
	if isEmptyList(*list1) {
		fmt.Println("List is Empty")
	} else {
		nodePtr := *&list1.head
		for nodePtr.next != *&list1.tail {
			nodePtr = nodePtr.next
			count++
		}
	}
	return count
}

func addAtPosition(position int, data int, list1 *list) {
	newNode := addNode(data)
	if *&list1.head == *&list1.tail {
		addFirst(data, list1)
	} else if position > count(list1) {
		addLast(data, list1)
	} else {
		nodePtr := *&list1.head
		for count := 1; count < position; count++ {
			nodePtr = nodePtr.next
		}
		newNode.next = nodePtr.next
		nodePtr.next = &newNode
	}

}

func removeFromPosition(position int, list1 *list) {
	if *&list1.head == *&list1.tail {
		fmt.Println("List is Empty")
	} else if position > count(list1) {
		removeLast(list1)
	} else if 1 == count(list1) {
		removeFirst(list1)
	} else {
		nodePtr := *&list1.head
		for count := 1; count < position-1; count++ {
			nodePtr = nodePtr.next
		}

		nodePtr.next = nodePtr.next.next
	}
}

func main() {
	list1 := list{
		head: nil,
		tail: nil,
	}
	addLast(10, &list1)
	addLast(20, &list1)
	addLast(30, &list1)
	addLast(40, &list1)
	addLast(50, &list1)
	addLast(60, &list1)
	addLast(70, &list1)
	addLast(80, &list1)
	addLast(90, &list1)
	addLast(100, &list1)
	fmt.Println("Original List")
	traverse(&list1)
	fmt.Println("After add First")
	addFirst(5, &list1)
	traverse(&list1)
	fmt.Println("After Add Last")
	addLast(105, &list1)
	traverse(&list1)
	fmt.Println("After at position")
	addAtPosition(3, 25, &list1)
	traverse(&list1)
	fmt.Println("After Remove first")
	removeFirst(&list1)
	traverse(&list1)
	fmt.Println("After Remove last")
	removeLast(&list1)
	traverse(&list1)
	fmt.Println("After Remove from position")
	removeFromPosition(3, &list1)
	traverse(&list1)
}
