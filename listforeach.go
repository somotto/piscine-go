package piscine

func ListForEach(l *List, f func(*NodeL)) {
	for current := l.Head; current != nil; current = current.Next {
		f(current)
	}
}

func Add2_node(node *NodeL) {
	switch node.Data.(type) {
	case int:
		node.Data = node.Data.(int) + 2
	case string:
		node.Data = node.Data.(string) + "2"
	}
}

func Subtract3_node(node *NodeL) {
	switch node.Data.(type) {
	case int:
		node.Data = node.Data.(int) - 3
	case string:
		node.Data = node.Data.(string) + "-3"
	}
}
