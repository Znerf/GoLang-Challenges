package kata

func DigitalRoot(n int) int {
  // ...
  if n<10{
    return n
  }
  
  s,temp :=0,n
  for temp!=0{
    s += temp%10
    temp /= 10
  }
  return DigitalRoot(s)
  
}

/////////////////////////
// Ginkgo BDD Testing Framework <http://onsi.github.io/ginkgo></http:>
// Gomega Matcher Library <http://onsi.github.io/gomega></http:>

package kata_test

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  . "codewarrior/kata"
)

var _ = Describe("Test Example", func() {
  It("fixed tests", func() {
    Expect(DigitalRoot(16)).To(Equal(7))
    Expect(DigitalRoot(195)).To(Equal(6))
    Expect(DigitalRoot(992)).To(Equal(2))
    Expect(DigitalRoot(167346)).To(Equal(9))
    Expect(DigitalRoot(0)).To(Equal(0))
  })
})


/////////////////
// Question
//  16  -->  1 + 6 = 7
//    942  -->  9 + 4 + 2 = 15  -->  1 + 5 = 6
// 132189  -->  1 + 3 + 2 + 1 + 8 + 9 = 24  -->  2 + 4 = 6
// 493193  -->  4 + 9 + 3 + 1 + 9 + 3 = 29  -->  2 + 9 = 11  -->  1 + 1 = 2
