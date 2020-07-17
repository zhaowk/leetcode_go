package leetcode_go

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	head := &Node{node.Val, nil}

	vi := make(map[int]bool)
	no := make(map[int]*Node)

	qu1 := make([]*Node, 0)
	qu2 := make([]*Node, 0)
	qu1 = append(qu1, node)
	qu2 = append(qu2, head)
	no[head.Val] = head

	for len(qu1) > 0 {
		p := qu1[0]
		qu1 = qu1[1:]

		q := qu2[0]
		qu2 = qu2[1:]

		if vi[q.Val] {
			continue
		}

		vi[q.Val] = true

		for _, n := range p.Neighbors {
			var r *Node
			if no[n.Val] == nil {
				r = &Node{n.Val, []*Node{}}
				no[n.Val] = r
			} else {
				r = no[n.Val]
			}
			q.Neighbors = append(q.Neighbors, r)

			if !vi[n.Val] {
				qu1 = append(qu1, n)
				qu2 = append(qu2, r)
			}
		}
	}

	return head
}
