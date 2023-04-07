// Given a string of words, you need to find the highest scoring word.

// Each letter of a word scores points according to its position in the alphabet: a = 1, b = 2, c = 3 etc.

// For example, the score of abad is 8 (1 + 2 + 1 + 4).

// You need to return the highest scoring word as a string.

// If two words score the same, return the word that appears earliest in the original string.

// All letters will be lowercase and all inputs will be valid.

package kata
import(
  "unicode"
  "strings"
)
func High(s string) string {
  res := strings.Split(s, " ")
  scoreline:=score(res[0])
  word:= res[0]
  for _,data := range res{
    sco := score(data)
    if sco > scoreline{
      scoreline =sco
      word=data
    }
  }
  return word
}

func score(s string) int{
  var sum int
	for _, x := range s {
		ascii := int(unicode.ToLower(x)) - 96
		sum += ascii
	}
  return sum
}


package kata_test

import (
    "fmt"
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
    . "codewarrior/kata"
)

var _ = Describe("Example tests",func() {
    arr := [...][2]string{
        {"man i need a taxi up to ubud","taxi"},
        {"what time are we climbing up the volcano","volcano"},
        {"take me to semynak","semynak"},
        {"aa b", "aa"},
        {"b aa", "b"},
        {"bb d", "bb"},
        {"d bb", "d"},
        {"aaa b", "aaa"},
    }
    for _,v := range arr {
        var input = v[0]
        var output = v[1] 
        It(fmt.Sprintf("Testing input \"%s\"",input),func() {Expect(High(input)).To(Equal(output))})
    }
})
