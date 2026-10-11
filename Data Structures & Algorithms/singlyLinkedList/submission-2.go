type Node struct {
	val  int
	Next *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		Head: nil,
		Tail: nil,
	}
}

func (ll *LinkedList) Get(index int) int {
	if index < 0 || ll.Head == nil {
		return -1
	}
	current := ll.Head // Fixed: added ':' for short declaration
	
	// If index is 0, we don't need to traverse, just return Head's value
	for i := 0; i < index; i++ { // Fixed: changed condition from '<=' to '<'
		if current.Next == nil {
			return -1
		}
		current = current.Next
	}
	return current.val
}

func (ll *LinkedList) InsertTail(val int) {
	node := &Node{val: val}
	
	// Edge Case: If list is empty, both Head and Tail point to the new node
	if ll.Head == nil {
		ll.Head = node
		ll.Tail = node
		return
	}
	
	ll.Tail.Next = node // Link old tail to new node
	ll.Tail = node      // Fixed: Actually update the Tail pointer of the list
}

func (ll *LinkedList) InsertHead(val int) {
	node := &Node{val: val}
	node.Next = ll.Head
	ll.Head = node
	
	// Edge Case: If it was an empty list, the new node is also the Tail
	if ll.Tail == nil {
		ll.Tail = node
	}
}

func (ll *LinkedList) Remove(index int) bool {
	if ll.Head == nil || index < 0 {
		return false
	}

	// Edge Case: Removing the Head node (index 0)
	if index == 0 {
		if ll.Head == ll.Tail { // If it was the only node left
			ll.Tail = nil
		}
		ll.Head = ll.Head.Next
		return true
	}

	current := ll.Head // Fixed: added ':'
	
	// We need to stop at the node RIGHT BEFORE the one we want to delete
	for i := 0; i < index-1; i++ { 
		if current.Next == nil {
			return false
		}
		current = current.Next
	}

	// Make sure the target node to delete actually exists
	if current.Next == nil {
		return false
	}

	rm := current.Next // Fixed: added ':'
	
	// Edge Case: If we are removing the very last node, update Tail
	if rm == ll.Tail {
		ll.Tail = current
	}

	current.Next = rm.Next
	return true
}

func (ll *LinkedList) GetValues() []int {
	var val []int // Instantiates an empty slice
	if ll.Head == nil {
		return val
	}

	c := ll.Head // Fixed: added ':'
	for {
		val = append(val, c.val) // Fixed: dynamically append values to prevent index panics
		
		if c == ll.Tail { // Correctly break when we reach the end
			break
		}
		c = c.Next
	}
	return val
}
