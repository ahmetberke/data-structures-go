package data_structures

import "testing"

func TestNewQueue(t *testing.T) {
	_ = NewQueue()
}

func TestQueue_Peek(t *testing.T) {
	q := NewQueue().Enqueue(1)
	expected := 1
	actual := q.Peek()
	if expected != actual {
		t.Errorf("peek node must be %v but it's %v", expected, actual)
	}
}

func TestQueue_Length(t *testing.T) {
	q := NewQueue().Enqueue(1).Enqueue(3).Enqueue("merhaba").Dequeue()
	expected := 2
	actual := q.Length()
	if expected != actual {
		t.Errorf("queue length must be %v but it's %v", expected, actual)
	}
}

func TestQueue_Enqueue(t *testing.T) {
	q := NewQueue().Enqueue(1)
	expected := 1
	actual := q.Peek()
	if expected != actual {
		t.Errorf("front node of queue must be %v but it's %v", expected, actual)
	}
}

func TestQueue_Dequeue(t *testing.T) {
	q := NewQueue().Enqueue(1).Enqueue(2).Dequeue()
	expected := 2
	actual := q.Peek()
	if expected != actual {
		t.Errorf("front node of queue must be %v but it's %v", expected, actual)
	}
}
