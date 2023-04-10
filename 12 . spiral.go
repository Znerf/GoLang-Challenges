package kata

func Snail(snaipMap [][]int) []int {
	res := []int{}
	i := len(snaipMap)
	len := i

	for {
		for j := len - i; j < i; j++ {
			res = append(res, snaipMap[len-i][j])
		}
		if i == 1 {
			break
		}
		for j := len - i + 1; j <= i-1; j++ {
			res = append(res, snaipMap[j][i-1])
		}

		for j := i - 2; j >= len-i; j-- {
			res = append(res, snaipMap[i-1][j])
		}
		for j := i - 2; j >= len-i+1; j-- {
			res = append(res, snaipMap[j][len-i])
		}
		i--
	}
	return res
}
/*
Snail Sort
Given an n x n array, return the array elements arranged from outermost elements to the middle element, traveling clockwise.

array = [[1,2,3],
         [4,5,6],
         [7,8,9]]
snail(array) #=> [1,2,3,6,9,8,7,4,5]
For better understanding, please follow the numbers of the next array consecutively:

array = [[1,2,3],
         [8,9,4],
         [7,6,5]]
snail(array) #=> [1,2,3,4,5,6,7,8,9]
This image will illustrate things more clearly:*/

// TODO: replace with your own tests (TDD). An example to get you started is included below.
// Ginkgo BDD Testing Framework <http://onsi.github.io/ginkgo/>
// Gomega Matcher Library <http://onsi.github.io/gomega/>

package kata_test

import (
	. "codewarrior/kata"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("snail tests function", func() {
	It("should resolve 3x3", func() {
		snailMap := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
		Expect(Snail(snailMap)).To(Equal([]int{1, 2, 3, 6, 9, 8, 7, 4, 5}))
	})
})

const snail = (arr) => {
  var finalArray = [];
  while (arr.length) {
    finalArray.push(...arr.shift());
    arr.map(row => finalArray.push(row.pop()))
    arr.reverse().map(row => row.reverse());
  }
  return finalArray
}

