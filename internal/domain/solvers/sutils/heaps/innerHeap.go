package heaps

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
