package data_structures

import (
	"testing"
)

func TestNewLinkedList(t *testing.T) {
	_ = NewLinkedList()
}

func TestLinkedList_Set(t *testing.T) {
	ll := NewLinkedList().Set(12, "selam", 15)
	expecteds := []any{12, "selam", 15}
	situation := 0
	currentNode := ll.head
	for currentNode.next != nil {
		output := currentNode.value
		if expecteds[situation] != output {
			t.Errorf("value is %v, must be %v", output, expecteds[situation])
		}
		currentNode = currentNode.next
		situation++
	}
}

func TestLinkedList_Append(t *testing.T) {
	ll := NewLinkedList().Append(12).Append("hello")
	expecteds := []any{12, "hello"}
	situation := 0
	currentNode := ll.head
	for currentNode.next != nil {
		output := currentNode.value
		if expecteds[situation] != output {
			t.Errorf("value is %v, must be %v", output, expecteds[situation])
		}
		currentNode = currentNode.next
		situation++
	}
}

func TestLinkedList_Prepend(t *testing.T) {
	ll := NewLinkedList().Prepend(12).Prepend("hello")
	expecteds := []any{"hello", 12}
	situation := 0
	currentNode := ll.head
	for currentNode.next != nil {
		output := currentNode.value
		if expecteds[situation] != output {
			t.Errorf("value is %v, must be %v", output, expecteds[situation])
		}
		currentNode = currentNode.next
		situation++
	}
}

func TestLinkedList_Find(t *testing.T) {
	ll := NewLinkedList().Set(31).Prepend(12).Append("hello")
	n := ll.Find("hello")
	if n != ll.head.next.next {
		t.Errorf("didn't match")
	}
}

func TestLinkedList_Index(t *testing.T) {
	ll := NewLinkedList().Set(31).Prepend(12).Append("hello")
	n, err := ll.Index("hello")
	if err != nil {
		t.Errorf("didn't match")
	}
	if n != 2 {
		t.Errorf("didn't match")
	}
}

func TestLinkedList_Length(t *testing.T) {
	ll := NewLinkedList().Set(31).Prepend(12).Append("hello")
	expected := 3
	expected2 := ll.len
	output := ll.Length()
	if output != expected || output != expected2 {
		t.Errorf("value is %v, must be %v and %v", output, expected, expected2)
	}
}

func TestNode_Index(t *testing.T) {
	ll := NewLinkedList().Set(31).Prepend(12).Append("hello")
	n := ll.Find(31)
	expected := 1
	output := n.Index()
	if expected != output {
		t.Errorf("value is %v, must be %v", output, expected)
	}
}

func TestLinkedList_Remove(t *testing.T) {
	ll := NewLinkedList().Set(31).Prepend(12).Append("hello").Remove(12)
	expected := "hello"
	output := ll.head.next.value
	if expected != output {
		t.Errorf("value is %v, must be %v", output, expected)
	}
}
