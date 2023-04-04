package kata


func MinMax(arr []int) [2]int {
  min, max := arr[0],arr[0]
  for _, s := range arr {
    if s>max{
      max=s
    }
    if s<min{
      min =s
    }
  }
  
  if min == max {
     return [2]int{min}
  }
  return [2]int{min,max}
}

// TODO: replace with your own tests (TDD). An example to get you started is included below.
// Ginkgo BDD Testing Framework <http://onsi.github.io/ginkgo/>
// Gomega Matcher Library <http://onsi.github.io/gomega/>

package kata_test
import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  . "codewarrior/kata"
)


var _ = Describe("Tests", func() {
     It("Sample tests", func() {
       Expect(MinMax([]int{1,2,3,4,5})).To(Equal([2]int{1,5}))
       Expect(MinMax([]int{2334454,5})).To(Equal([2]int{5,2334454}))
     })
})


// Write a function that returns both the minimum and maximum number of the given list/array.

// Examples (Input --> Output)
// [1,2,3,4,5] --> [1,5]
// [2334454,5] --> [5,2334454]
// [1]         --> [1,1]
