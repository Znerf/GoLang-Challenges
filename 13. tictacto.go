package kata

func contain0(board [3][3]int) int{
  for i:=0 ; i<3 ; i++ {
     for j:=0 ; j<3 ; i++ {
        if board[i][j] == 0 {return -1}
    }
  }
  return 0
}

func IsSolved(board [3][3]int) int {
  // Your code here
  for i := 0; i < 3; i++ {
		if board[i][1] == board[i][2] && board[i][1] == board[i][0] && board[i][2] != 0 {
			return board[i][2]
		}
		if board[1][i] == board[2][i] && board[1][i] == board[0][i] && board[2][i] != 0 {
			return board[2][i]
		}
	}

	if (board[1][1] == board[2][2] && board[1][1] == board[0][0]) || (board[0][2] == board[1][1] && board[2][0] == board[1][1]) {
		return board[1][1]
	}  
  return contain0(board)
}

package kata_test

/*
If we were to set up a Tic-Tac-Toe game, we would want to know whether the board's current state is solved, wouldn't we? Our goal is to create a function that will check that for us!

Assume that the board comes in the form of a 3x3 array, where the value is 0 if a spot is empty, 1 if it is an "X", or 2 if it is an "O", like so:

[[0, 0, 1],
 [0, 1, 2],
 [2, 1, 0]]
We want our function to return:

-1 if the board is not yet finished AND no one has won yet (there are empty spots),
1 if "X" won,
2 if "O" won,
0 if it's a cat's game (i.e. a draw).
You may assume that the board passed in is valid in the context of a game of Tic-Tac-Toe
*/
import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  . "codewarrior/kata"
)

func doTest(board [3][3]int, expected int) {
  actual := IsSolved(board)
  Expect(actual).To(Equal(expected), "With board = %v\nExpected %d but got %d\n", board, expected, actual)
}

var _ = Describe("Test Example", func() {
  It("should test that the solution returns the correct value", func() {
    var board [3][3]int
    // not yet finished
    board = [3][3]int{
      {0, 0, 1},
      {0, 1, 2},
      {2, 1, 0},
    }
    doTest(board, -1)

    // winning row
    board = [3][3]int{
      {1, 1, 1},
      {0, 2, 2},
      {0, 0, 0},
    }
    doTest(board, 1)

    // winning column
    board = [3][3]int{
      {2, 1, 2},
      {2, 1, 1},
      {1, 1, 2},
    }
    doTest(board, 1)

    // draw
    board = [3][3]int{
      {2, 1, 2},
      {2, 1, 1},
      {1, 2, 1},
    }
    doTest(board, 0)
  })
})
