package _1xx

import (
	"bytes"
	"fmt"
	"strconv"
	"testing"
)

func (n *Node) String() string {
	qu := make([]*Node, 0)

	if n == nil {
		return "[]"
	}

	vi := make(map[int]bool)

	qu = append(qu, n)
	tmp := make([]int, 0)
	for len(qu) > 0 {
		p := qu[0]
		qu = qu[1:]

		if !vi[p.Val] {
			tmp = append(tmp, p.Val)
		}
		vi[p.Val] = true

		for _, v := range p.Neighbors {
			if !vi[v.Val] {
				qu = append(qu, v)
			}
		}
	}
	b := bytes.Buffer{}
	b.WriteString("[")
	b.WriteString(strconv.Itoa(tmp[0]))
	for i := 1; i < len(tmp); i++ {
		b.WriteString(", ")
		b.WriteString(strconv.Itoa(tmp[i]))
	}
	b.WriteString("]")
	return string(b.Bytes())
}

func buildNode(n [][]int) *Node {
	m := make(map[int]*Node)

	if len(n) == 0 {
		return nil
	}

	for i := 1; i <= len(n); i++ {
		m[i] = &Node{i, []*Node{}}
	}

	for i := 0; i < len(n); i++ {
		for j := 0; j < len(n[i]); j++ {
			m[i+1].Neighbors = append(m[i+1].Neighbors, m[n[i][j]])
		}
	}
	return m[1]
}

func TestCloneGraph(t *testing.T) {
	fmt.Println(buildNode([][]int{
		{2, 4}, {1, 3}, {2, 4}, {1, 3},
	}))
}
