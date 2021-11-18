package leetcode_go

type Trie struct {
	Val   int
	Tries []*Trie
}

/** Initialize your data structure here. */
func Constructor1() Trie {
	return Trie{0, []*Trie{}}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	t := this
	for _, w := range word {
		var next *Trie = nil
		for _, v := range t.Tries {
			if v != nil && v.Val == int(w) {
				next = v
			}
		}

		if next == nil {
			n := &Trie{int(w), []*Trie{}}
			t.Tries = append(t.Tries, n)
			next = n
		}
		t = next
	}
	t.Tries = append(t.Tries, nil)
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	t := this
	for _, w := range word {
		if t == nil {
			return false
		}
		var next *Trie = nil
		for _, v := range t.Tries {
			if v != nil && v.Val == int(w) {
				next = v
			}
		}

		if next == nil {
			return false
		}
		t = next
	}

	for _, v := range t.Tries {
		if v == nil {
			return true
		}
	}
	return false
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	t := this
	for _, w := range prefix {
		var next *Trie = nil
		for _, v := range t.Tries {
			if v != nil && v.Val == int(w) {
				next = v
			}
		}

		if next == nil {
			return false
		}
		t = next
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
