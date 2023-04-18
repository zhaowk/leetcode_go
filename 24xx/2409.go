package _4xx

import "fmt"

type date struct {
	m, d int
}

func (d date) Less(d2 date) bool {
	return d.m < d2.m || (d.m == d2.m && d.d < d2.d)
}

func (d date) Days(d2 date) int {
	days := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if d.m == d2.m {
		return d2.d - d.d
	}

	var r int
	for i := d.m; i < d2.m-1; i++ {
		r += days[i]
	}
	return r + days[d.m-1] - d.d + d2.d
}

func min(d1, d2 date) date {
	if d1.Less(d2) {
		return d1
	}
	return d2
}

func max(d1, d2 date) date {
	if d1.Less(d2) {
		return d2
	}
	return d1
}

func countDaysTogether(arriveAlice string, leaveAlice string, arriveBob string, leaveBob string) int {
	var alice1, alice2, bob1, bob2 date

	_, _ = fmt.Sscanf(arriveAlice, "%d-%d", &alice1.m, &alice1.d)
	_, _ = fmt.Sscanf(leaveAlice, "%d-%d", &alice2.m, &alice2.d)
	_, _ = fmt.Sscanf(arriveBob, "%d-%d", &bob1.m, &bob1.d)
	_, _ = fmt.Sscanf(leaveBob, "%d-%d", &bob2.m, &bob2.d)

	if alice2.Less(bob1) || bob2.Less(alice1) {
		return 0
	}

	return max(alice1, bob1).Days(min(alice2, bob2)) + 1
}
