package ring

type Node struct {
	next, prev *Node
	value      interface{} //what is this interface ?
}

// Ring : Presents a circular doubly linked list
type Ring struct {
	length  int
	current *Node
}

func (this *Ring) Len() int {
	return this.length
}

func (this *Ring) Next() {
	this.current = this.current.next
}

func (this *Ring) Prev() {
	this.current = this.current.prev
}

func (this *Ring) SetCurrent(value interface{}) {
	this.current.value = value
}

func (this *Ring) Insert(value interface{}) {
	this.length++
	node := &Node{
		next:  this.current.next,
		prev:  this.current,
		value: value,
	}
	this.current.next.prev = node
	this.current.next = node
}

//Delete the current node
func (this *Ring) Delete() {
	this.length--
	this.current.prev.next = this.current.next
	this.current.next.prev = this.current.prev
	this.current = this.current.prev
}

// Finds a nearby node with specified value and makes it current
// If the value did not find, current node remains unchanged
func (this *Ring) Find(value interface{}) {
	// Iterate over nodes until returned to the current node
	for node := this.current; ; node = node.next {
		// Make node current if value coincide with param
		if node.value == value {
			this.current = node
			break
		}
		// Check for full cycle
		if node == this.current.prev {
			break
		}
	}
}

// Ring constructor
// Returns a new one-node ring
func NewRing() Ring {
	ring := Ring{}
	ring.current = new(Node)
	ring.current.next = ring.current
	ring.current.prev = ring.current
	return ring
}
