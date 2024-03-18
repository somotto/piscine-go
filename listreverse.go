package piscine

func ListReverse(l *List) {
	var prev, current *NodeL
	for current = l.Head; current != nil; current, prev = prev, current {
		current.Next = prev
	}
	l.Head, l.Tail = l.Tail, l.Head
}
