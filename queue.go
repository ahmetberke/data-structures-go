package data_structures

type queue struct {
	front  *node
	rear   *node
	length int
}

func NewQueue() *queue {
	return &queue{}
}

func (q *queue) Enqueue(value any) *queue {
	if q.front == nil {
		q.front = &node{
			value: value,
		}
		q.rear = q.front
	} else {
		q.front.next = &node{
			value: value,
		}
	}
	q.length++
	return q
}

func (q *queue) Dequeue() *queue {
	if q.front == nil {
		return q
	}
	if q.front.next == nil {
		q.front = nil
		q.rear = nil
	}
	q.front = q.front.next
	q.length--
	return q
}

func (q *queue) Peek() any {
	return q.front.value
}

func (q *queue) Length() int {
	return q.length
}
