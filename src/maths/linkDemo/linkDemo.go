package linkDemo

type node struct {
	val  int
	next *node
}

//    cur  pre
//    [1]->[2]->[3]
//    [1]->[]
//cur= cur.next

// for curr != nil {
//        next := curr.Next
//        curr.Next = prev
//        prev = curr
//        curr = next
//    }
//round1
//cur: [1]->[2]->[3]
//next:  [2]-[3]
//cur.next: []
//pre: [1]->[]
//cur: [2]->[3]

//round2
//cur: [2]->[3]
//next: [3]
//cur.next: [2]->[1]
//pre: [2]->[1]
//cur=[3]

//round3
//cur:[3]
//next: []
//cur:[3]->[2]->[1]
//pre=cur [3]->[2]->[1]
//cur=[]
//return pre
func reverse(head *node) *node {
	cur, pre := head, &node{}

	for cur != nil {
		//[2]
		next := cur.next
		//[]
		cur.next = pre
		//[2]->[]
		pre = cur
		//[2]->[3]
		cur = next

	}

	return pre

}

//DFS是递归函数
//[1]->[2]-[3]
//head=[2]->[3]
//newHead=[3]
//head.next.next=[3]->[2]->[3]
//head.next=[]
//return [3]->[2]-[]
//newhead=[3]->[2]-[]

//round2
//head->[1]-[2]-[]
//newhead:=[3]-[2]-[]
//head.next.next=[2]-[1]
//header.next=nil=[3]-[2]-[1]

func DFS(head *node) *node {
	if head == nil || head.next == nil {
		return head
	}
	newHead := DFS(head.next)
	head.next.next = head
	head.next = nil
	return newHead
}

func dfsReverseLink(head *node) *node {
	if head.next == nil || head == nil {
		return head
	}

	next := dfsReverseLink(head.next)
	next.next = head
	head.next = nil
	return next
}

//func bpfReverseLink(head *node) *node{
//var    queueNode []node=make([]node,0)
//
//tmp:=head
//for tmp.next!=nil  {
//
//
//
//	tmp=head.next
//	queueNode=append(queueNode,*tmp)
//
//
//
//}
//
//
//}
