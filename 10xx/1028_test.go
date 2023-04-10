package _0xx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
import (
	. "github.com/zhaowk/leetcode_go"
)

//输入："1-2--3--4-5--6--7"
//输出：[1,2,5,3,4,6,7]
//
//输入："1-2--3---4-5--6---7"
//输出：[1,2,5,3,null,6,null,4,null,7]
//
//输入："1-401--349---90--88"
//输出：[1,401,null,349,88,90]

func TestRecoverFromPreorder(t *testing.T) {

	//fmt.Println(recoverFromPreorder("1-2--3--4-5--6--7"))
	//fmt.Println(recoverFromPreorder("1-2--3---4-5--6---7"))
	//fmt.Println(recoverFromPreorder("1-401--349---90--88"))

	assert.Equal(t, BuildTree(1, 2, 5, 3, 4, 6, 7),
		recoverFromPreorder("1-2--3--4-5--6--7"))

	assert.Equal(t, BuildTree(1, 2, 5, 3, nil, 6, nil, 4, nil, 7),
		recoverFromPreorder("1-2--3---4-5--6---7"))

	assert.Equal(t, BuildTree(1, 401, nil, 349, 88, 90),
		recoverFromPreorder("1-401--349---90--88"))
}
