package bot

import (
	"github.com/notnil/chess"
	"math/rand"
	"time"
)

// Random plays random moves.
func Random(p *chess.Position) *chess.Move {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	moves := p.ValidMoves()
	return moves[r1.Intn(len(moves))]
}
