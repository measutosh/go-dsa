# go-dsa

This repo will be theplaceholder for all the DSA questions I will practice using Golang. I have tested all the answers using go playground.

**1. Insert node at the beginning of a linked list.**

**Write a program to create and add new node at the beginning of linked list.**

```Go
package main

import "fmt"

type LinkNode struct {
	data int
	next *LinkNode
}

func getLinkNode(data int) *LinkNode {
	var me *LinkNode = &LinkNode{}
	me.data = data
	me.next = nil
	return me
}

type SingleLL struct {
	head *LinkNode
}

func getSingleLL() *SingleLL {
	var me *SingleLL = &SingleLL{}
	me.head = nil
	return me
}

// add new node at the beginning of the linkedlist
func (this *SingleLL) addNode(data int) {
	// create a new node
	var node *LinkNode = getLinkNode(data)
	// connect the current node to the previous node
	node.next = this.head
	this.head = node
}

// display all elements of the linked list
func (this SingleLL) display() {
	if this.head == nil {
		return
	}

	var temp *LinkNode = this.head
	for temp != nil {
		fmt.Println(temp.data)
		// jumping to next node
		temp = temp.next
	}
	fmt.Print("NULL\n")
}

func main() {
	var sll *SingleLL = getSingleLL()
	sll.addNode(8)
	sll.addNode(9)
	sll.addNode(10)
	sll.display()
}

/*

OUTPUT:

10
9
8
NULL

*/
```



**2. Insert node at end of linked list**

**Write a program to add a new node at the tail of a linked list.**

```go
package main

import (
	"fmt"
)

// linkedlist node declaration
type LinkNode struct {
  data int
  next * LinkNode
}

// constructor for linkedlist node
func getLinkNode(data int) * LinkNode {
  return &LinkNode {data, nil}
}

type SingleLL  struct {
  head * LinkNode
}

func getSingleLL() * SingleLL {
  return &SingleLL {nil}
}

// adding new node at the end of the linkedlist
func(this *SingleLL) addNode (value int) {
  var node * LinkNode = getLinkNode(value)
  if this.head == nil {
    this.head = node
  } else {
    var temp * LinkNode = this.head
    // loop till you get the last node
    for (temp.next != nil) {
      temp = temp.next
    }
    temp.next = node
  }
}


// dsiplay all values
func (this SingleLL) display() {
  if this.head != nil {
    var temp * LinkNode = this.head
    // loop till the last node
    for (temp != nil) {
      // display the current node value
      fmt.Print(" ", temp.data)
      // jump to next node
      temp = temp.next
    } 
  } else {
      fmt.Print("Empty linked list \n")
  }
  
}


func main() {
	var sll * SingleLL = getSingleLL()
  sll.addNode(1)
  sll.addNode(2)
  sll.addNode(3)
  sll.addNode(4)
  sll.addNode(5)

  fmt.Print("Linked List \n")
  sll.display()
  
}

```
