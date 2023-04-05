package kata
import (
  "strings"
)
func DuplicateEncode(word string) string {
	var str string = ""
	for _, alpha := range strings.ToLower(word) {
		if strings.Count(strings.ToLower(word), string(alpha)) == 1 {
			str = str + "("
		} else {
			str = str + ")"
		}
	}
	return str
}

// TODO: replace with your own tests (TDD). An example to get you started is included below.
// Ginkgo BDD Testing Framework <http://onsi.github.io/ginkgo></http:>
// Gomega Matcher Library <http://onsi.github.io/gomega></http:>

package kata_test

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  . "codewarrior/kata"
)
var _ = Describe("Test Example", func() {
  It("should return the correct value", func() {
    Expect(DuplicateEncode("din")).To(Equal("((("))
    Expect(DuplicateEncode("recede")).To(Equal("()()()"))
    Expect(DuplicateEncode("(( @")).To(Equal("))(("))
  })

  It("should ignore case", func() {
    Expect(DuplicateEncode("Success")).To(Equal(")())())"))
  })
})

//The goal of this exercise is to convert a string to a new string where each character in the new string is "(" if that character appears only once in the original string, or ")" if that character appears more than once in the original string. Ignore capitalization when determining if a character is a duplicate.

//Examples
//"din"      =>  "((("
//"recede"   =>  "()()()"
//"Success"  =>  ")())())"
//"(( @"     =>  "))((" 
