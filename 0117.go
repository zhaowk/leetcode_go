package leetcode_go

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */
//func connect1(root *Node) *Node {
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
//		if prevHead.Left != nil {
//			nextHead = prevHead.Left
//		} else if prevHead.Right != nil {
//			nextHead = prevHead.Right
//		} else {
//			prevHead = prevHead.Next
//			continue
//		}
//
//		p := prevHead
//		var prev *Node = nil
//
//		for p != nil {
//			if prev != nil {
//				if p.Left != nil {
//					prev.Next = p.Left
//					prev = p.Left
//				}
//
//				if p.Right != nil {
//					prev.Next = p.Right
//					prev = p.Right
//				}
//			} else {
//				if p.Left != nil {
//					prev = p.Left
//				}
//
//				if p.Right != nil {
//					if prev != nil {
//						prev.Next = p.Right
//					}
//					prev = p.Right
//				}
//			}
//
//			p = p.Next
//		}
//		prevHead = nextHead
//	}
//
//	return root
//}
