
type Node struct {
	val  int32 // Downsized from int (8 bytes) to int32 (4 bytes)
	next int32 // Downsized from *Node (8 bytes pointer) to int32 index (4 bytes)
}

// LinkedList encapsulates an arena allocation strategy, avoiding individual heap allocations.
type LinkedList struct {
	nodes []Node // Single contiguous block of memory (Arena)
	head  int32  // -1 signifies nil
	tail  int32  // -1 signifies nil
	free  int32  // Head of the internal free-list for reusing deleted nodes
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		nodes: make([]Node, 0, 8), // Pre-allocating small capacity avoids aggressive allocations
		head:  -1,
		tail:  -1,
		free:  -1,
	}
}

// allocNode handles O(1) node recycling or allocation out of the backing slice.
func (ll *LinkedList) allocNode(val int) int32 {
	v32 := int32(val)
	if ll.free != -1 {
		// Recycle a previously deleted index
		index := ll.free
		ll.free = ll.nodes[index].next
		ll.nodes[index] = Node{val: v32, next: -1}
		return index
	}
	// Grow slice sequentially
	ll.nodes = append(ll.nodes, Node{val: v32, next: -1})
	return int32(len(ll.nodes) - 1)
}

func (ll *LinkedList) Get(index int) int {
	if index < 0 || ll.head == -1 {
		return -1
	}
	currIdx := ll.head
	for i := 0; i < index; i++ {
		currIdx = ll.nodes[currIdx].next
		if currIdx == -1 {
			return -1
		}
	}
	return int(ll.nodes[currIdx].val)
}

func (ll *LinkedList) InsertTail(val int) {
	nodeIdx := ll.allocNode(val)
	if ll.head == -1 {
		ll.head = nodeIdx
		ll.tail = nodeIdx
		return
	}
	ll.nodes[ll.tail].next = nodeIdx
	ll.tail = nodeIdx
}

func (ll *LinkedList) InsertHead(val int) {
	nodeIdx := ll.allocNode(val)
	ll.nodes[nodeIdx].next = ll.head
	ll.head = nodeIdx
	if ll.tail == -1 {
		ll.tail = nodeIdx
	}
}

func (ll *LinkedList) Remove(index int) bool {
	if ll.head == -1 || index < 0 {
		return false
	}

	if index == 0 {
		rmIdx := ll.head
		if ll.head == ll.tail {
			ll.tail = -1
		}
		ll.head = ll.nodes[ll.head].next
		
		// Recycle memory into free-list
		ll.nodes[rmIdx].next = ll.free
		ll.free = rmIdx
		return true
	}

	currIdx := ll.head
	for i := 0; i < index-1; i++ {
		currIdx = ll.nodes[currIdx].next
		if currIdx == -1 || ll.nodes[currIdx].next == -1 {
			return false
		}
	}

	rmIdx := ll.nodes[currIdx].next
	if rmIdx == -1 {
		return false
	}

	if rmIdx == ll.tail {
		ll.tail = currIdx
	}

	ll.nodes[currIdx].next = ll.nodes[rmIdx].next

	// Recycle memory into free-list
	ll.nodes[rmIdx].next = ll.free
	ll.free = rmIdx
	return true
}

func (ll *LinkedList) GetValues() []int {
	if ll.head == -1 {
		return nil
	}
	
	// Count elements first to allocate a single, perfectly sized slice
	count := 0
	currIdx := ll.head
	for currIdx != -1 {
		count++
		currIdx = ll.nodes[currIdx].next
	}

	val := make([]int, count)
	currIdx = ll.head
	for i := 0; i < count; i++ {
		val[i] = int(ll.nodes[currIdx].val)
		currIdx = ll.nodes[currIdx].next
	}
	return val
}
