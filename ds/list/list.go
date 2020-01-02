package list

type ListNode struct {
	Val int
	Next *ListNode
}


type List struct {
	Root *ListNode
}

func (l *List) Add(val int) {
	node := &ListNode{
		Val:val,
		Next:nil,
	}
	if l.Root == nil {
		l.Root = node
	} else{
		var curr *ListNode = l.Root
		for curr.Next != nil {
			curr = curr.Next
		}
		curr.Next = node
	}
}