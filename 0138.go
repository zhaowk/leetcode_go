package leetcode_go

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

// Node 定义冲突
//type Node struct {
//	Val int
//	Next *Node
//	Random *Node
//}
//
// AC
//func copyRandomList(head *Node) *Node {
//
//	if head == nil {
//		return nil
//	}
//
//	pHead := &Node{}
//
//	tmp := make([]*Node, 0)
//	tmp2 := make([]*Node, 0)
//
//	for p:=head;p!=nil;p=p.Next {
//		tmp = append(tmp, p)
//	}
//
//	pHead.Next = &Node{head.Val,nil,nil}
//
//	for p, q := head, pHead; p!=nil;p=p.Next{
//		q.Next = &Node{p.Val,nil,nil}
//		q = q.Next
//		tmp2 = append(tmp2, q)
//	}
//
//	for p,q :=head,pHead.Next;p!=nil;p,q=p.Next,q.Next {
//		if p.Random == nil {
//			q.Random = nil
//		} else {
//			for idx, r := range tmp {
//				if p.Random == r {
//					q.Random = tmp2[idx]
//				}
//			}
//		}
//	}
//
//	return pHead.Next
//}
