package queue

type Queue[T any] interface {
	Enqueue(T)
	Dequeue() T
	Length() int
}

type ringQueue[T any] struct {
	values []T
	front  int
	back   int
	length int
}

func New[T any]() Queue[T] {
	values := make([]T, 1000)
	return &ringQueue[T]{values, 0, 0, 0}
}

func (q *ringQueue[T]) Length() int {
	return q.length
}

func (q *ringQueue[T]) Enqueue(v T) {
	if q.length < len(q.values) {
		q.values[q.back] = v
		q.back = q.nextIndex(q.back)
		q.length++
	} else {
		q.increaseCapacity()
		q.Enqueue(v)
	}
}

func (q *ringQueue[T]) nextIndex(i int) int {
	n := i + 1
	if i == len(q.values) {
		n = 0
	}
	return n
}

func (q *ringQueue[T]) increaseCapacity() {
	new_values := make([]T, len(q.values)*2)
	for i := 0; q.Length() > 0; i++ {
		new_values[i] = q.Dequeue()
	}
	q.values = new_values
	q.front = 0
	q.back = q.length
}

func (q *ringQueue[T]) Dequeue() T {
	if q.length > 0 {
		v := q.values[q.front]
		q.front++
		q.length--
		return v
	} else {
		var n T
		return n
	}
}
