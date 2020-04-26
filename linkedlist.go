package main;

import "fmt"

type Node struct {
	data int
	next *Node
}

type Linkedlist struct {
    head *Node
}

func Inserthead(l *Linkedlist, d int){
    var newnode Node
    newnode.data = d
	if l.head == nil {
		l.head = &newnode
	} else {
        newnode.next = l.head
        l.head = &newnode
    }
}

func InsertEnd(l *Linkedlist, d int){
    var newnode Node
    temp := l.head
    for temp.next != nil {
        temp = temp.next
    }
    newnode.data = d
    temp.next = &newnode
}

func Remove(l *Linkedlist, d int){
    temp := l.head
    var prev *Node
    for temp.data != d{
        prev = temp
        temp = temp.next
    }
    if temp.next == nil {
        fmt.Println("Element not present")
    } else {
        prev.next = temp.next
    }
}

func (l Linkedlist) print(){
    temp := l.head
    for temp.next != nil {
        fmt.Println(temp.data)
        temp = temp.next
    }
    fmt.Println(temp.data)
}
func main() {
    var list Linkedlist
    Inserthead(&list, 4)
    Inserthead(&list, 6)
    Inserthead(&list, 8)
    list.print()
    InsertEnd(&list, 10)
    list.print()
    Remove(&list, 6)
    list.print()
}
