package leetcode_go

type DList struct {
	Key  int
	Val  int
	Prev *DList
	Next *DList
}
type LRUCache struct {
	LRUMap  map[int]*DList
	LRUHead *DList
	LRULen  int
	LRUCap  int
}

func DConstructor(capacity int) LRUCache {

	head := &DList{0, 0, nil, nil}
	head.Next = head
	head.Prev = head

	return LRUCache{make(map[int]*DList, 0), head, 0, capacity}
}

func (this *LRUCache) Get(key int) int {
	node := this.LRUMap[key]
	if node != nil {
		// 更新
		node.Prev.Next = node.Next
		node.Next.Prev = node.Prev

		node.Next = this.LRUHead.Next
		node.Prev = this.LRUHead
		node.Next.Prev = node
		this.LRUHead.Next = node
		return node.Val
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	node := this.LRUMap[key]

	if node != nil {
		// 更新
		node.Prev.Next = node.Next
		node.Next.Prev = node.Prev

		node.Next = this.LRUHead.Next
		node.Prev = this.LRUHead
		node.Next.Prev = node
		this.LRUHead.Next = node

		node.Val = value
	} else {
		// 添加
		node = &DList{key, value, nil, nil}

		if this.LRULen < this.LRUCap {
			this.LRULen++
		} else { // LRU
			delete(this.LRUMap, this.LRUHead.Prev.Key)

			this.LRUHead.Prev = this.LRUHead.Prev.Prev
			this.LRUHead.Prev.Next = this.LRUHead
		}

		this.LRUMap[key] = node

		node.Next = this.LRUHead.Next
		node.Prev = this.LRUHead
		node.Next.Prev = node
		this.LRUHead.Next = node

		// 循环 尾结点
		if this.LRUHead.Prev == this.LRUHead {
			this.LRUHead.Prev = node
		}
	}
}
