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
