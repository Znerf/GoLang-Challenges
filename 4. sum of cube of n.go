import (
    "math"
)

func FindNb(m int) int {
	data := 1 + 8*math.Sqrt(float64(m))
	
  result := ((1 + math.Sqrt(data)) / 2) - 1
	if result == float64(int(result)) {
		return int(result)
	}
	return -1
}

//test

package kata_test
import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  . "codewarrior/kata"
)

func dotest(n int, exp int) {
    var ans = FindNb(n)
    Expect(ans).To(Equal(exp))
}

var _ = Describe("Test Example", func() {

    It("should handle basic cases", func() {
        dotest(4183059834009, 2022)
        dotest(24723578342962, -1)
    })
})


// Your task is to construct a building which will be a pile of n cubes. The cube at the bottom will have a volume of 
// n^3 , the cube above will have volume of (n-1)^3  and so on until the top which will have a volume of 
// 1^3. You are given the total volume m of the building. Being given m can you find the number n of cubes you will have to build?

// The parameter of the function findNb (find_nb, find-nb, findNb, ...) will be an integer m and you have to return the integer n such as 

// n^3+ (n-1)^3... 1^3 =m if such a n exists or -1 if there is no such n.

// Examples:
// findNb(1071225) --> 45

// findNb(91716553919377) --> -1
