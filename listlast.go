package piscine

func ListLast(l *List) interface{} {
	var s interface{}
	for l.Head != nil {
		s = l.Head.Data
		l.Head = l.Head.Next
	}
	return s
}
