package main

type Piece interface {
	// GetPossibleMoves(currentPosition int) []int
	// CalcNextPosition(currentPosition int) []int

	GetPieceType() string
	GetPieceTeam() string
	GetIconPath() string
}
