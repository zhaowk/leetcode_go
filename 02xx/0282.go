package _2xx

import "strconv"

const (
	oTypNum = iota
	oTypAdd
	oTypMinus
	oTypMul
)

type oOperator struct {
	typ int
	val int
}

type oOperators []oOperator

func (o oOperators) String() (res string) {
	for _, v := range o {
		switch v.typ {
		case oTypNum:
			res += strconv.Itoa(v.val)
		case oTypAdd:
			res += "+"
		case oTypMinus:
			res += "-"
		case oTypMul:
			res += "*"
		default:

		}
	}
	return
}

// 1 <= num.length <= 10, target 32bit signed integer
func addOperators(num string, target int) []string {
	res := make([]string, 0)
	var operate func(nums oOperators, suffix string)

	var compute = func(nums oOperators, target int) bool {
		if len(nums) == 0 {
			return false
		}

		if len(nums) == 1 {
			return false
		}

		tmp := make(oOperators, 0)
		lastIdx, lastNum := 0, nums[0].val
		for i := 1; i < len(nums); i += 2 {
			if lastIdx != i-1 {
				lastNum = nums[i-1].val
			}

			if nums[i].typ == oTypMul {
				lastIdx, lastNum = i+1, lastNum*nums[i+1].val
			} else {
				tmp = append(tmp, oOperator{typ: oTypNum, val: lastNum}, oOperator{typ: nums[i].typ})
				lastNum = nums[i+1].val
			}
		}

		tmp = append(tmp, oOperator{typ: oTypNum, val: lastNum})
		lastNum = tmp[0].val
		for i := 1; i < len(tmp); i += 2 {
			switch tmp[i].typ {
			case oTypAdd:
				lastNum += tmp[i+1].val
			case oTypMinus:
				lastNum -= tmp[i+1].val
			default:

			}
		}

		return lastNum == target
	}

	operate = func(nums oOperators, suffix string) {
		if len(suffix) == 0 {
			return
		}

		if val, ok := conv(suffix); ok {
			tmp := append(nums, oOperator{typ: oTypNum, val: val})
			if compute(tmp, target) {
				res = append(res, tmp.String())
			}
		}

		for i := 1; i < len(suffix); i++ {
			if val, ok := conv(suffix[:i]); ok {
				operate(append(nums, oOperator{val: val, typ: oTypNum}, oOperator{typ: oTypAdd}), suffix[i:])
				operate(append(nums, oOperator{val: val, typ: oTypNum}, oOperator{typ: oTypMinus}), suffix[i:])
				operate(append(nums, oOperator{val: val, typ: oTypNum}, oOperator{typ: oTypMul}), suffix[i:])
			}
		}
	}

	if n, ok := conv(num); ok && n == target {
		if len(num) == 1 || len(num) > 1 && num[0] != '0' {
			res = append(res, num)
		}
	}

	operate(oOperators{}, num)

	return res
}

func conv(s string) (res int, ok bool) {
	res, err := strconv.Atoi(s)

	if err != nil || len(s) == 0 {
		ok = false
		return
	}
	if len(s) == 1 {
		ok = true
		return
	}

	if s[0] == '0' {
		ok = false
		return
	}

	ok = true
	return
}
