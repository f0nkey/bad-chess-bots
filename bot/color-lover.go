package bot

import (
	"github.com/notnil/chess"
	"math/rand"
)

// OppositeColor puts their pieces on the opposite color (white tries to place pieces on dark squares).
// Otherwise, it plays a random move.
func OppositeColor(p *chess.Position) *chess.Move {
	myColor := p.Turn()
	desiredSquareColor := chess.Black
	if myColor == chess.Black {
		desiredSquareColor = chess.White
	}
	var preferredMoves []*chess.Move
	for _, move := range p.ValidMoves() {
		endingSquare := move.S2()
		squareColor := chess.White
		if ((endingSquare / 8) % 2) == (endingSquare % 2) { // https://github.com/notnil/chess/blob/master/square.go#L30
			squareColor = chess.Black
		}
		if squareColor == desiredSquareColor {
			preferredMoves = append(preferredMoves, move)
		}
	}
	if len(preferredMoves) == 0 {
		return Random(p)
	}
	return preferredMoves[rand.Intn(len(preferredMoves))]
}

// SameColor puts their pieces on the same color (white tries to place pieces on light squares).
// Otherwise, it plays a random move.
func SameColor(p *chess.Position) *chess.Move {
	myColor := p.Turn()
	desiredSquareColor := chess.White
	if myColor == chess.Black {
		desiredSquareColor = chess.Black
	}
	var preferredMoves []*chess.Move
	for _, move := range p.ValidMoves() {
		endingSquare := move.S2()
		squareColor := chess.White
		if ((endingSquare / 8) % 2) == (endingSquare % 2) { // https://github.com/notnil/chess/blob/master/square.go#L30
			squareColor = chess.Black
		}
		if squareColor == desiredSquareColor {
			preferredMoves = append(preferredMoves, move)
		}
	}
	if len(preferredMoves) == 0 {
		return Random(p)
	}
	return preferredMoves[rand.Intn(len(preferredMoves))]
}
