package leetcode_go

import (
	"fmt"
	"strconv"
)

var ipAddressesResult []string

func restoreIpAddresses(s string) []string {
	ipAddressesResult = make([]string, 0)
	ipAddresses(s, "", 0)
	return ipAddressesResult
}

func ipAddresses(s, t string, depth int) {

	if len(s) > 3*(4-depth) && len(s) < 4-depth {
		return
	}

	if depth == 4 {
		if len(s) == 0 {
			ipAddressesResult = append(ipAddressesResult, t)
		}
	} else {
		f := ""
		if depth == 0 {
			f = "%s%s"
		} else {
			f = "%s.%s"
		}

		for i := 1; i < 4; i++ {
			if i > len(s) {
				continue
			}
			if i > 1 && s[0] == '0' {
				continue
			}
			n, e := strconv.Atoi(s[:i])
			if e == nil && n < 256 {
				ipAddresses(s[i:], fmt.Sprintf(f, t, s[:i]), depth+1)
			}
		}
	}
}
