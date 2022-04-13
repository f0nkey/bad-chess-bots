package main

import (
	"github.com/notnil/chess"
	"math/rand"
	"time"
)

// randomMove makes random moves.
func randomMove(p *chess.Position) *chess.Move {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	moves := p.ValidMoves()
	return moves[r1.Intn(len(moves))]
}
