package linkedlist

import "fmt"

type DoubleLikedList struct {
	size int
	head *Node
	tail *Node
}

type Node struct {
	value      int
	prev, next *Node
}

// Clears the whole list
func (d *DoubleLikedList) Clear() {
	trav := d.head
	for trav != nil {
		next := trav.next
		trav.prev, trav.next = nil, nil
		trav = next
	}
	d.head, d.tail, d.size = nil, nil, 0
}

func (d *DoubleLikedList) AddFirst(value int) {
	if d.size == 0 {
		new := &Node{value: value}
		d.head = new
		d.tail = new
	} else {
		d.head.prev = &Node{value, nil, d.head}
		d.head = d.head.prev
	}

	d.size++
}

func (d *DoubleLikedList) AddLast(value int) {
	if d.size == 0 {
		new := &Node{value: value}
		d.head = new
		d.tail = new
	} else {
		d.tail.next = &Node{value: value, prev: d.tail}
		d.tail = d.tail.next
	}

	d.size++
}

func (d *DoubleLikedList) PeekFirst() (int, error) {
	if d.size == 0 {
		return 0, fmt.Errorf("empty list")
	}
	return d.head.value, nil
}

func (d *DoubleLikedList) PeekLast() (int, error) {
	if d.size == 0 {
		return 0, fmt.Errorf("empty list")
	}
	return d.tail.value, nil
}

func (d *DoubleLikedList) RemoveFirst() (int, error) {
	if d.size == 0 {
		return 0, fmt.Errorf("empty list")
	}
	data := d.head.value
	d.head = d.head.next
	d.size--
	if d.size == 0 {
		d.tail = nil
	} else {
		d.head.prev = nil
	}
	return data, nil
}

func (d *DoubleLikedList) RemoveLast() (int, error) {
	if d.size == 0 {
		return 0, fmt.Errorf("empty list")
	}
	data := d.tail.value
	d.tail = d.head.prev
	d.size--
	if d.size == 0 {
		d.head = nil
	} else {
		d.tail.next = nil
	}
	return data, nil
}

func (d *DoubleLikedList) remove(node *Node) (int, error) {
	if node.prev == nil {
		return d.RemoveFirst()
	}
	if node.next == nil {
		return d.RemoveLast()
	}
	node.next.prev = node.prev
	node.prev.next = node.next
	d.size--
	return node.value, nil
}

func (d *DoubleLikedList) RemoveAt(index int) (int, error) {
	if index < 0 || index >= d.size {
		return 0, fmt.Errorf("index value %d is invalid", index)
	}
	var i int
	var trav *Node
	if index < d.size/2 {
		for i = 0; i != index; i++ {
			trav = trav.next
		}
	} else {
		for i = 0; i != index; i-- {
			trav = trav.prev
		}
	}
	return d.remove(trav)
}

func (d *DoubleLikedList) IndexOf(value int) []int {
	// Todo
	return nil
}
