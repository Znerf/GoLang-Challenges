/*A child is playing with a ball on the nth floor of a tall building. The height of this floor above ground level, h, is known.

He drops the ball out of the window. The ball bounces (for example), to two-thirds of its height (a bounce of 0.66).

His mother looks out of a window 1.5 meters from the ground.

How many times will the mother see the ball pass in front of her window (including when it's falling and bouncing?

Three conditions must be met for a valid experiment:
Float parameter "h" in meters must be greater than 0
Float parameter "bounce" must be greater than 0 and less than 1
Float parameter "window" must be less than h.
If all three conditions above are fulfilled, return a positive integer, otherwise return -1.

Note:
The ball can only be seen if the height of the rebounding ball is strictly greater than the window parameter.

Examples:
- h = 3, bounce = 0.66, window = 1.5, result is 3

- h = 3, bounce = 1, window = 1.5, result is -1 

(Condition 2) not fulfilled).
8*/

package kata

func BouncingBall(h, bounce, window float64) int {
  // your code
  if bounce<0 || bounce >= 1 {return -1}
  times := -1
  for h > window {
    times= times+2
    h=h*(bounce)
  }
  return times
}

package kata_test
import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  . "codewarrior/kata"
)

func testequal(h, bounce, window float64, exp int) {
    var ans = BouncingBall(h, bounce, window)
    Expect(ans).To(Equal(exp))
}

var _ = Describe("Test Example", func() {

    It("should handle basic cases", func() {
      testequal(3, 0.66, 1.5, 3)
      testequal(40, 0.4, 10, 3)
      testequal(10, 0.6, 10, -1)
      testequal(40, 1, 10, -1)
      testequal(5, -1, 1.5, -1)
    })
})



func BouncingBall(h, bounce, window float64) int {
	if h <= window || bounce <= 0 || bounce >= 1 {
		return -1
	} else {
		return 2 + BouncingBall((h*bounce), bounce, window)
	}
}

func BouncingBall(h, bounce, window float64) int {
  failCond := h < 0 || bounce < 0 || bounce >= 1 || window >= h
  if (failCond) { return -1 }
  
  // Find n such that h*bounce^n = w
  n := math.Log(window/h) / math.Log(bounce)
  
  // Take into account that ball can only be seen when h*b^m > w ; m >= 1
  if (n == math.Floor(n)) { n-- }
  
  return (int(math.Floor(n))*2) + 1
}
/*
Learning points
1. recursive v while loop (memory extensive and easy clear code)
*/
