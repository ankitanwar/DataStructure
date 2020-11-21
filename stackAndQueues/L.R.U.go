package stackandqueues

import "github.com/ankitanwar/Golang-DataStructure/linkedlist"

//Dict : Dictonary of LRu
var Dict = make(map[string]*linkedlist.Dnode)

//LRU : To implement the LRU
type LRU struct {
	Head  *linkedlist.Dnode
	Tail  *linkedlist.Dnode
	Count int
}

//Insert : To insert the value into the LRU
func (l *LRU) Insert(value string) {
	n := &linkedlist.Dnode{}
	n.Value = value
	if len(Dict) < l.Count {
		if l.Head == nil {
			Dict[value] = l.Head
			l.Head = n
			l.Tail = n
			return
		}
		l.Tail.Next = n
		n.Prev = l.Tail
		Dict[value] = n
		l.Tail = l.Tail.Next
	} else {
		current := l.Tail.Value
		delete(Dict, current)
		prev := l.Tail.Prev
		prev.Next = nil
		l.Tail.Prev = nil
		l.Tail = prev
		l.Head.Prev = n
		n.Next = l.Head
		l.Head = n
		Dict[value] = n

	}
}

//Get : To get the value from the lru
func (l *LRU) Get(value string) string {
	_, found := Dict[value]
	if found == false {
		return "NOt present"
	}
	pointer := Dict[value]
	if pointer == l.Head {
		return l.Head.Value
	}
	if pointer.Next == nil {
		pointer.Prev = nil
		pointer.Next = l.Head
		l.Head.Prev = pointer
		l.Head.Next = nil
		l.Head = pointer
		return l.Head.Value
	}
	pointer.Prev.Next = pointer.Next
	pointer.Next.Prev = pointer.Prev
	pointer.Prev = nil
	pointer.Next = nil
	pointer.Next = l.Head
	l.Head.Prev = pointer
	l.Head = pointer
	return l.Head.Value

}

//Print : To print the LRU
func (l *LRU) Print() {
	current := l.Head
	ans := ""
	for {
		if current == nil {
			break
		}
		ans += current.Value + "-->"
		current = current.Next

	}
	println(ans)
}
