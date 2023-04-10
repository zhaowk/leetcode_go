package _1xx

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */
//func connect(root *Node) *Node {
//
//	if root == nil {
//		return nil
//	}
//
//	root.Next = nil
//	prevHead := root
//	nextHead := root.Left
//
//	for prevHead != nil {
//		nextHead = prevHead.Left
//		p := prevHead
//		var prev *Node = nil
//
//		for p != nil {
//			if prev != nil {
//				prev.Next = p.Left
//			}
//
//			if p.Left != nil {
//				p.Left.Next = p.Right
//				prev = p.Right
//			} else {
//				return root
//			}
//
//			p = p.Next
//		}
//		prevHead = nextHead
//	}
//
//	return root
//}
