package kata

import (
"math"
)


func IsPrime(n int) bool {
  if n == 2 {
		return true
	}

	if n%2 == 0 || n < 2 {
		return false
	}

	for i := 3; i <= int(math.Sqrt(float64(n))); i = i + 2 {
		if n%i == 0 {
			return false
		}
	}

	return true
}



package kata

func IsPrime(n int) bool {
  if(n <= 3) {
    return n > 1;
  }
  if n % 2 == 0 || n % 3 == 0 {
    return false;
  }
  for i := 5; i * i <= n; i += 6 {
    if n % i == 0 || n % (i + 2) == 0 {
      return false;
    }
  }
  return true;
}
//test
// TODO: replace with your own tests (TDD). An example to get you started is included below.
// Ginkgo BDD Testing Framework <http://onsi.github.io/ginkgo/>
// Gomega Matcher Library <http://onsi.github.io/gomega/>

package kata_test
import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  . "codewarrior/kata"
)

func doTest(n int, expected bool) {
  Expect(IsPrime(n)).To(Equal(expected), "With n = %d", n)
}

var _ = Describe("Basic", func() {
   It("Basic tests", func() {
     doTest(0, false)
     doTest(1, false)
     doTest(2, true)
     doTest(73, true)
     doTest(75, false)
     doTest(-1, false)
   })

   It("Test prime", func() {
     doTest(3, true)
     doTest(5, true)
     doTest(7, true)
     doTest(41, true)
     doTest(5099, true)
   })

   It("Test not prime", func() {
     doTest(4, false)
     doTest(6, false)
     doTest(8, false)
     doTest(9, false)
     doTest(45, false)
     doTest(-5, false)
     doTest(-8, false)
     doTest(-41, false)
   })
})

//Define a function that takes an integer argument and returns a logical value true or false depending on if the integer is a prime.

//Per Wikipedia, a prime number ( or a prime ) is a natural number greater than 1 that has no positive divisors other than 1 and itself.

//Requirements
//You can assume you will be given an integer input.
//You can not assume that the integer will be only positive. You may be given negative numbers as well ( or 0 ).
//NOTE on performance: There are no fancy optimizations required, but still the most trivial solutions might time out. Numbers go up to 2^31 ( or similar, depending on language ). Looping all the way up to n, or n/2, will be too slow.
