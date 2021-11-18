package leetcode_go

type WordDictionary struct {
	val map[uint8]*WordDictionary
	end bool
}

/** Initialize your data structure here. */
func ConstructorW() WordDictionary {
	return WordDictionary{
		val: make(map[uint8]*WordDictionary),
		end: false,
	}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	wd := this
	for _, c := range word {
		letter := uint8(c)
		if wd.val == nil {
			wd.val = map[uint8]*WordDictionary{
				letter: {},
			}
		} else if _, ok := wd.val[letter]; !ok {
			wd.val[letter] = &WordDictionary{}
		} else {
			// found
		}
		wd = wd.val[letter]
	}
	wd.end = true
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	wd := this
	for idx, c := range word {
		if wd.val == nil || len(wd.val) == 0 {
			return false
		}

		letter := uint8(c)
		if c == '.' {
			for _, item := range wd.val {
				if item.Search(word[idx+1:]) {
					return true
				}
			}
			return false
		} else if n, ok := wd.val[letter]; ok {
			wd = n
		} else {
			// not found
			return false
		}
	}
	return wd.end
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
