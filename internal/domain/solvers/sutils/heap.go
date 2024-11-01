package sutils

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

// innerHeap реализует интерфейс heap.Interface, описанный в container/heap.
type innerHeap []*Item

func (h innerHeap) Len() int { return len(h) }

func (h innerHeap) Less(i, j int) bool {
	return h[i].Weight < h[j].Weight
}

func (h innerHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *innerHeap) Push(x any) {
	*h = append(*h, x.(*Item))
}

func (h *innerHeap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*h = old[0 : n-1]

	return item
}
