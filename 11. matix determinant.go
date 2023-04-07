//matrix determinant

package kata

func Determinant(matrix [][]int) int {
      // your code here
  if len(matrix) == 1 {
    return Determinant1(matrix)
  }
  if len(matrix) == 2 {
    return Determinant2(matrix)
  }
  if len(matrix) == 3 {
    sum:=0
    sum += matrix[0][0] * Determinant2([][]int{{matrix[1][1],matrix[1][2]},{matrix[2][1],matrix[2][2]} })
    sum -= matrix[0][1] * Determinant2([][]int{{matrix[1][0],matrix[1][2]},{matrix[2][0],matrix[2][2]} })
    sum += matrix[0][2] * Determinant2([][]int{{matrix[1][0],matrix[1][1]},{matrix[2][0],matrix[2][1]} })
    return sum 
  }
  return 0
 
}
func Determinant1(matrix [][] int) int{
  return matrix[0][0]
}
func Determinant2(matrix [][] int) int {
  return matrix[0][0]*matrix[1][1]-matrix[0][1]*matrix[1][0]
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
       Expect(Determinant([][]int{{1}})).To(Equal(1))
       Expect(Determinant([][]int{{1, 3}, {2, 5}})).To(Equal(-1))
       Expect(Determinant([][]int{{2, 5, 3}, {1, -2, -1}, {1, 3, 4}})).To(Equal(-20))
     })
})
