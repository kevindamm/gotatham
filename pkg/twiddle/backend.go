package twiddle

import "github.com/kevindamm/gotatham/pkg/puzzle"

// ...

type twiddleState struct {
  Width int
  Height int
  Size int

  Grid TwiddleGrid
  Completed bool
  Autosolve bool
  MoveCount int
  MoveTarget int
  PlayerDoes TwiddleMove
}

type TwiddleGrid []int

type TwiddleMove int

func NewTwiddleMove(x, y, size int, clockwise bool) {
  var tm int
  if clockwise {
    tm + 1
  }
  tm |= x & MASK_DIM << X_OFFSET
  tm |= y & MASK_DIM << Y_OFFSET
  tm |= size & MASK_SIZE << SIZE_OFFSET
  return TwiddleMove(tm)
}

func (t TwiddleMove) SelectX() int {
  return int(t)>>X_OFFSET & MASK_DIM
}

func (t TwiddleMove) SelectY() int {
  return int(t)>>Y_OFFSET & MASK_DIM
}

func (t TwiddleMove) Size() int {
  return int(t)>>SIZE_OFFSET & MASK_SIZE
}

func (t TwiddleMove) Direction() bool {
  return int(t) & 1 > 0
}
