package _2xx

import (
	"bytes"
	"encoding/binary"
	"unsafe"
)
import (
	. "github.com/zhaowk/leetcode_go"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
}

func Constructor4() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	q := make([]*TreeNode, 0)
	m := make(map[uintptr]int)

	if root == nil {
		return ""
	}

	q = append(q, root)
	m[uintptr(unsafe.Pointer(root))] = 0

	for idx := 0; idx < len(q); idx++ {
		if q[idx].Left != nil {
			m[uintptr(unsafe.Pointer(q[idx].Left))] = len(q)
			q = append(q, q[idx].Left)
		}

		if q[idx].Right != nil {
			m[uintptr(unsafe.Pointer(q[idx].Right))] = len(q)
			q = append(q, q[idx].Right)
		}
	}

	buf := bytes.NewBuffer(make([]byte, 0))

	for _, node := range q {
		_ = binary.Write(buf, binary.BigEndian, int16(node.Val))
		_ = binary.Write(buf, binary.BigEndian, int16(m[uintptr(unsafe.Pointer(node.Left))]))
		_ = binary.Write(buf, binary.BigEndian, int16(m[uintptr(unsafe.Pointer(node.Right))]))
	}
	return string(buf.Bytes())
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}

	m := make([]struct {
		n    *TreeNode
		l, r int16
	}, 0)

	buf := bytes.NewBufferString(data)

	for buf.Len() > 0 { // 3 * sizeof(int16)
		var v, l, r int16
		_ = binary.Read(buf, binary.BigEndian, &v)
		_ = binary.Read(buf, binary.BigEndian, &l)
		_ = binary.Read(buf, binary.BigEndian, &r)
		node := &TreeNode{Val: int(v)}
		m = append(m, struct {
			n    *TreeNode
			l, r int16
		}{n: node, l: l, r: r})
	}

	for _, v := range m {
		if v.l > 0 {
			v.n.Left = m[v.l].n
		}
		if v.r > 0 {
			v.n.Right = m[v.r].n
		}
	}

	if len(m) > 0 {
		return m[0].n
	}
	return nil
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
