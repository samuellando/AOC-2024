package priorityqueue

import (
	"container/heap"
)

// Item represents a single element in the priority queue
type Item[T any] struct {
	Value    T // The value of the item; arbitrary data
	Priority int    // The priority of the item; lower value means higher priority
	Index    int    // The index of the item in the heap; needed for heap.Interface
}

// PriorityQueue implements heap.Interface and holds Items
type PriorityQueue[T any] []*Item[T]

func New[T any]() PriorityQueue[T] {
    pq := make(PriorityQueue[T], 0)
    heap.Init(&pq)
    return pq
}

func (pq PriorityQueue[T]) Len() int { return len(pq) }

func (pq PriorityQueue[T]) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].Priority < pq[j].Priority
}

func (pq PriorityQueue[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue[T]) Push(x any) {
	n := len(*pq)
	item := x.(*Item[T])
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue[T]) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // don't stop the GC from reclaiming the item eventually
	item.Index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue[T]) Update(item *Item[T], value T, priority int) {
	item.Value = value
	item.Priority = priority
	heap.Fix(pq, item.Index)
}
