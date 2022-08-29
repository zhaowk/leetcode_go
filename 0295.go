package leetcode_go

import "container/heap"

type miHp struct {
	nums []int
}

func (m miHp) Len() int {
	return len(m.nums)
}

func (m miHp) Less(i, j int) bool {
	return m.nums[i] < m.nums[j]
}

func (m miHp) Swap(i, j int) {
	m.nums[i], m.nums[j] = m.nums[j], m.nums[i]
}

func (m *miHp) Push(x interface{}) {
	if i, ok := x.(int); ok {
		m.nums = append(m.nums, i)
	}
}

func (m *miHp) Pop() interface{} {
	if len(m.nums) == 0 {
		return nil
	}
	i := m.nums[len(m.nums) - 1]
	m.nums = m.nums[:len(m.nums) - 1]
	return i
}

type mxHp struct {
	nums []int
}

func (m mxHp) Len() int {
	return len(m.nums)
}

func (m mxHp) Less(i, j int) bool {
	return m.nums[i] > m.nums[j]
}

func (m mxHp) Swap(i, j int) {
	m.nums[i], m.nums[j] = m.nums[j], m.nums[i]
}

func (m *mxHp) Push(x interface{}) {
	if i, ok := x.(int); ok {
		m.nums = append(m.nums, i)
	}
}

func (m *mxHp) Pop() interface{} {
	if len(m.nums) == 0 {
		return nil
	}
	i := m.nums[len(m.nums) - 1]
	m.nums = m.nums[:len(m.nums) - 1]
	return i
}

type MedianFinder struct {
	mx *mxHp  // 较小的数， 使用大顶堆， 默认 len(mx) >= len(mi)
	mi *miHp  // 较大的数， 使用小顶堆
	size int
}

/** initialize your data structure here. */
func Constructor3() MedianFinder {
	return MedianFinder{
		mx: &mxHp{nums: make([]int, 0)},
		mi: &miHp{nums: make([]int, 0)},
	}
}

func (this *MedianFinder) AddNum(num int)  {
	if this.size == 0 {
		heap.Push(this.mx, num)
	} else if this.size == 1 {
		if this.mx.nums[0] <= num {
			heap.Push(this.mi, num)
		} else {
			heap.Push(this.mi, heap.Pop(this.mx))
			heap.Push(this.mx, num)
		}
	} else {
		if this.mx.nums[0] <= num {
			heap.Push(this.mi, num)
			// 调整
			if this.mx.Len() < this.mi.Len() {
				heap.Push(this.mx, heap.Pop(this.mi))
			}
		} else {
			if this.mx.Len() > this.mi.Len() {
				heap.Push(this.mi, heap.Pop(this.mx))
			}
			heap.Push(this.mx, num)
		}
	}

	this.size ++
}

func (this *MedianFinder) FindMedian() float64 {
	if this.size == 0 {
		return 0
	}
	if this.size & 1 == 1 {
		return float64(this.mx.nums[0])
	} else {
		return (float64(this.mi.nums[0]) + float64(this.mx.nums[0])) / 2.0
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */