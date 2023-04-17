package kata

func solution(str string, ending string) bool {
  // Your code here!
  start := len(str)-len(ending)
  if start <0{
    return false
  }
  a := str[start:]
  return a == ending
  

//////////////////////////////////////////////////////////
/*Complete the solution so that it returns true if the first argument(string) passed in ends with the 2nd argument (also a string).

Examples:

solution("abc", "bc") // returns true
solution("abc", "d") // returns false

https://www.codewars.com/kata/51f2d1cafc9c0f745c00037d/train/go

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


// Some solutions
package kata

import "strings"

func solution(str, ending string) bool {
	return strings.HasSuffix(str, ending)
}
	
package kata

func solution(str, ending string) bool {
  return len(str) >= len(ending) && str[len(str) - len(ending):] == ending
}
	
	
package kata

import "regexp"

func solution(str, ending string) bool {
  // Your code here!
  endsWith, _ := regexp.MatchString(ending + "$", str)
  return endsWith
  
}
	
// Insight 
// rune is also a datatype like char. it's ASCII representation 
