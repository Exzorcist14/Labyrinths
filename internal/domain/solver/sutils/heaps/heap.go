package heaps

import (
	"container/heap"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/maze/cells"
)

// Heap - куча минимумов.
type Heap struct {
	heap innerHeap // Heap - обёртка innerHeap, абстрагируюющая от понимания применения container/heap
}

// Item содержит координаты клетки и её тип.
type Item struct {
	Vertex cells.Coordinates
	Weight cells.Type
}

// New возвращает инициализированный Heap.
func New() Heap {
	h := Heap{
		heap: innerHeap{},
	}

	heap.Init(&h.heap)

	return h
}

// Len возвращает количество элементов в куче.
func (h *Heap) Len() int {
	return len(h.heap)
}

// Push добавляет item в кучу.
func (h *Heap) Push(item Item) {
	heap.Push(&h.heap, &item)
}

// Pop возвращает Item с наименьшим Item.Weight в куче, удаляя из неё.
func (h *Heap) Pop() *Item {
	return heap.Pop(&h.heap).(*Item)
}
