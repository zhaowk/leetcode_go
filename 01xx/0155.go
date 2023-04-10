package _1xx

type MinStack struct {
	Data []int
	Min  []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{[]int{}, []int{}}
}

func (this *MinStack) Push(x int) {
	if len(this.Min) > 0 && x >= this.Min[len(this.Min)-1] {
		this.Min = append(this.Min, this.Min[len(this.Min)-1])
	} else {
		this.Min = append(this.Min, x)
	}

	this.Data = append(this.Data, x)
}

func (this *MinStack) Pop() {
	if len(this.Data) > 0 {
		this.Data = this.Data[:len(this.Data)-1]
	}

	if len(this.Min) > 0 {
		this.Min = this.Min[:len(this.Min)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.Data) > 0 {
		return this.Data[len(this.Data)-1]
	}
	return -1
}

func (this *MinStack) GetMin() int {
	if len(this.Min) > 0 {
		return this.Min[len(this.Min)-1]
	}
	return -1
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
