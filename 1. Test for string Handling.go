package kata

func solution(str string, ending string) bool {
  // Your code here!
  start := len(str)-len(ending)
  if start <0{
    return false
  }
  a := str[start:]
  return a == ending
  

////////////////////////////////////////////////////////
/*Complete the solution so that it returns true if the first argument(string) passed in ends with the 2nd argument (also a string).

Examples:

solution("abc", "bc") // returns true
solution("abc", "d") // returns false



package kata

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
)

var _ = Describe("Example test:", func() {
	It("Should work for fixed tests", func() {
		Expect(solution("", "")).To(Equal(true))
		Expect(solution(" ", "")).To(Equal(true))
		Expect(solution("abc", "c")).To(Equal(true))
		Expect(solution("banana", "ana")).To(Equal(true))
		Expect(solution("a", "z")).To(Equal(false))
		Expect(solution("", "t")).To(Equal(false))
		Expect(solution("$a = $b + 1", "+1")).To(Equal(false))
		Expect(solution("    ", "   ")).To(Equal(true))
		Expect(solution("1oo", "100")).To(Equal(false))
	})
}) 
*/
