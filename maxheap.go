package main

type HeapElement struct {
	index int
	value string
}
type maxheap struct {
	heapArray []HeapElement
	values    map[string]int
	size      int
	maxsize   int
}

func (m *maxheap) leaf(index int) bool {
	if index >= (m.size/2) && index <= m.size {
		return true
	}
	return false
}

func (m *maxheap) parent(index int) int {
	return (index - 1) / 2
}

func (m *maxheap) leftchild(index int) int {
	return 2*index + 1
}

func (m *maxheap) rightchild(index int) int {
	return 2*index + 2
}

func (m *maxheap) insert(item *HeapElement) {
	if _, ok := m.values[item.value]; ok {
		for i := 0; i < len(m.heapArray); i++ {
			if m.heapArray[i].value == item.value {
				m.heapArray[i].index = item.index
				if m.size == 5 {
					var ha = make([]HeapElement, m.size)
					ha[0] = m.heapArray[i]
					for j := 0; j < len(m.heapArray); j++ {
						if j != i {
							if j < i {
								ha[j+1] = m.heapArray[j]
							} else {
								ha[j] = m.heapArray[j]
							}
						}
					}
					m.heapArray = ha
					//m.heapArray[0], m.heapArray[1], m.heapArray[2], m.heapArray[3], m.heapArray[4] = ha[1], ha[2], ha[3], ha[4], ha[0]
					i = 0
					for m.heapArray[i].index < m.heapArray[i+1].index {
						m.swap(i, i+1)
						i++
					}
				} else {
					m.upHeapify(i)
				}
				return
			}
		}
	}
	if m.size == 5 {
		if m.heapArray[len(m.heapArray)-1].index < item.index {
			delete(m.values, m.heapArray[len(m.heapArray)-1].value)
			m.heapArray[len(m.heapArray)-1] = *item
			m.values[item.value] = 1
			m.upHeapify(0)
		} else {
			return
		}
	} else {

		m.heapArray = append(m.heapArray, *item)
		m.size++
		m.values[item.value] = 1
		m.upHeapify(0)
	}
}

func (m *maxheap) swap(first, second int) {
	m.heapArray[first], m.heapArray[second] = m.heapArray[second], m.heapArray[first]
}

func (m *maxheap) upHeapify(index int) {
	for m.heapArray[index].index > m.heapArray[m.parent(index)].index {
		m.swap(index, m.parent(index))
		index = m.parent(index)
	}
}

func (m *maxheap) downHeapify(current int) {
	if m.leaf(current) {
		return
	}
	largest := current
	leftChildIndex := m.leftchild(current)
	rightRightIndex := m.rightchild(current)
	//If current is smallest then return
	if leftChildIndex < m.size && m.heapArray[leftChildIndex].index > m.heapArray[largest].index {
		largest = leftChildIndex
	}
	if rightRightIndex < m.size && m.heapArray[rightRightIndex].index > m.heapArray[largest].index {
		largest = rightRightIndex
	}
	if largest != current {
		m.swap(current, largest)
		m.downHeapify(largest)
	}
	return
}

func (m *maxheap) buildMaxHeap() {
	for index := (m.size / 2) - 1; index >= 0; index-- {
		m.downHeapify(index)
	}
}

func (m *maxheap) remove() HeapElement {
	top := m.heapArray[0]
	delete(m.values, top.value)
	m.heapArray[0] = m.heapArray[m.size-1]
	m.heapArray = m.heapArray[:(m.size)-1]
	m.size--
	m.downHeapify(0)
	return top
}
