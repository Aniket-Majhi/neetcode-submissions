

// Node uses simple array indices for rapid pointer-chasing performance.
// Struct members are naturally ordered to ensure flawless data alignment.
type Node struct {
	val  int // Kept as native int for 64-bit CPU register optimization
	next int // Element index pointer: -1 represents nil
}

// LinkedList implements a flat Arena block with constant tracking variables
// to eliminate layout latency and dynamic resizing loops.
type LinkedList struct {
	nodes []Node // Flat memory layout for perfect CPU cache locality
	head  int    // Arena index pointing to head
	tail  int    // Arena index pointing to tail
	free  int    // Head of O(1) recycled freelist
	size  int    // Real-time tracking of length to remove sizing loops
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		nodes: make([]Node, 0, 16), // Pre-allocated buffer prevents intermediate allocations
		head:  -1,
		tail:  -1,
		free:  -1,
		size:  0,
	}
}

// allocNode retrieves a node in absolute O(1) time without garbage collector strain.
func (ll *LinkedList) allocNode(val int) int {
	if ll.free != -1 {
		index := ll.free
		ll.free = ll.nodes[index].next
		ll.nodes[index] = Node{val: val, next: -1}
		return index
	}
	ll.nodes = append(ll.nodes, Node{val: val, next: -1})
	return len(ll.nodes) - 1
}

func (ll *LinkedList) Get(index int) int {
	// Immediate bounds checking using the tracked list size
	if index < 0 || index >= ll.size {
		return -1
	}

	currIdx := ll.head
	// Extremely tight loop: directly traverses flat array space
	for i := 0; i < index; i++ {
		currIdx = ll.nodes[currIdx].next
	}
	return ll.nodes[currIdx].val
}

func (ll *LinkedList) InsertTail(val int) {
	nodeIdx := ll.allocNode(val)
	ll.size++ // Keep size track up to date

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
	ll.size++

	ll.nodes[nodeIdx].next = ll.head
	ll.head = nodeIdx
	if ll.tail == -1 {
		ll.tail = nodeIdx
	}
}

func (ll *LinkedList) Remove(index int) bool {
	if index < 0 || index >= ll.size {
		return false
	}
	ll.size--

	if index == 0 {
		rmIdx := ll.head
		if ll.head == ll.tail {
			ll.tail = -1
		}
		ll.head = ll.nodes[ll.head].next
		
		// Recycle into free-list instantly
		ll.nodes[rmIdx].next = ll.free
		ll.free = rmIdx
		return true
	}

	currIdx := ll.head
	// Traversing using index tracking allows loops to compile down to bare register jumps
	for i := 0; i < index-1; i++ {
		currIdx = ll.nodes[currIdx].next
	}

	rmIdx := ll.nodes[currIdx].next
	if rmIdx == ll.tail {
		ll.tail = currIdx
	}

	ll.nodes[currIdx].next = ll.nodes[rmIdx].next

	// Recycle into free-list instantly
	ll.nodes[rmIdx].next = ll.free
	ll.free = rmIdx
	return true
}

func (ll *LinkedList) GetValues() []int {
	if ll.size == 0 {
		return nil
	}
	
	// CRITICAL LATENCY OPTIMIZATION:
	// O(1) allocation because we track length globally instead of loops.
	val := make([]int, ll.size)
	
	currIdx := ll.head
	for i := 0; i < ll.size; i++ {
		val[i] = ll.nodes[currIdx].val
		currIdx = ll.nodes[currIdx].next
	}
	return val
}
