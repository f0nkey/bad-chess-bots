package bot

import (
	"github.com/notnil/chess"
	"math"
)

// Centrist tries to put all of its pieces in the center.
func Centrist(p *chess.Position) *chess.Move {
	type DistanceMove struct {
		move              *chess.Move
		currentDistance   float64
		resultingDistance float64
	}
	moves := make([]DistanceMove, 0, 0)
	for _, m := range p.ValidMoves() {
		currDist := distanceToCenter(m.S1())
		resDist := distanceToCenter(m.S2())
		if resDist > currDist || resDist == currDist {
			continue
		}
		if m.S1() == chess.D4 || m.S1() == chess.D5 || m.S1() == chess.E4 || m.S1() == chess.E5 {
			continue
		}
		moves = append(moves, DistanceMove{
			move:              m,
			currentDistance:   currDist,
			resultingDistance: resDist,
		})
	}

	if len(moves) == 0 {
		return Random(p)
	}

	smallestDiff := moves[0].resultingDistance
	smallestDiffMove := moves[0]
	if len(moves) > 1 {
		for i := 1; i < len(moves); i++ {
			diff := moves[i].resultingDistance
			if diff < smallestDiff {
				smallestDiff = diff
				smallestDiffMove = moves[i]
			}
		}
	}
	return smallestDiffMove.move
}

func distanceToCenter(sq chess.Square) float64 {
	lowestDistance := euclideanDistance(sq, chess.D4)
	if euclideanDistance(sq, chess.D5) < lowestDistance {
		lowestDistance = euclideanDistance(sq, chess.D5)
	}
	if euclideanDistance(sq, chess.E4) < lowestDistance {
		lowestDistance = euclideanDistance(sq, chess.E4)
	}
	if euclideanDistance(sq, chess.E5) < lowestDistance {
		lowestDistance = euclideanDistance(sq, chess.E5)
	}
	return lowestDistance
}

func euclideanDistance(a, b chess.Square) float64 {
	left := (float64(a.Rank()) - float64(b.Rank())) * (float64(a.Rank()) - float64(b.Rank()))
	right := (float64(a.File()) - float64(b.File())) * (float64(a.File()) - float64(b.File()))
	return math.Sqrt(left + right)
}
