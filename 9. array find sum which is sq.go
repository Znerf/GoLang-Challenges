package kata
import (
  "math"
)
func ListSquared(m, n int) [][]int {
	// your code
	res := [][]int{}
	for i := m; i <= n; i++ {
		sum := 0
		for j := 1; j <= i; j++ {
			if i%j == 0 {
				sum = sum + j*j
			}
		}

		if math.Sqrt(float64(sum)) == float64(int(math.Sqrt(float64(sum)))) {
			a := []int{i, sum}
			res = append(res, a)
		}
	}
	return res
}

// information 
/*
 i learnt that array doesn't need to be defined before hand as in c and c++
 You can use make([]int, size) but this one is effective as this is memory less extentive.
 Use of append()
 difference between initialization and assingent as assigment uses memoery space but initiliazation is null and no any value assigned only address
*/

package kata_test
import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  . "codewarrior/kata"
)

func dotest(m, n int, exp [][]int) {
    var ans = ListSquared(m, n)
    Expect(ans).To(Equal(exp))
}

var _ = Describe("Test Example", func() {

    It("should handle basic cases", func() {
        dotest(1, 250, [][]int{{1, 1}, {42, 2500}, {246, 84100}})
        dotest(250, 500, [][]int{{287, 84100}})
        dotest(300, 600, [][]int{})
    })
})
