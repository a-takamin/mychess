package main

type Piece interface {
	CalcPossibleNextMove(currentPosition int) []*Move

	GetPieceType() string
	GetPieceTeam() string
	GetIconPath() string
}
