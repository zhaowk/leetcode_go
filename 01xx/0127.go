package _1xx

func ladderLength(beginWord string, endWord string, wordList []string) int {

	found := false
	for _, w := range wordList {
		if w == endWord {
			found = true
			break
		}
	}

	if !found {
		return 0
	}

	mTop := make(map[string]bool, 0)
	mButtom := make(map[string]bool, 0)

	topDepth, buttomDepth := 1, 0
	prevTop := make([]string, 1)
	prevTop[0] = beginWord
	prevButtom := make([]string, 1)
	prevButtom[0] = endWord

	for topDepth <= len(wordList) && buttomDepth <= len(wordList) {
		//fmt.Println(prevTop, prevButtom)
		currTop := make([]string, 0)

		found = false
		for _, prevWord := range prevTop {
			for _, word := range wordList {
				if !mTop[word] && canTrans(prevWord, word) {
					found = true
					mTop[word] = true
					currTop = append(currTop, word)
				}
			}
		}
		if found {
			topDepth++
		} else {
			return 0
		}

		for _, prevW := range prevButtom {
			for _, currW := range currTop {
				if prevW == currW {
					return topDepth + buttomDepth
				}
			}
		}

		prevTop = currTop
		currButtom := make([]string, 0)
		found = false
		for _, prevWord := range prevButtom {
			for _, word := range wordList {
				if !mButtom[word] && canTrans(prevWord, word) {
					found = true
					mButtom[word] = true
					currButtom = append(currButtom, word)
				}
			}
		}
		if found {
			buttomDepth++
		} else {
			return 0
		}

		for _, prevW := range currButtom {
			for _, currW := range currTop {
				if prevW == currW {
					return topDepth + buttomDepth
				}
			}
		}
		prevButtom = currButtom
	}

	return 0

}

func canTrans(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	can := 0

	for i := range a {
		if a[i] != b[i] {
			can++
		}

		if can > 1 {
			return false
		}
	}

	return can == 1
}
