package bot

import (
	"github.com/notnil/chess"
	"math/rand"
	"sort"
)

// Huddle puts their pieces as close as possible to their own king.
func Huddle(p *chess.Position) *chess.Move {
	return closestMoveToKing(p, p.Turn())
}

// Swarm puts their pieces as close as possible to the opposing king.
func Swarm(p *chess.Position) *chess.Move {
	opposingColor := chess.White
	if p.Turn() == chess.White {
		opposingColor = chess.Black
	}
	return closestMoveToKing(p, opposingColor)
}

// closestMoveToKing returns a move with its dest square having the shortest distance to the king with targetKingColor.
func closestMoveToKing(p *chess.Position, targetKingColor chess.Color) *chess.Move {
	moves := p.ValidMoves()
	targetKing := chess.BlackKing
	if targetKingColor == chess.White {
		targetKing = chess.WhiteKing
	}
	kingSquare := func(kingWanted chess.Piece) chess.Square {
		for file := range []chess.File{chess.FileA, chess.FileB, chess.FileC, chess.FileD, chess.FileE, chess.FileF, chess.FileG, chess.FileH} {
			for rank := range []chess.Rank{chess.Rank1, chess.Rank2, chess.Rank3, chess.Rank4, chess.Rank5, chess.Rank6, chess.Rank7, chess.Rank8} {
				sq := chess.NewSquare(chess.File(file), chess.Rank(rank))
				if kingWanted == p.Board().Piece(sq) {
					return sq
				}
			}
		}
		return chess.NoSquare
	}(targetKing)
	type chebyMove struct {
		move              *chess.Move
		chebyshevDistance int8
	}
	chebyMoves := make([]chebyMove, 0, len(moves))
	for _, move := range moves {
		if move.S1() == kingSquare {
			continue
		}
		cheby := max(
			abs(int8(move.S2().File()-kingSquare.File())),
			abs(int8(move.S2().Rank()-kingSquare.Rank())),
		)
		chebyMoves = append(chebyMoves, chebyMove{
			move:              move,
			chebyshevDistance: cheby,
		})
	}
	sort.Slice(chebyMoves, func(i, j int) bool {
		return chebyMoves[i].chebyshevDistance < chebyMoves[j].chebyshevDistance
	})
	topMoves := make([]chebyMove, 0, len(chebyMoves))
	if len(chebyMoves) == 0 {
		return Random(p)
	}
	topMoves = append(topMoves, chebyMoves[0])
	for i, move := range chebyMoves {
		if i == 0 {
			continue
		}
		if topMoves[0].chebyshevDistance == move.chebyshevDistance {
			topMoves = append(topMoves, move)
			continue
		}
		break
	}
	return topMoves[rand.Intn(len(topMoves))].move
}

func abs(x int8) int8 {
	if x < 0 {
		return x * -1
	}
	return x
}

func max(x, y int8) int8 {
	if x > y {
		return x
	}
	return y
}
