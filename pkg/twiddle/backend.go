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
  PreviousMove TwiddleMove
}

type TwiddleMove struct {
  VSelect int
  HSelect int
  Size int
  Direction int
}

type TwiddleGrid []int
