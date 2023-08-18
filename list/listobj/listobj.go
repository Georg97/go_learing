package listobj

import "fmt"

type List[T comparable] struct {
	next *List[T]
	val  T
}

func NewList[T comparable](val T) (*List[T], error) {
    list := &List[T]{nil, val}
    return list, nil
}

func (list *List[T]) Add(elem T) {
    if list.next != nil {
        list.next.Add(elem)
        return
    }
    newElem := &List[T]{nil, elem}
    list.next = newElem
}

func (list *List[T]) GetIndexOf(elem T) int {
    return list.getIndexOf(elem, 0)
}
func (list *List[T]) getIndexOf(elem T, startIndex int) int {
    if list.val == elem {
        return startIndex
    }
    if list.next == nil {
        return -1
    }
    return list.next.getIndexOf(elem, startIndex + 1)
}
func (list *List[T]) Get(index int) T {
    return list.get(index, 0)
}
func (list *List[T]) get(index, start int) T {
    if start == index {
        return list.val
    }
    return list.next.get(index, start + 1)
}

func (list *List[T]) String() string {
    return fmt.Sprintf("%v -> %v", list.val, list.next)
}


