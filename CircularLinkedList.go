package main;

import "fmt"

type Node struct {
    data int
    next *Node
}

type Circulatlist struct {
    head *Node
    tail *Node
}

func create(C *Circulatlist, d int) {
    var newnode Node
    newnode.data = d
    C.head = &newnode
    C.tail = &newnode
    newnode.next = &newnode
}

func size(C *Circulatlist) int {
    pos := 0
    head := C.head
    for head.next != C.head {
        head = head.next
        pos++
    }
    return pos
}

func (C *Circulatlist) insert(d,p int) {
    var newnode Node
    newnode.data = d
    if p == 0 {
        newnode.next = (*C).head
        (*C).head = &newnode
        (*C).tail.next = &newnode
    } else if p == size(C)+1 {
        newnode.next = (*C).head
        (*C).tail.next = &newnode
        (*C).tail = &newnode
    } else {
        pos := 0
        current := (*C).head
        prev := (*C).head
        for pos != p {
            prev = current
            current = current.next
            pos++
        }
        prev.next = &newnode
        newnode.next = current
    }
}

func (C *Circulatlist) print() {
    start := C.head
    for start.next != C.head {
        fmt.Println(start.data)
        start = start.next
    }
    fmt.Println(start.data)
}

func main(){
    var circular_list Circulatlist
    create(&circular_list,9)
    (&circular_list).insert(10,0)
    (&circular_list).insert(11,1)
    (&circular_list).insert(12,1)
    (&circular_list).insert(13,1)
    (&circular_list).insert(14,55555)
    circular_list.print()
}
