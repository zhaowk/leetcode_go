package _3xx

type NestedInteger struct {
	isInt bool
	val   int
	list  []*NestedInteger
}

func (this NestedInteger) IsInteger() bool { return this.isInt }

func (this NestedInteger) GetInteger() int {
	if this.IsInteger() {
		return this.val
	} else {
		return 0
	}
}

func (this *NestedInteger) SetInteger(value int) { this.isInt = true; this.val = value }

func (this *NestedInteger) Add(elem NestedInteger) {
	this.isInt = false
	if this.list == nil {
		this.list = []*NestedInteger{&elem}
	} else {
		this.list = append(this.list, &elem)
	}
}

func (this NestedInteger) GetList() []*NestedInteger {
	if this.IsInteger() {
		return []*NestedInteger{}
	}
	return this.list
}

/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */

type stackTyp struct {
	idx      int
	currList *[]*NestedInteger
}

type NestedIterator struct {
	stack []*stackTyp
}

func Constructor5(nestedList []*NestedInteger) *NestedIterator {
	it := &NestedIterator{}
	if len(nestedList) == 0 {
		return it
	}

	it.stack = []*stackTyp{{idx: 0, currList: &nestedList}}
	it.findNext()
	return it
}

func (this *NestedIterator) Next() int {
	i := (*this.stack[len(this.stack)-1].currList)[this.stack[len(this.stack)-1].idx].GetInteger()
	this.stack[len(this.stack)-1].idx++
	this.findNext()
	return i
}

func (this *NestedIterator) HasNext() bool {
	return len(this.stack) != 0
}

func (this *NestedIterator) findNext() {
	for len(this.stack) > 0 {
		idx := len(this.stack) - 1
		// list over
		if this.stack[idx].idx >= len(*this.stack[idx].currList) {
			this.stack = this.stack[:idx]
			if len(this.stack) > 0 {
				this.stack[len(this.stack)-1].idx++
			}

			continue
		}

		// find next integer
		if (*this.stack[idx].currList)[this.stack[idx].idx].IsInteger() {
			break
		}
		list := (*this.stack[idx].currList)[this.stack[idx].idx].GetList()
		// another nested list
		this.stack = append(this.stack, &stackTyp{
			idx:      0,
			currList: &list,
		})
	}
}
