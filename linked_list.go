package data_structures

import "fmt"

type node struct {
	value any
	next  *node
}

type linkedList struct {
	head *node
	len  int
}

func NewLinkedList() *linkedList {
	return &linkedList{}
}

func (ll *linkedList) Set(vals ...any) *linkedList {

	if len(vals) <= 0 {
		ll.head = nil
		return ll
	}
	ll.head = &node{
		value: vals[0],
		next:  nil,
	}
	ll.len = 1

	currentNode := ll.head
	for _, v := range vals[1:] {
		currentNode.next = &node{
			value: v,
			next:  nil,
		}
		ll.len++
		currentNode = currentNode.next
	}

	return ll
}

func (ll *linkedList) Length() int {
	return ll.len
}

func (ll *linkedList) Append(val any) *linkedList {
	if ll.len == 0 {
		ll.head = &node{value: val}
	} else {
		currentNode := ll.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = &node{value: val}
	}
	ll.len++
	return ll
}

func (ll *linkedList) Prepend(val any) *linkedList {
	ll.head = &node{
		value: val,
		next:  ll.head,
	}
	ll.len++
	return ll
}

func (ll *linkedList) Remove(val any) *linkedList {
	if ll.head.value == val {
		ll.head = ll.head.next
		ll.len--
		return ll
	}

	bNode := ll.head
	cNode := ll.head.next
	for cNode != nil {
		if cNode.value == val {
			bNode.next = cNode.next
			ll.len--
			return ll
		}
		bNode = cNode
		cNode = cNode.next
	}
	return ll
}

func (ll *linkedList) Find(val any) *node {
	if ll.Length() <= 0 {
		return nil
	}
	currentNode := ll.head
	for currentNode.next != nil {
		if currentNode.value == val {
			return currentNode
		}
		currentNode = currentNode.next
	}
	if currentNode.value == val {
		return currentNode
	}
	return nil
}

func (ll *linkedList) Index(val any) (int, error) {
	if ll.Length() <= 0 {
		return -1, fmt.Errorf("value not found")
	}
	index := 0
	currentNode := ll.head
	for currentNode.next != nil {
		if currentNode.value == val {
			return index, nil
		}
		index++
		currentNode = currentNode.next
	}
	if currentNode.value == val {
		return index, nil
	}
	return -1, fmt.Errorf("value not found")
}

func (ll *linkedList) Print() {
	if ll.len == 0 {
		return
	}
	fmt.Printf("[")
	currentNode := ll.head
	for currentNode.next != nil {
		fmt.Printf("%v ", currentNode.value)
		currentNode = currentNode.next
	}
	fmt.Printf("%v ]", currentNode.value)
}
