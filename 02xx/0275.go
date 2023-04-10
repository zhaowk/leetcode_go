package _2xx

func hIndex2(citations []int) int {
	if len(citations) == 1 {
		if citations[0] > 0 {
			return 1
		} else {
			return 0
		}
	}
	start, end := 0, len(citations)

	for start < end {
		if citations[start] >= len(citations)-start {
			return len(citations) - start
		}
		mid := (start + end) / 2
		if citations[mid] < len(citations)-mid {
			start = mid + 1
		} else if citations[mid] == len(citations)-mid {
			return len(citations) - mid
		} else {
			end = mid
		}

		if start+1 == end {
			if citations[start] >= len(citations)-start {
				return len(citations) - start
			}
			return len(citations) - end
		}
	}

	return len(citations) - end
}
