package main

import (
	"github.com/notnil/chess"
)

// oppositeColor puts their pieces on the opposite color (white tries to place pieces on dark squares).
// Otherwise, it plays a random move.
func oppositeColor(p *chess.Position) *chess.Move {
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
		return randomMove(p)
	}
	return preferredMoves[0]
}

// sameColor puts their pieces on the same color (white tries to place pieces on light squares).
// Otherwise, it plays a random move.
func sameColor(p *chess.Position) *chess.Move {
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
		return randomMove(p)
	}
	return preferredMoves[0]
}
